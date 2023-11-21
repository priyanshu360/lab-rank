package models

import (
	"encoding/json"
	"net/http"
	"reflect"
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
	// Name         string
}

type CreateUserAPIRequest struct {
	CollegeID    uuid.UUID `json:"college_id" validate:"required"`
	AccessID     uuid.UUID `json:"access_id" validate:"required"` // To do: Implement AccessID functionality
	Email        string    `json:"email" validate:"required,email"`
	ContactNo    string    `json:"contact_no" validate:"required"`
	DOB          time.Time `json:"dob" validate:"required"`
	UniversityID string    `json:"university_id"`
	// Name         string    `json:"name"`
}

type GetUserAPIRequest struct {
	UserID  string `json:"user_id"`
	EmailID string `json:"email_id"`
}

type UpdateUserAPIRequest struct {
	ID        uuid.UUID  `json:"id" validate:"required"`
	Status    UserStatus `json:"status" validate:"oneof=ACTIVE INACTIVE DELETED SPAM"`
	Email     string     `json:"email" validate:"email"`
	ContactNo string     `json:"contact_no"`
}

func (r *UpdateUserAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

func (r *UpdateUserAPIRequest) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}

	// Todo : add custom validations
	return nil
}

func (r *UpdateUserAPIRequest) ToUser(user User) *User {
	newUser := &User{
		ID:           user.ID,
		CollegeID:    user.CollegeID,
		UniversityID: user.UniversityID,
		DOB:          user.DOB,
	}

	setField(&newUser.Email, r.Email, user.Email)
	setField(&newUser.ContactNo, r.ContactNo, user.ContactNo)
	setField(&newUser.Status, r.Status, user.Status)

	return newUser
}

func setField(field interface{}, value, defaultValue interface{}) {
	fieldValue := reflect.ValueOf(field).Elem()
	if value != nil {
		fieldValue.Set(reflect.ValueOf(value))
	} else {
		fieldValue.Set(reflect.ValueOf(defaultValue))
	}
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
		// Name:         r.Name,
	}
}

// Implement the Parse method for POST request
func (r *CreateUserAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

type CreateUserAPIResponse struct {
	Message *User
}

// Implement the Write method for UserapiResponse
func (cr *CreateUserAPIResponse) Write(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(cr)
}

func NewCreateUserAPIResponse(user *User) *CreateUserAPIResponse {
	return &CreateUserAPIResponse{
		Message: user,
	}
}

type DeleteUserAPIResponse struct {
	user_id string
}

// Implement the Write method for UserapiResponse
func (cr *DeleteUserAPIResponse) Write(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(cr)
}

func NewDeleteUserAPIResponse(userId string) *DeleteUserAPIResponse {
	return &DeleteUserAPIResponse{
		user_id: userId,
	}
}
