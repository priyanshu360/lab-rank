package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-jwt/jwt"
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
	"github.com/priyanshu360/lab-rank/dashboard/models"
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
	rbac        map[http.Handler]models.AccessLevelModeEnum
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
		rbac:        make(map[http.Handler]models.AccessLevelModeEnum),
	}
}

func (s *APIServer) add(path string, role models.AccessLevelModeEnum, handler http.Handler) {
	s.router.PathPrefix(path).Handler(handler)
	s.rbac[handler] = role
}

func (s *APIServer) RbacMiddleware(next http.Handler) http.Handler {
	requiredRole := s.rbac[next]
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !config.AuthEnabled() {
			next.ServeHTTP(w, req)
		}
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if _, err := validateTokenWithRole(authHeader, requiredRole); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		next.ServeHTTP(w, req)
	})
}

func (s *APIServer) initRoutesAndMiddleware() {
	fileStorage := filesys.NewK8sCMStore(clientset, "storage")
	publisher, err := queue.InitRabbitMQPublisher("lab-rank")
	if err != nil {
		log.Fatal(err)
	}

	s.add("/auth", models.AccessLevelAdmin, handler.NewAuthHandler(auth.New(psql.NewAuthPostgresRepo(db), syllabus.New(psql.NewSyllabusPostgresRepo(db)))))
	s.add("/user", models.AccessLevelAdmin, handler.NewUserHandler(user.New(psql.NewUserPostgresRepo(db))))
	s.add("/subject", models.AccessLevelAdmin, handler.NewSubjectHandler(subject.New(psql.NewSubjectPostgresRepo(db), syllabus.New(psql.NewSyllabusPostgresRepo(db)))))
	s.add("/college", models.AccessLevelAdmin, handler.NewCollegeHandler(college.New(psql.NewCollegePostgresRepo(db), syllabus.New(psql.NewSyllabusPostgresRepo(db)))))
	s.add("/university", models.AccessLevelAdmin, handler.NewUniversityHandler(university.New(psql.NewUniversityPostgresRepo(db))))
	s.add("/submission", models.AccessLevelStudent, handler.NewSubmissionsHandler(submission.New(psql.NewSubmissionPostgresRepo(db), fileStorage, publisher)))
	s.add("/environment", models.AccessLevelAdmin, handler.NewEnvironmentHandler(environment.New(psql.NewEnvironmentPostgresRepo(db), fileStorage)))
	s.add("/problem", models.AccessLevelAdmin, handler.NewProblemsHandler(problem.New(psql.NewProblemPostgresRepo(db), fileStorage)))
	s.add("/syllabus", models.AccessLevelAdmin, handler.NewSyllabusHandler(syllabus.New(psql.NewSyllabusPostgresRepo(db))))

	s.middlewares = []mux.MiddlewareFunc{
		mux.CORSMethodMiddleware(s.router),
		handler.NewReqIDMiddleware().Decorate,
		OptionMiddleware,
		// s.RbacMiddleware,
	}
	s.router.Use(s.middlewares...)
	s.httpServer.Handler = s.router
}

func OptionMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusAccepted)
			return
		}

		next.ServeHTTP(w, req)
	})
}
func validateTokenWithRole(tokenString string, requiredRole models.AccessLevelModeEnum) (*jwt.StandardClaims, error) {
	// Replace the following with your own secret key
	secretKey := []byte("your_secret_key")

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	// Extract and return the claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("Failed to extract claims")
	}

	if claims.Subject != string(requiredRole) {
		return nil, errors.New("Unauthorised")
	}
	return claims, nil
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
