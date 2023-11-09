package handler

import (
	"context"
	"log"
	"net/http"

	subject_svc "github.com/priyanshu360/lab-rank/dashboard/internal/subject"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type subjectHandler struct {
    svc subject_svc.SubjectService
}

func NewSubjectHandler(svc subject_svc.SubjectService) Handler {
    return &subjectHandler{
        svc: svc,
    }
}

func (h *subjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *subjectHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
    var request models.CreateSubjectAPIRequest
    if err := request.Parse(r); err != nil {
        log.Println(err)
        return newAPIError(models.BadRequest.Add(err))
    }

    subject := request.ToSubject() // Implement ToSubject() method in your models
    subject, err := h.svc.Create(ctx, subject)
    if err != models.NoError {
        return newAPIError(models.InternalError.Add(err))
    }

    return models.NewCreateSubjectAPIResponse(subject)
}

func (h *subjectHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	subject,err := h.svc.Fetch(ctx,id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateSubjectAPIResponse(subject) // Reusing the same Response from Create in Get 
}
