package handler

import (
	"net/http"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	user_svc "github.com/priyanshu360/lab-rank/dashboard/service/user"
)

type userHandler struct {
	svc user_svc.UserService
}

func NewUserHandler(svc user_svc.UserService) Handler {
	return &userHandler{svc}
}


func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request HTTPRequest
	var response HTTPResponse
	switch r.Method {
	case http.MethodPost:
		request = &models.UserAPIPostRequest{}
		if err := request.Parse(r); err != nil {
			http.Error(w, "Failed to parse request", http.StatusBadRequest)
			return
		}
		response = h.handleCreate(request)

	// case http.MethodPut:
	// 	var updateRequest YourUpdateRequestStruct
	// 	if err := h.parseRequest(r, &updateRequest); err != nil {
	// 		http.Error(w, "Failed to parse request", http.StatusBadRequest)
	// 		return
	// 	}
	// 	response := h.handleUpdate(updateRequest)
	// 	if err := h.writeResponse(w, response); err != nil {
	// 		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	// 	}

	// case http.MethodDelete:
	// 	var deleteRequest YourDeleteRequestStruct
	// 	if err := h.parseRequest(r, &deleteRequest); err != nil {
	// 		http.Error(w, "Failed to parse request", http.StatusBadRequest)
	// 		return
	// 	}
	// 	response := h.handleDelete(deleteRequest)
	// 	if err := h.writeResponse(w, response); err != nil {
	// 		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	// 	}

	

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}


	if err := response.Write(w); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}


// func (h *userHandler) handleGet(r models.HTTPRequest) models.HTTPResponse {
// 	// Handle the GET request to retrieve a user's information
// 	// You can extract user ID or other relevant data from the request
// 	// Call the user service to fetch the user's data
// 	// Serialize the data and write it to the response
// }

func (h *userHandler) handleCreate(r HTTPRequest) HTTPResponse {
	// Handle the POST request to create a new user
	// Parse the request body to extract user data
	// Call the user service to create the user
	// Return a response, possibly with the newly created user's information
	return nil
}

// func (h *userHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
// 	// Handle the PUT request to update a user's information
// 	// Extract user ID and updated data from the request
// 	// Call the user service to update the user's information
// 	// Return a response, possibly with the updated user's information
// }

// func (h *userHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
// 	// Handle the DELETE request to delete a user's account
// 	// Extract user ID from the request
// 	// Call the user service to delete the user's account
// 	// Return a response indicating the success or failure of the operation
// }