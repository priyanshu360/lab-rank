package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	editor_svc "github.com/priyanshu360/lab-rank/dashboard/internal/editor"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type editorHandler struct {
	svc     editor_svc.Service
	eRouter *mux.Router
}

func NewEditorHandler(svc editor_svc.Service) *editorHandler {
	h := &editorHandler{
		svc:     svc,
		eRouter: mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *editorHandler) initRoutes() *editorHandler {
	h.eRouter.HandleFunc("/editor/start", serveHTTPWrapper(h.handleStart, models.AccessLevelAdmin)).Methods("POST")
	// h.eRouter.HandleFunc("/editor/submit", serveHTTPWrapper(h.handleCreate, models.AccessLevelAdmin)).Methods("POST")
	// Add other routes as needed

	return h
}

func (h *editorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.eRouter.ServeHTTP(w, r)
}
func (h *editorHandler) handleStart(ctx context.Context, r *http.Request) apiResponse {
	var request models.StartEditorReq
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	editor, err := h.svc.Start(ctx, request.ToEditor())
	if err != models.NoError {
		return newAPIError(err)
	}

	return models.NewCreateEditorAPIResponse(&editor)
}

// func (h *editorHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
// 	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
// 	environment, err := h.svc.Fetch(ctx, id)
// 	if err != models.NoError {
// 		return newAPIError(models.InternalError.Add(err))
// 	}
// 	return models.NewCreateEnvironmentAPIResponse(environment) // Reusing the same Response from Create in Get
// }
