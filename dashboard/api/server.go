package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Import Gorilla Mux
	"github.com/priyanshu360/lab-rank/dashboard/api/handler"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/internal/college"
	"github.com/priyanshu360/lab-rank/dashboard/internal/environment"
	"github.com/priyanshu360/lab-rank/dashboard/internal/problem"
	"github.com/priyanshu360/lab-rank/dashboard/internal/subject"
	"github.com/priyanshu360/lab-rank/dashboard/internal/submission"
	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/internal/university"
	"github.com/priyanshu360/lab-rank/dashboard/internal/user"
	filesys "github.com/priyanshu360/lab-rank/dashboard/repository/fs"
	psql "github.com/priyanshu360/lab-rank/dashboard/repository/postgres"
	"github.com/priyanshu360/lab-rank/queue/queue"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var db *gorm.DB
var clientset *kubernetes.Clientset

// ServerConfig is your server configuration interface.

// APIServer is your API server instance.
type APIServer struct {
	Config   config.ServerConfig
	Router   *mux.Router // Replace ServeMux with Gorilla Mux
	Handlers map[string]handler.Handler
}

func NewServer(cfg config.ServerConfig) *APIServer {
	return &APIServer{
		Config:   cfg,
		Router:   mux.NewRouter(), // Initialize Gorilla Mux router
		Handlers: make(map[string]handler.Handler),
	}
}

// TODO: change this not looking good (maybe option or decorator pattern)
func (s *APIServer) initRoutes() {
	fileStorage := filesys.NewK8sCMStore(clientset, "storage")
	// Todo : make queue name env / handle error
	publisher, _ := queue.InitRabbitMQPublisher("lab-rank")
	// Initialize routes and handlers for different entities
	userHandler := handler.NewUserHandler(user.NewUserService(psql.NewUserPostgresRepo(db)))
	s.Handlers["/user"] = handler.NewReqIDMiddleware().Decorate(userHandler)

	subjectHandler := handler.NewSubjectHandler(subject.NewSubjectService(psql.NewSubjectPostgresRepo(db)))
	s.Handlers["/subject"] = handler.NewReqIDMiddleware().Decorate(subjectHandler)

	collegeHandler := handler.NewCollegeHandler(college.NewCollegeService(psql.NewCollegePostgresRepo(db)))
	s.Handlers["/college"] = handler.NewReqIDMiddleware().Decorate(collegeHandler)

	universityHandler := handler.NewUniversityHandler(university.NewUniversityService(psql.NewUniversityPostgresRepo(db)))
	s.Handlers["/university"] = handler.NewReqIDMiddleware().Decorate(universityHandler)

	submissionsHandler := handler.NewSubmissionsHandler(submission.NewSubmissionService(psql.NewSubmissionPostgresRepo(db), fileStorage, publisher))
	s.Handlers["/submission"] = handler.NewReqIDMiddleware().Decorate(submissionsHandler)

	environmentHandler := handler.NewEnvironmentHandler(environment.NewEnvironmentService(psql.NewEnvironmentPostgresRepo(db), fileStorage))
	s.Handlers["/environment"] = handler.NewReqIDMiddleware().Decorate(environmentHandler)

	problemsHandler := handler.NewProblemsHandler(problem.NewProblemService(psql.NewProblemPostgresRepo(db), fileStorage))
	s.Handlers["/problem"] = handler.NewReqIDMiddleware().Decorate(problemsHandler)

	syllabusHandler := handler.NewSyllabusHandler(syllabus.NewSyllabusService(psql.NewSyllabusPostgresRepo(db)))
	s.Handlers["/syllabus"] = handler.NewReqIDMiddleware().Decorate(syllabusHandler)
}

func (s *APIServer) run() {
	address := s.Config.GetAddress()
	port := s.Config.GetPort()
	addr := fmt.Sprintf("%s:%s", address, port)

	fmt.Printf("APIServer is running on http://%s\n", addr)

	for route, handler := range s.Handlers {
		s.Router.Handle(route, handler)
	}

	http.Handle("/", s.Router) // Set the Gorilla Mux router as the default handler

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func StartHttpServer(cfg config.ServerConfig) {
	server := NewServer(cfg)
	server.initRoutes()
	server.run()
}

func InitDB(cfg config.DBConfig) {
	var err error
	if db, err = gorm.Open(postgres.Open(cfg.GetURL()), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}

	err = db.Exec("SET search_path TO lab_rank").Error
	if err != nil {
		log.Fatal(err)
	}

	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println(table)
	}

}

func InitK8sClientset(kubeconfig string) error {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}
