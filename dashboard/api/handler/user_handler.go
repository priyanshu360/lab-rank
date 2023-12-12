package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	user_svc "github.com/priyanshu360/lab-rank/dashboard/internal/user"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type userHandler struct {
	svc     user_svc.UserService
	uRouter mux.Router
}

func NewUserHandler(svc user_svc.UserService) *userHandler {
	h := &userHandler{
		svc:     svc,
		uRouter: *mux.NewRouter(),
	}

	return h.initRoutes()
}

func (h *userHandler) initRoutes() *userHandler {
	h.uRouter.HandleFunc("/user", serveHTTPWrapper(h.handleGet)).Methods("GET")
	h.uRouter.HandleFunc("/user", serveHTTPWrapper(h.handleCreate)).Methods("POST")
	h.uRouter.HandleFunc("/user", serveHTTPWrapper(h.handleUpdate)).Methods("PUT")
	h.uRouter.HandleFunc("/user", serveHTTPWrapper(h.handleDelete)).Methods("DELETE")

	return h
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.uRouter.ServeHTTP(w, r)
}

func (h *userHandler) handleDelete(ctx context.Context, r *http.Request) apiResponse {
	id := r.URL.Query().Get("id")
	if err := h.svc.Delete(ctx, id); err != models.NoError {
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewDeleteUserAPIResponse(id)
}
func (h *userHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	var request models.UpdateUserAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}
	user, err := h.svc.Update(ctx, &request)
	if err != models.NoError {
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewCreateUserAPIResponse(user) // Reusing the same Response for Update
}

func (h *userHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	email := r.URL.Query().Get("email")
	userId := r.URL.Query().Get("id")
	limit := r.URL.Query().Get("limit")
	request := models.GetUserAPIRequest{
		EmailID: email,
		UserID:  userId,
		Limit:   limit,
	}
	users, err := h.svc.Fetch(ctx, &request)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}
	switch len(users) {
	case 1:
		return models.NewCreateUserAPIResponse(users[0]) // Reusing the same Response from Create in Get
	default:
		response := models.NewFetchUserAPIResponse(users) // Create FetchUsersAPIResponse instance
		return response
	}
}

func (h *userHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateUserAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}

	user := request.ToUser()
	user, err := h.svc.Create(ctx, user)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateUserAPIResponse(user)
}
