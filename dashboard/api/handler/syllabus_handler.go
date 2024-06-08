package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	syllabus_svc "github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type syllabusHandler struct {
	svc     syllabus_svc.Service
	sRouter *mux.Router
}

func NewSyllabusHandler(svc syllabus_svc.Service) *syllabusHandler {
	h := &syllabusHandler{
		svc:     svc,
		sRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *syllabusHandler) initRoutes() *syllabusHandler {
	h.sRouter.HandleFunc("/syllabus", serveHTTPWrapper(h.handleGet, models.AccessLevelStudent)).Methods("GET")
	h.sRouter.HandleFunc("/syllabus/by_subject/{subject_id}", serveHTTPWrapper(h.handleGetBySubjectID, models.AccessLevelStudent)).Methods("GET")

	h.sRouter.HandleFunc("/syllabus", serveHTTPWrapper(h.handleCreate, models.AccessLevelAdmin)).Methods("POST")
	// h.sRouter.HandleFunc("/syllabus", serveHTTPWrapper(h.handleUpdate)).Methods("PUT")
	// h.sRouter.HandleFunc("/syllabus", serveHTTPWrapper(h.handleDelete)).Methods("DELETE")
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
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	syllabus, err := h.svc.Fetch(ctx, id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateSyllabusAPIResponse(syllabus) // Reusing the same Response from Create in Get

}

// Implement other handler methods for syllabus-related operations

func (h *syllabusHandler) handleGetBySubjectID(ctx context.Context, r *http.Request) apiResponse {
	vars := mux.Vars(r)
	subjectID, err := strconv.Atoi(vars["subject_id"])
	if err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	syllabus, err := h.svc.FetchBySubjectID(ctx, subjectID)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	response := models.NewListSyllabusAPIResponse(syllabus)
	return response
}
