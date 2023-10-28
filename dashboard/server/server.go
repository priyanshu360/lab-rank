// server/server.go

package server

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
	userHandler := handler.NewUserHandler(user_svc.NewUserService(psql.NewUserPostgresRepo(db)))
	s.Handlers["/user"] = handler.NewReqIDMiddleware().Decorate(userHandler)
	// Initialize with the repository
	// Add more routes and custom handlers as needed
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