package models

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	UserID       uuid.UUID `json:"user_id"`
	AccessIDs    string    `json:"access_ids"`
	Salt         string    `json:"salt"`
	PasswordHash string    `json:"password_hash"`
}

type SignUpAPIRequest struct {
	CollegeID    uuid.UUID `json:"college_id" validate:"required"`
	AccessID     uuid.UUID `json:"access_id" validate:"required"` // To do: Implement AccessID functionality
	Email        string    `json:"email" validate:"required,email"`
	ContactNo    string    `json:"contact_no" validate:"required"`
	DOB          time.Time `json:"dob" validate:"required"`
	UniversityID string    `json:"university_id"`
	Password     string    `json:"password"`
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
		// Name:         r.Name,
	}
}
