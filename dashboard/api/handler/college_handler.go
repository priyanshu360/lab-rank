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
	case http.MethodGet:
		response = h.handleGet(ctx, r)
	case http.MethodPut:
		response = h.handleUpdate(ctx, r)
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

func (h *collegeHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	limit := r.URL.Query().Get("limit")
	colleges, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(colleges) == 1 {
		return models.NewCreateCollegeAPIResponse(colleges[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListCollegesAPIResponse(colleges)
		return response
	}
}

func (h *collegeHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	var request models.UpdateCollegeAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}
	user, err := h.svc.Update(ctx, &request)
	if err != models.NoError {
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewCreateCollegeAPIResponse(user) // Reusing the same Response for Update
}
// Implement other handler methods for college-related operations
