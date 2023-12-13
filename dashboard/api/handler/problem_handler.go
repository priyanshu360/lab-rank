package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
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
	h.pRouter.HandleFunc("/problem", serveHTTPWrapper(h.handleGet)).Methods("GET")
	h.pRouter.HandleFunc("/problem", serveHTTPWrapper(h.handleCreate)).Methods("POST")
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
	lang := r.URL.Query().Get("lang")
	id := r.URL.Query().Get("id")

	// todo : fix this hack
	if id != "" && lang != "" {
		uuid, _ := uuid.Parse(id)
		resp, _ := h.svc.GetInitCode(ctx, uuid, lang)
		return resp
	}

	limit := r.URL.Query().Get("limit")
	problems, err := h.svc.Fetch(ctx, id, limit)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	if len(problems) == 1 {
		return models.NewCreateProblemAPIResponse(problems[0]) // Reusing the same Response from Create in Get
	} else {
		response := models.NewListProblemsAPIResponse(problems)
		return response
	}
}

// Implement other handler methods for problems-related operations
