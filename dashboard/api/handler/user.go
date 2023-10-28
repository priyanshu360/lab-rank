package handler

import (
	"context"
	"log"
	"net/http"

	user_svc "github.com/priyanshu360/lab-rank/dashboard/internal/user"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type userHandler struct {
	svc user_svc.UserService
}

func NewUserHandler(svc user_svc.UserService) Handler {
	return &userHandler{
		svc : svc,
	}
}

type HandleFunc func(*apiRequest) apiResponse

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response apiResponse
	var ctx = r.Context()

	switch r.Method {
	case http.MethodPost:
		response = h.handleCreate(ctx, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}



	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}





func (h *userHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateUserAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return models.BadRequest.Add(err)
	}



	user := request.ToUser()
	response, err := h.svc.Create(ctx,user)
	if err != nil {
		return models.CustomError.Add(err)
	}

	return &models.CreateUserapiResponse{
		Message: response,
	}
}

