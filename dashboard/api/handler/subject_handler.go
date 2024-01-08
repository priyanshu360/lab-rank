package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	subject_svc "github.com/priyanshu360/lab-rank/dashboard/internal/subject"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type subjectHandler struct {
	svc     subject_svc.Service
	sRouter *mux.Router
}

func NewSubjectHandler(svc subject_svc.Service) *subjectHandler {
	h := &subjectHandler{
		svc:     svc,
		sRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *subjectHandler) initRoutes() *subjectHandler {
	h.sRouter.HandleFunc("/subject/{university_id}", serveHTTPWrapper(h.handleGetByUniversityID, models.AccessLevelStudent)).Methods("GET")
	h.sRouter.HandleFunc("/subject", serveHTTPWrapper(h.handleGet, models.AccessLevelStudent)).Methods("GET")
	h.sRouter.HandleFunc("/subject", serveHTTPWrapper(h.handleCreate, models.AccessLevelAdmin)).Methods("POST")
	// h.sRouter.HandleFunc("/subjects", serveHTTPWrapper(h.handleUpdate)).Methods("PUT")
	// Add other routes as needed

	return h
}

func (h *subjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.sRouter.ServeHTTP(w, r)
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

func (h *subjectHandler) handleGetByUniversityID(ctx context.Context, r *http.Request) apiResponse {
	vars := mux.Vars(r)
	universityId, err := uuid.Parse(vars["university_id"])
	if err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	subjects, appError := h.svc.FetchByUniversityID(ctx, universityId)
	if appError != models.NoError {
		return newAPIError(appError)
	}
	response := models.NewListSubjectsAPIResponse(subjects)
	return response
}

func (h *subjectHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	limit := r.URL.Query().Get("limit")
	subjects, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(subjects) == 1 {
		return models.NewCreateSubjectAPIResponse(subjects[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListSubjectsAPIResponse(subjects)
		return response
	}
}
