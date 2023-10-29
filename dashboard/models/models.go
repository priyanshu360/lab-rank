package models

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()



type AppError struct {
	Type        ErrorType    
	Reason string `json:"reason"` 
}



func NewAppError(code ErrorType, err string) AppError {
	return AppError{
		Type: code,
		Reason : err,
	
	}
}

func (e AppError) Add(err error) AppError {
	fmt.Println(err)
	e.Reason = fmt.Sprintf("%s : %s", e.Reason, err.Error())
	return e
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s, %s", e.Type, e.Reason)
}



var (
	UserNotFoundError = NewAppError(ErrorNotFound, "user not found")
	UserInvalidInput = NewAppError(ErrorBadData, "invalid input")
	InternalError = NewAppError(ErrorInternal, "internal server error")
	BadRequest = NewAppError(ErrorBadData, "bad request")
	NoError = NewAppError(ErrorNone, "")
)




type ErrorType string


const (
	ErrorNone          ErrorType = ""
	ErrorTimeout       ErrorType = "timeout"
	ErrorCanceled      ErrorType = "canceled"
	ErrorExec          ErrorType = "execution"
	ErrorBadData       ErrorType = "bad_data"
	ErrorInternal      ErrorType = "internal"
	ErrorUnavailable   ErrorType = "unavailable"
	ErrorNotFound      ErrorType = "not_found"
	ErrorNotAcceptable ErrorType = "not_acceptable"
)