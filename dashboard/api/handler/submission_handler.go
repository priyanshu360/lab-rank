package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/priyanshu360/lab-rank/dashboard/internal/submission"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type submissionsHandler struct {
	svc submission.SubmissionService
}

func NewSubmissionsHandler(svc submission.SubmissionService) Handler {
	return &submissionsHandler{
		svc: svc,
	}
}

func (h *submissionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *submissionsHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateSubmissionAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	submission := request.ToSubmissions()
	submission, err := h.svc.Create(ctx, submission)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateSubmissionAPIResponse(submission)
}

// Implement other handler methods for submissions-related operations
