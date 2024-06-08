package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	problems_svc "github.com/priyanshu360/lab-rank/dashboard/internal/problem"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type problemsHandler struct {
	svc     problems_svc.Service
	pRouter *mux.Router
}

func NewProblemsHandler(svc problems_svc.Service) *problemsHandler {
	h := &problemsHandler{
		svc:     svc,
		pRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *problemsHandler) initRoutes() *problemsHandler {
	h.pRouter.HandleFunc("/problem/{subject_id}/{college_id}", serveHTTPWrapper(h.handleGetProblemsForSubject, models.AccessLevelStudent)).Methods("GET")
	h.pRouter.HandleFunc("/problem", serveHTTPWrapper(h.handleGet, models.AccessLevelStudent)).Methods("GET")
	h.pRouter.HandleFunc("/problem", serveHTTPWrapper(h.handleCreate, models.AccessLevelTeacher)).Methods("POST")
	// Add other routes as needed

	return h
}

func (h *problemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.pRouter.ServeHTTP(w, r)
}

func (h *problemsHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateProblemAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	problem := request.ToProblem()
	problem, err := h.svc.Create(ctx, problem)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateProblemAPIResponse(problem)
}

func (h *problemsHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	// Todo : cleanup
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	problem, err := h.svc.Fetch(ctx, id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateProblemAPIResponse(problem)

}

// Implement other handler methods for problems-related operations
func (h *problemsHandler) handleGetProblemsForSubject(ctx context.Context, r *http.Request) apiResponse {
	vars := mux.Vars(r)
	subjectID, err := strconv.Atoi(vars["subject_id"])
	if err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	collegeID, err := strconv.Atoi(vars["college_id"])
	if err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	problems, appError := h.svc.GetProblemsForSubject(ctx, subjectID, collegeID)
	if appError != models.NoError {
		return newAPIError(appError)
	}
	response := models.NewListProblemsAPIResponse(problems)
	return response
}
