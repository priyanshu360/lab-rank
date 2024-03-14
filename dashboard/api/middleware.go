package api

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/priyanshu360/lab-rank/dashboard/config"
	"github.com/priyanshu360/lab-rank/dashboard/internal/auth"
	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	psql "github.com/priyanshu360/lab-rank/dashboard/repository/postgres"
	redisrepo "github.com/priyanshu360/lab-rank/dashboard/repository/redis"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL.Path, req.Header)
		if strings.HasPrefix(req.URL.Path, "/auth") || !config.AuthEnabled() {
			next.ServeHTTP(w, req)
			return
		}
		// Extract JWT token from the request headers
		jwtToken := req.Header.Get("Authorization")
		if jwtToken == "" || !strings.HasPrefix(jwtToken, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(jwtToken, "Bearer ")
		log.Println(token)

		svc := auth.New(psql.NewAuthPostgresRepo(db), redisrepo.NewRedisSessionRepository(rClient), syllabus.New(psql.NewSyllabusPostgresRepo(db)))
		authSession, appErr := svc.Authenticate(context.Background(), token)
		if appErr != models.NoError {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		type AuthKey string
		nReq := req.WithContext(context.WithValue(req.Context(), AuthKey("authSession"), authSession))
		next.ServeHTTP(w, nReq)
	})
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
