package handler

import (
	"context"
	"log"
	"net/http"

	university_svc "github.com/priyanshu360/lab-rank/dashboard/internal/university"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type universityHandler struct {
    svc university_svc.UniversityService
}

func NewUniversityHandler(svc university_svc.UniversityService) Handler {
    return &universityHandler{
        svc: svc,
    }
}

func (h *universityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *universityHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
    var request models.CreateUniversityAPIRequest
    if err := request.Parse(r); err != nil {
        log.Println(err)
        return newAPIError(models.BadRequest.Add(err))
    }

    university := request.ToUniversity()
    university, err := h.svc.Create(ctx, university)
    if err != models.NoError {
        return newAPIError(models.InternalError.Add(err))
    }

    return models.NewCreateUniversityAPIResponse(university)
}


func (h *universityHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	university,err := h.svc.Fetch(ctx,id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateUniversityAPIResponse(university) // Reusing the same Response from Create in Get 
}

// Implement other handler methods for university-related operations
