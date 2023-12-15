package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux" // Import Gorilla Mux
	"github.com/priyanshu360/lab-rank/dashboard/api/handler"
	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/internal/auth"
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
	httpServer  *http.Server
	middlewares []mux.MiddlewareFunc
	router      *mux.Router
}

func NewServer(cfg config.ServerConfig) *APIServer {
	return &APIServer{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", cfg.GetAddress(), cfg.GetPort()),
			WriteTimeout: time.Duration(cfg.GetWriteTimeout()) * time.Second,
			ReadTimeout:  time.Duration(cfg.GetReadTimeout()) * time.Second,
		},
		middlewares: []mux.MiddlewareFunc{},
		router:      mux.NewRouter(),
	}
}

func (s *APIServer) add(path string, handler http.Handler) {
	s.router.PathPrefix(path).Handler(handler)
}

func (s *APIServer) initRoutesAndMiddleware() {
	fileStorage := filesys.NewK8sCMStore(clientset, "storage")
	publisher, err := queue.InitRabbitMQPublisher("lab-rank")
	if err != nil {
		log.Fatal(err)
	}

	s.add("/auth", handler.NewAuthHandler(auth.New(psql.NewAuthPostgresRepo(db), syllabus.New(psql.NewSyllabusPostgresRepo(db)))))
	s.add("/user", handler.NewUserHandler(user.New(psql.NewUserPostgresRepo(db))))
	s.add("/subject", handler.NewSubjectHandler(subject.New(psql.NewSubjectPostgresRepo(db), syllabus.New(psql.NewSyllabusPostgresRepo(db)))))
	s.add("/college", handler.NewCollegeHandler(college.New(psql.NewCollegePostgresRepo(db), syllabus.New(psql.NewSyllabusPostgresRepo(db)))))
	s.add("/university", handler.NewUniversityHandler(university.New(psql.NewUniversityPostgresRepo(db))))
	s.add("/submission", handler.NewSubmissionsHandler(submission.New(psql.NewSubmissionPostgresRepo(db), fileStorage, publisher)))
	s.add("/environment", handler.NewEnvironmentHandler(environment.New(psql.NewEnvironmentPostgresRepo(db), fileStorage)))
	s.add("/problem", handler.NewProblemsHandler(problem.New(psql.NewProblemPostgresRepo(db), fileStorage)))
	s.add("/syllabus", handler.NewSyllabusHandler(syllabus.New(psql.NewSyllabusPostgresRepo(db))))

	s.middlewares = []mux.MiddlewareFunc{
		mux.CORSMethodMiddleware(s.router),
		handler.NewReqIDMiddleware().Decorate,
	}
	s.router.Use(s.middlewares...)
	s.httpServer.Handler = s.router
}

func (s *APIServer) run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	}()

	log.Println("[*] Server running .... ")

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Println("Received signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		fmt.Println("Error during server shutdown:", err)
	}
	fmt.Println("Server gracefully stopped")
}

func StartHttpServer(cfg config.ServerConfig) {
	server := NewServer(cfg)
	server.initRoutesAndMiddleware()
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
