package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

var validate = validator.New()



type APIError struct {
	code        int    
	Reason string `json:"reason"` 
}



func NewAPIError(code int, err string) APIError {
	return APIError{
		code: code,
		Reason : err,
	
	}
}

func (e APIError) Add(err error) APIError {
	fmt.Println(err)
	e.Reason = fmt.Sprintf("%s : %s", e.Reason, err.Error())
	return e
}

func (e APIError) Error() string {
	return e.Reason
}

func (e APIError) Write(w http.ResponseWriter) error {
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
var BadRequest = NewAPIError(http.StatusBadRequest, "bad request")

type DBError error

// Define specific DBError variables for different errors.
var (
    ErrUserNotFound   DBError = errors.New("user not found")
    ErrDuplicateUser  DBError = errors.New("duplicate user")
    // Add more DBError variables for other database-related errors as needed.
)

type ServiceError error