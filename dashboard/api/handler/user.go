package handler

import (
	"fmt"
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


func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.HTTPRequest
	var response models.HTTPResponse
	switch r.Method {
	case http.MethodPost:
		request = &models.CreateUserAPIRequest{}
		log.Println("hellow ", request)
		if err := request.Parse(r); err != nil {
			log.Println(err)
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return
		}
		response = h.handleCreate(request)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}


	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}




func (h *userHandler) handleCreate(r models.HTTPRequest) models.HTTPResponse {
	// Handle the POST request to create a new user
	// Parse the request body to extract user data
	// Call the user service to create the user
	// Return a response, possibly with the newly created user's information
	return models.CustomError.Error(fmt.Errorf("Custom Error"))
}

