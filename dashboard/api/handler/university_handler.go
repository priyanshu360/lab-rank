package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	university_svc "github.com/priyanshu360/lab-rank/dashboard/internal/university"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type universityHandler struct {
	svc     university_svc.UniversityService
	uRouter *mux.Router
}

func NewUniversityHandler(svc university_svc.UniversityService) *universityHandler {
	h := &universityHandler{
		svc:     svc,
		uRouter: mux.NewRouter(),
	}

	return h.InitRoutes()
}

func (h *universityHandler) InitRoutes() *universityHandler {
	h.uRouter.HandleFunc("/university", ServeHTTPWrapper(h.handleGet)).Methods("GET")
	h.uRouter.HandleFunc("/university", ServeHTTPWrapper(h.handleCreate)).Methods("POST")
	// h.uRouter.HandleFunc("/university", ServeHTTPWrapper(h.handleUpdate)).Methods("PUT")
	// h.uRouter.HandleFunc("/university", ServeHTTPWrapper(h.handleDelete)).Methods("DELETE")
	// Add other routes as needed

	return h
}

func (h *universityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.uRouter.ServeHTTP(w, r)
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
	limit := r.URL.Query().Get("limit")
	universities, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(universities) == 1 {
		return models.NewCreateUniversityAPIResponse(universities[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListUniversitiesAPIResponse(universities)
		return response
	}
}

// Implement other handler methods for university-related operations
