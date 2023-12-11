package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	subject_svc "github.com/priyanshu360/lab-rank/dashboard/internal/subject"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type subjectHandler struct {
	svc     subject_svc.SubjectService
	sRouter *mux.Router
}

func NewSubjectHandler(svc subject_svc.SubjectService) *subjectHandler {
	h := &subjectHandler{
		svc:     svc,
		sRouter: mux.NewRouter(),
	}

	return h.InitRoutes()
}

func (h *subjectHandler) InitRoutes() *subjectHandler {
	h.sRouter.HandleFunc("/subjects", ServeHTTPWrapper(h.handleGet)).Methods("GET")
	h.sRouter.HandleFunc("/subjects", ServeHTTPWrapper(h.handleCreate)).Methods("POST")
	// h.sRouter.HandleFunc("/subjects", ServeHTTPWrapper(h.handleUpdate)).Methods("PUT")
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
