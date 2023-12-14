package models

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Auth struct {
	UserID       uuid.UUID `json:"user_id"`
	AccessIDs    string    `json:"access_ids"`
	Salt         []byte    `json:"salt"`
	PasswordHash string    `json:"password_hash"`
}

type SignUpAPIRequest struct {
	CreateUserAPIRequest
	Password string `json:"password"`
}

type SignUpAPIResponse CreateUserAPIResponse

type LoginAPIResponse struct {
	Message string
}

func NewLoginAPIResponse(jwt string) *LoginAPIResponse {
	return &LoginAPIResponse{
		Message: jwt,
	}
}

func (res *LoginAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}

type LoginAPIRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (r SignUpAPIRequest) ToUser() *User {
	return &User{
		ID:           uuid.New(),
		CollegeID:    r.CollegeID,
		Status:       UserStatusActive, // You can set an initial status here if needed
		Email:        r.Email,
		ContactNo:    r.ContactNo,
		UniversityID: r.UniversityID,
		DOB:          r.DOB,
		Name:         r.Name,
		UserName:     r.UserName,
	}
}
