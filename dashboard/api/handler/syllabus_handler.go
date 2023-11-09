package handler

import (
	"context"
	"log"
	"net/http"

	syllabus_svc "github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type syllabusHandler struct {
    svc syllabus_svc.SyllabusService
}

func NewSyllabusHandler(svc syllabus_svc.SyllabusService) Handler {
    return &syllabusHandler{
        svc: svc,
    }
}

func (h *syllabusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var response apiResponse
    var ctx = r.Context()

    switch r.Method {
    case http.MethodPost:
        response = h.handleCreate(ctx, r)
    case http.MethodGet:
        response = h.handleGet(ctx, r)
    // Implement other HTTP methods as needed
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    if err := response.Write(w); err != nil {
        http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
    }
}

func (h *syllabusHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
    var request models.CreateSyllabusAPIRequest
    if err := request.Parse(r); err != nil {
        log.Println(err)
        return newAPIError(models.BadRequest.Add(err))
    }

    syllabus := request.ToSyllabus()
    syllabus, err := h.svc.Create(ctx, syllabus)
    if err != models.NoError {
        return newAPIError(models.InternalError.Add(err))
    }

    return models.NewCreateSyllabusAPIResponse(syllabus)
}


func (h *syllabusHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	syllabus,err := h.svc.Fetch(ctx,id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateSyllabusAPIResponse(syllabus) // Reusing the same Response from Create in Get 
}

// Implement other handler methods for syllabus-related operations
