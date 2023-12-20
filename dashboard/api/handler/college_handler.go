package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	college_svc "github.com/priyanshu360/lab-rank/dashboard/internal/college"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type collegeHandler struct {
	svc     college_svc.Service
	cRouter *mux.Router
}

func NewCollegeHandler(svc college_svc.Service) *collegeHandler {
	h := &collegeHandler{
		svc:     svc,
		cRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *collegeHandler) initRoutes() *collegeHandler {
	h.cRouter.HandleFunc("/college/names/{university_id}", serveHTTPWrapper(h.handleGetName)).Methods("GET")
	h.cRouter.HandleFunc("/college", serveHTTPWrapper(h.handleGet)).Methods("GET")
	h.cRouter.HandleFunc("/college", serveHTTPWrapper(h.handleCreate)).Methods("POST")

	return h
}

func (h *collegeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.cRouter.ServeHTTP(w, r)
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

func (h *collegeHandler) handleGetName(ctx context.Context, r *http.Request) apiResponse {
	vars := mux.Vars(r)
	universityId, err := uuid.Parse(vars["university_id"])
	if err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}
	response, svcErr := h.svc.GetCollegeIDsForUniversityID(ctx, universityId)
	if svcErr != models.NoError {
		return newAPIError(svcErr)
	}
	return models.NewListCollegesIdNamesAPIResponse(response)
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

// Implement other handler methods for college-related operations
