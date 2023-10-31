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

func NewUserHandler(svc user_svc.UserService) *userHandler {
	return &userHandler{
		svc : svc,
	}
}


func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var response apiResponse
	var ctx = r.Context()

	switch r.Method {
	case http.MethodPost:
		response = h.handleCreate(ctx, r)
	case http.MethodGet:
		response = h.handleGet(ctx, r)
	case http.MethodPut:
		response = h.handleUpdate(ctx,r)
	case http.MethodDelete:
		response = h.handleDelete(ctx,r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}



	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}


func (h *userHandler) handleDelete(ctx context.Context, r *http.Request) apiResponse{
	id := r.URL.Query().Get("id")
	if err := h.svc.Delete(ctx, id); err != models.NoError{
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewDeleteUserAPIResponse(id)
}
func (h *userHandler) handleUpdate(ctx context.Context, r *http.Request) apiResponse {
	var request models.UpdateUserAPIRequest
	if err := request.Parse(r); err!= nil{
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}
	user, err := h.svc.Update(ctx,&request)
	if err != models.NoError{
		return newAPIError(models.BadRequest.Add(err))
	}
	return models.NewCreateUserAPIResponse(user) // Reusing the same Response for Update 
}


func (h *userHandler) handleGet(ctx context.Context, r *http.Request) apiResponse {
	email := r.URL.Query().Get("email")
    userId := r.URL.Query().Get("userID")
	request := models.GetUserAPIRequest{
        EmailID: email,
        UserID:  userId,
    }
	user,err := h.svc.Fetch(ctx,&request)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateUserAPIResponse(user) // Reusing the same Response from Create in Get 
}


func (h *userHandler) handleCreate(ctx context.Context, r *http.Request) apiResponse {
	var request models.CreateUserAPIRequest
	if err := request.Parse(r); err != nil {
		log.Println(err)
		return newAPIError(models.BadRequest.Add(err))
	}



	user := request.ToUser()
	user, err := h.svc.Create(ctx,user)
	if err != models.NoError {
		return newAPIError(models.InternalError.Add(err))
	}

	return models.NewCreateUserAPIResponse(user)
}


