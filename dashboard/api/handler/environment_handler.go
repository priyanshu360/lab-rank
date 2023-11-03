package handler

import (
	"context"
	"log"
	"net/http"

	environment_svc "github.com/priyanshu360/lab-rank/dashboard/internal/environment"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type environmentHandler struct {
	svc environment_svc.EnvironmentService
}

func NewEnvironmentHandler(svc environment_svc.EnvironmentService) Handler {
	return &environmentHandler{
		svc: svc,
	}
}

func (h *environmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response apiResponse
	var ctx = r.Context()

	switch r.Method {
	case http.MethodPost:
		response = h.handleCreate(ctx, r)
	case http.MethodGet:
		response = h.handleGet(ctx, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
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
	environment,err := h.svc.Fetch(ctx,id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateEnvironmentAPIResponse(environment) // Reusing the same Response from Create in Get 
}

// Implement other handler methods for environment-related operations
