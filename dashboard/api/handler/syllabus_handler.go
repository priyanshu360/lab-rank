package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	syllabus_svc "github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type syllabusHandler struct {
	svc     syllabus_svc.SyllabusService
	sRouter *mux.Router
}

func NewSyllabusHandler(svc syllabus_svc.SyllabusService) *syllabusHandler {
	h := &syllabusHandler{
		svc:     svc,
		sRouter: mux.NewRouter(),
	}

	return h.InitRoutes()
}

func (h *syllabusHandler) InitRoutes() *syllabusHandler {
	h.sRouter.HandleFunc("/syllabus", ServeHTTPWrapper(h.handleGet)).Methods("GET")
	h.sRouter.HandleFunc("/syllabus", ServeHTTPWrapper(h.handleCreate)).Methods("POST")
	// h.sRouter.HandleFunc("/syllabus", ServeHTTPWrapper(h.handleUpdate)).Methods("PUT")
	// h.sRouter.HandleFunc("/syllabus", ServeHTTPWrapper(h.handleDelete)).Methods("DELETE")
	// Add other routes as needed

	return h
}

func (h *syllabusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.sRouter.ServeHTTP(w, r)
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
	limit := r.URL.Query().Get("limit")
	syllabus, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(syllabus) == 1 {
		return models.NewCreateSyllabusAPIResponse(syllabus[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListSyllabusAPIResponse(syllabus)
		return response
	}
}

// Implement other handler methods for syllabus-related operations
