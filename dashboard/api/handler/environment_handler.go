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
	case http.MethodPut:
		response = h.handleUpdate(ctx, r)
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

func (h *environmentHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	var request models.UpdateEnvironmentAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}
	user, err := h.svc.Update(ctx, &request)
	if err != models.NoError {
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewCreateEnvironmentAPIResponse(user) // Reusing the same Response for Update
}
// Implement other handler methods for environment-related operations
