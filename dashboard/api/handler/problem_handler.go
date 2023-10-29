package handler

import (
	"context"
	"log"
	"net/http"

	problems_svc "github.com/priyanshu360/lab-rank/dashboard/internal/problem"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type problemsHandler struct {
    svc problems_svc.ProblemService
}

func NewProblemsHandler(svc problems_svc.ProblemService) Handler {
    return &problemsHandler{
        svc: svc,
    }
}

func (h *problemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var response apiResponse
    var ctx = r.Context()

    switch r.Method {
    case http.MethodPost:
        response = h.handleCreate(ctx, r)
    // Implement other HTTP methods as needed
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    if err := response.Write(w); err != nil {
        http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
    }
}

func (h *problemsHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
    var request models.CreateProblemAPIRequest
    if err := request.Parse(r); err != nil {
        log.Println(err)
        return newAPIError(models.BadRequest.Add(err))
    }

    problem := request.ToProblem()
    problem, err := h.svc.Create(ctx, problem)
    if err != models.NoError {
        return newAPIError(models.InternalError.Add(err))
    }

    return models.NewCreateProblemAPIResponse(problem)
}

// Implement other handler methods for problems-related operations
