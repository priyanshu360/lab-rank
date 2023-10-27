// server/server.go

package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux" // Import Gorilla Mux
	"github.com/priyanshu360/lab-rank/dashboard/api/handler"
	user_handler "github.com/priyanshu360/lab-rank/dashboard/api/handler/user"
	"github.com/priyanshu360/lab-rank/dashboard/repository/postgres"
	user_svc "github.com/priyanshu360/lab-rank/dashboard/service/user"
)

// ServerConfig is your server configuration interface.
type ServerConfig interface {
	GetAddress() string
	GetPort() string
}

// HTTPServer is your HTTP server instance.
type HTTPServer struct {
	Config   ServerConfig
	Router   *mux.Router // Replace ServeMux with Gorilla Mux
	Handlers map[string]handler.Handler
}

func NewServer(config ServerConfig) *HTTPServer {
	return &HTTPServer{
		Config:   config,
		Router:   mux.NewRouter(), // Initialize Gorilla Mux router
		Handlers: make(map[string]handler.Handler),
	}
}

func (s *HTTPServer) initRoutes() {
	s.Handlers["/user"] = user_handler.NewUserHandler(user_svc.NewUserAPIService(postgres.NewPostgresRepo())) // Initialize with the repository
	// Add more routes and custom handlers as needed
}

func (s *HTTPServer) run() {
	address := s.Config.GetAddress()
	port := s.Config.GetPort()
	addr := fmt.Sprintf("%s:%s", address, port)

	fmt.Printf("HTTPServer is running on http://%s\n", addr)

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
