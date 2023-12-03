package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	Problems_svc "github.com/priyanshu360/lab-rank/dashboard/internal/problem"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type ProblemsHandler struct {
	svc Problems_svc.ProblemService
}

func NewProblemsHandler(svc Problems_svc.ProblemService) Handler {
	return &ProblemsHandler{
		svc: svc,
	}
}

func (h *ProblemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response apiResponse
	var ctx = r.Context()

	switch r.Method {
	case http.MethodPost:
		response = h.handleCreate(ctx, r)
	case http.MethodGet:
		response = h.handleGet(ctx, r)
	case http.MethodPut:
		response = h.handleUpdate(ctx, r)
	// Implement other HTTP methods as needed
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}

func (h *ProblemsHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateProblemAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	Problem := request.ToProblem()
	Problem, err := h.svc.Create(ctx, Problem)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateProblemAPIResponse(Problem)
}

func (h *ProblemsHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	// Todo : cleanup
	lang := r.URL.Query().Get("lang")
	id := r.URL.Query().Get("id")

	// todo : fix this hack
	if id != "" && lang != "" {
		uuid, _ := uuid.Parse(id)
		resp, _ := h.svc.GetInitCode(ctx, uuid, lang)
		return resp
	}

	limit := r.URL.Query().Get("limit")
	Problems, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(Problems) == 1 {
		return models.NewCreateProblemAPIResponse(Problems[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListProblemsAPIResponse(Problems)
		return response
	}
}

func (h *ProblemsHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	var request models.UpdateProblemAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}
	user, err := h.svc.Update(ctx, &request)
	if err != models.NoError {
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewCreateProblemAPIResponse(user) // Reusing the same Response for Update
}
// Implement other handler methods for Problems-related operations
