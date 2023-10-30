package handler

import (
	"context"
	"log"
	"net/http"

	college_svc "github.com/priyanshu360/lab-rank/dashboard/internal/college"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type collegeHandler struct {
	svc college_svc.CollegeService
}

func NewCollegeHandler(svc college_svc.CollegeService) Handler {
	return &collegeHandler{
		svc: svc,
	}
}

func (h *collegeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *collegeHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateCollegeAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	college := request.ToCollege()
	college, err := h.svc.Create(ctx, college)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateCollegeAPIResponse(college)
}

// Implement other handler methods for college-related operations
