package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Import Gorilla Mux
	"github.com/priyanshu360/lab-rank/dashboard/api/handler"
	user_svc "github.com/priyanshu360/lab-rank/dashboard/internal/user"
	psql "github.com/priyanshu360/lab-rank/dashboard/repository/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ServerConfig is your server configuration interface.
type ServerConfig interface {
	GetAddress() string
	GetPort() string
}

// APIServer is your API server instance.
type APIServer struct {
	Config   ServerConfig
	Router   *mux.Router // Replace ServeMux with Gorilla Mux
	Handlers map[string]handler.Handler
}

func NewServer(config ServerConfig) *APIServer {
	return &APIServer{
		Config:   config,
		Router:   mux.NewRouter(), // Initialize Gorilla Mux router
		Handlers: make(map[string]handler.Handler),
	}
}

func (s *APIServer) initRoutes() {
	// Initialize routes and handlers for different entities
    userHandler := handler.NewUserHandler(user_svc.NewUserService(psql.NewUserPostgresRepo(db)))
    s.Handlers["/user"] = handler.NewReqIDMiddleware().Decorate(userHandler)

    subjectHandler := handler.NewSubjectHandler(subject_svc.NewSubjectService(psql.NewSubjectPostgresRepo(db)))
    s.Handlers["/subject"] = handler.NewReqIDMiddleware().Decorate(subjectHandler)

    collegeHandler := handler.NewCollegeHandler(college_svc.NewCollegeService(psql.NewCollegePostgresRepo(db)))
    s.Handlers["/college"] = handler.NewReqIDMiddleware().Decorate(collegeHandler)

    universityHandler := handler.NewUniversityHandler(university_svc.NewUniversityService(psql.NewUniversityPostgresRepo(db)))
    s.Handlers["/university"] = handler.NewReqIDMiddleware().Decorate(universityHandler)

    submissionsHandler := handler.NewSubmissionsHandler(submissions_svc.NewSubmissionsService(psql.NewSubmissionsPostgresRepo(db)))
    s.Handlers["/submissions"] = handler.NewReqIDMiddleware().Decorate(submissionsHandler)

    environmentHandler := handler.NewEnvironmentHandler(environment_svc.NewEnvironmentService(psql.NewEnvironmentPostgresRepo(db)))
    s.Handlers["/environment"] = handler.NewReqIDMiddleware().Decorate(environmentHandler)

    problemsHandler := handler.NewProblemsHandler(problems_svc.NewProblemsService(psql.NewProblemsPostgresRepo(db)))
    s.Handlers["/problems"] = handler.NewReqIDMiddleware().Decorate(problemsHandler)

    syllabusHandler := handler.NewSyllabusHandler(syllabus_svc.NewSyllabusService(psql.NewSyllabusPostgresRepo(db)))
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

func StartHttpServer(config ServerConfig) {
	server := NewServer(config)
	server.initRoutes()
	server.run()
}

func InitDB(){
	dbURL := "postgres://baeldung:baeldung@localhost:5432/postgres"
	var err error 
	if db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{}); err != nil {
		log.Fatal(err)
	}

}