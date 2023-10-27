package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type APIError struct {
	code        int    
	Reason string `json:"reason"` 
}

func NewAPIError(code int, err string) *APIError {
	return &APIError{
		code: code,
		Reason : err,
	
	}
}

func (e *APIError) Error(err error) *APIError {
	e.Reason = fmt.Sprintf("%s : %s", e.Reason, err)
	return e
}

func (e *APIError) Write(w http.ResponseWriter) error {
	// Implement serialization and writing logic for the User API response
	// Serialize the struct r and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.code)
	return json.NewEncoder(w).Encode(e)
}


var (
	UserNotFoundError = NewAPIError(http.StatusNotFound, "user not found")
	UserInvalidInput = NewAPIError(http.StatusBadRequest, "invalid input")
)

var CustomError = NewAPIError(http.StatusInternalServerError, "internal server error")