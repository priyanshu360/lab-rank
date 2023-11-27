package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
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
	case http.MethodGet:
		response = h.handleGet(ctx, r)
	case http.MethodPut:
		response = h.handleUpdate(ctx, r)
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

func (h *submissionsHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	limit := r.URL.Query().Get("limit")
	submissions, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(submissions) == 1 {
		return models.NewCreateSubmissionAPIResponse(submissions[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListSubmissionsAPIResponse(submissions)
		return response
	}
}

func (h *submissionsHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	id, err := uuid.Parse(r.URL.Query().Get("id"))

	var request models.UpdateSubmissionAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	submission := request.ToSubmissions()
	submission, err = h.svc.Update(ctx, id, submission)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewUpdateSubmissionAPIResponse(submission)
}
