package models

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserStatus string

const (
    UserStatusActive   UserStatus = "ACTIVE"
    UserStatusInactive UserStatus = "INACTIVE"
    UserStatusDeleted  UserStatus = "DELETED"
    UserStatusSpam     UserStatus = "SPAM"
)

type User struct {
    ID         uuid.UUID `json:"id" validate:"required"`
    CollegeID  uuid.UUID `json:"college_id" validate:"required"`
    Status     UserStatus `json:"status" validate:"required,oneof=ACTIVE INACTIVE DELETED SPAM"`
    CreatedAt  time.Time `json:"created_at" validate:"required"`
    Email      string    `json:"email" validate:"required,email"`
    ContactNo  string    `json:"contact_no" validate:"required,len=15"`
    UniversityID uuid.UUID `json:"university_id"`
    DOB        time.Time `json:"dob" validate:"required"`
}



type CreateUserAPIRequest struct {
    CollegeID   uuid.UUID `json:"college_id" validate:"required"`
    Status      string    `json:"status" validate:"required"`
    AccessID    uuid.UUID `json:"access_id" validate:"required"`
    Email       string    `json:"email" validate:"required,email"`
    ContactNo   string    `json:"contact_no" validate:"required,contact_number"`
    UniversityID string    `json:"university_id"`
    DOB         time.Time `json:"dob" validate:"required"`
}

// UserAPIResponse for all request methods
type CreateUserAPIResponse struct {
	Message *User 
}



// Implement the Parse method for POST request
func (r *CreateUserAPIRequest) Parse(req *http.Request) error {
	// Implement parsing logic for the POST request
	// Extract necessary data from the request and populate r
	log.Println("hello",  r, req)
	return json.NewDecoder(req.Body).Decode(r)
}


// Implement the Write method for UserAPIResponse
func (r *CreateUserAPIResponse) Write(w http.ResponseWriter) error {
	// Implement serialization and writing logic for the User API response
	// Serialize the struct r and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}

