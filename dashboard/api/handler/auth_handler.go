package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priyanshu360/lab-rank/dashboard/internal/auth"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type authHandler struct {
	svc     auth.Service
	aRouter *mux.Router
}

func NewAuthHandler(svc auth.Service) *authHandler {
	h := &authHandler{
		svc:     svc,
		aRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *authHandler) initRoutes() *authHandler {
	log.Print("auth handler")
	h.aRouter.HandleFunc("/auth/login", serveHTTPWrapper(h.handleLogin)).Methods("POST")
	h.aRouter.HandleFunc("/auth/signup", serveHTTPWrapper(h.handleSignUp)).Methods("POST")
	// h.aRouter.HandleFunc("/auth/logout", serveHTTPWrapper(h.handleLogout)).Methods("POST")
	// h.aRouter.HandleFunc("/auth/reset", serveHTTPWrapper(h.handleLogout)).Methods("PUT")
	// Add other routes as needed

	return h
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("auth handler called")
	h.aRouter.ServeHTTP(w, r)
}

func (h *authHandler) handleLogin(ctx context.Context, r *http.Request) apiResponse {
	var req models.LoginAPIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	email, password := req.Email, req.Password
	response, err := h.svc.Login(ctx, email, password)
	if err != models.NoError {
		return newAPIError(err)
	}
	return response
}

func (h *authHandler) handleLogout(ctx context.Context, r *http.Request) apiResponse {
	return nil
}

func (h *authHandler) handleSignUp(ctx context.Context, r *http.Request) apiResponse {
	var req models.SignUpAPIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	user := req.ToUser()
	password := req.Password
	err := h.svc.SignUp(ctx, user, password)
	if err != models.NoError {
		return newAPIError(err)
	}
	return newAPIError(models.NoError)
}
