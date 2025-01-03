package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

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
	h.cRouter.HandleFunc("/college/names/{university_id}", serveHTTPWrapper(h.handleGetName, models.AccessLevelStudent)).Methods("GET")
	h.cRouter.HandleFunc("/college", serveHTTPWrapper(h.handleGet, models.AccessLevelStudent)).Methods("GET")
	h.cRouter.HandleFunc("/college", serveHTTPWrapper(h.handleCreate, models.AccessLevelAdmin)).Methods("POST")

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
	universityId, err := strconv.Atoi(vars["university_id"])
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
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	college, err := h.svc.Fetch(ctx, id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	return models.NewCreateCollegeAPIResponse(college)
}

// Implement other handler methods for college-related operations
