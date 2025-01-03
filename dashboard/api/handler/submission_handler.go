package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/priyanshu360/lab-rank/dashboard/internal/submission"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type submissionsHandler struct {
	svc     submission.Service
	sRouter *mux.Router
}

func NewSubmissionsHandler(svc submission.Service) *submissionsHandler {
	h := &submissionsHandler{
		svc:     svc,
		sRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *submissionsHandler) initRoutes() *submissionsHandler {

	h.sRouter.HandleFunc("/submission/user/{user_id}", serveHTTPWrapper(h.handleGetForUserID, models.AccessLevelStudent)).Methods("GET")
	h.sRouter.HandleFunc("/submission/problem/{problem_id}", serveHTTPWrapper(h.handleGet, models.AccessLevelStudent)).Methods("GET")
	h.sRouter.HandleFunc("/submission", serveHTTPWrapper(h.handleGet, models.AccessLevelStudent)).Methods("GET")
	h.sRouter.HandleFunc("/submission", serveHTTPWrapper(h.handleCreate, models.AccessLevelStudent)).Methods("POST")
	h.sRouter.HandleFunc("/submission", serveHTTPWrapper(h.handleUpdate, models.AccessLevelAdmin)).Methods("PUT")
	// Add other routes as needed

	return h
}

func (h *submissionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.sRouter.ServeHTTP(w, r)
}

func (h *submissionsHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateSubmissionAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	submission := request.ToSubmissions()
	submission, err := h.svc.Create(ctx, submission)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateSubmissionAPIResponse(submission)
}

func (h *submissionsHandler) handleGetForUserID(ctx context.Context, r *http.Request) apiResponse {
	vars := mux.Vars(r)
	userID, err := uuid.Parse(vars["user_id"])
	if err != nil {
		return newAPIError(models.BadRequest.Add(err))
	}

	submissions, err := h.svc.FetchForUserID(ctx, userID)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	response := models.NewListSubmissionsWithProbTitleAPIResponse(submissions)
	return response

}

func (h *submissionsHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	submission, err := h.svc.Fetch(ctx, id)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	return models.NewCreateSubmissionAPIResponse(submission) // Reusing the same Response from Create in Get

}

func (h *submissionsHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	var request models.UpdateSubmissionAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	submission := request.ToSubmissions()
	submission, err = h.svc.Update(ctx, id, submission)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewUpdateSubmissionAPIResponse(submission)
}
