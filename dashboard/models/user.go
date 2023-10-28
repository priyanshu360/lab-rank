package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
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
	ID           uuid.UUID  `json:"id" validate:"required"`
	CollegeID    uuid.UUID  `json:"college_id" validate:"required"`
	Status       UserStatus `json:"status" validate:"required,oneof=ACTIVE INACTIVE DELETED SPAM"`
	CreatedAt    time.Time  `json:"created_at" validate:"required"`
	Email        string     `json:"email" validate:"required,email"`
	ContactNo    string     `json:"contact_no" validate:"required,len=10"`
	UniversityID string     `json:"university_id"`
	DOB          time.Time  `json:"dob" validate:"required"`
}

type CreateUserAPIRequest struct {
	CollegeID    uuid.UUID `json:"college_id" validate:"required"`
	AccessID     uuid.UUID `json:"access_id" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	ContactNo    string    `json:"contact_no" validate:"required"`
	DOB          time.Time `json:"dob" validate:"required"`
	UniversityID string		`json:"university_id"`
}

func (r *CreateUserAPIRequest) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}

	// Todo : add custom validations
	return nil
}

func (r *User) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}

	// Todo : add custom validations
	return nil
}

func (r *User) GenerateReqFields() error {
	r.ID = uuid.New()
	r.Status = UserStatusInactive
	r.CreatedAt = time.Now()
	return r.validate()
}

func (r *CreateUserAPIRequest) ToUser() *User {
	return &User{
		ID:           uuid.New(),
		CollegeID:    r.CollegeID,
		Status:       UserStatusActive, // You can set an initial status here if needed
		Email:        r.Email,
		ContactNo:    r.ContactNo,
		UniversityID: r.UniversityID,
		DOB:          r.DOB,
	}
}

// Implement the Parse method for POST request
func (r *CreateUserAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

type CreateUserapiResponse struct {
	Message *User
}


// Implement the Write method for UserapiResponse
func (cr *CreateUserapiResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(cr)
}
