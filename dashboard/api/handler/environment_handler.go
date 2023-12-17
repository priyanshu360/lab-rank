package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	environment_svc "github.com/priyanshu360/lab-rank/dashboard/internal/environment"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type environmentHandler struct {
	svc     environment_svc.Service
	eRouter *mux.Router
}

func NewEnvironmentHandler(svc environment_svc.Service) *environmentHandler {
	h := &environmentHandler{
		svc:     svc,
		eRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *environmentHandler) initRoutes() *environmentHandler {
	h.eRouter.HandleFunc("/environment", serveHTTPWrapper(h.handleGet)).Methods("GET")
	h.eRouter.HandleFunc("/environment", serveHTTPWrapper(h.handleCreate)).Methods("POST")
	// Add other routes as needed

	return h
}

func (h *environmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.eRouter.ServeHTTP(w, r)
}
func (h *environmentHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateEnvironmentAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	environment := request.ToEnvironment()
	environment, err := h.svc.Create(ctx, environment)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateEnvironmentAPIResponse(environment)
}

func (h *environmentHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	limit := r.URL.Query().Get("limit")
	environments, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(environments) == 1 {
		return models.NewCreateEnvironmentAPIResponse(environments[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListEnvironmentsAPIResponse(environments)
		return response
	}
}

// Implement other handler methods for environment-related operations
