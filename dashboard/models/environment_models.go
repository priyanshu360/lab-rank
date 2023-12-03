// environment.go

package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// Environment struct
type Environment struct {
	ID             uuid.UUID       `json:"id" validate:"required"`
	Title          string          `json:"title" validate:"required"`
	Link           string          `json:"link" validate:"required"`
	CreatedBy      uuid.UUID       `json:"created_by" validate:"required"`
	CreatedAt      time.Time       `json:"created_at" validate:"required"`
	UpdateEvents   json.RawMessage `json:"update_events" validate:"required"`
	LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids" validate:"required" gorm:"column:live_dockerc_ids"`
	File           []byte          `json:"file" gorm:"-"`
}

// CreateEnvironmentAPIRequest struct
type CreateEnvironmentAPIRequest struct {
	Title          string          `json:"title" validate:"required"`
	CreatedBy      uuid.UUID       `json:"created_by" validate:"required"`
	UpdateEvents   json.RawMessage `json:"update_events"`
	LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids"`
	File           []byte          `json:"file" validate:"required"`
}

// EnvironmentAPIResponse struct
type EnvironmentAPIResponse struct {
	Message *Environment
}

// Implement the Parse method for POST request for CreateEnvironmentAPIRequest
func (r *CreateEnvironmentAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for EnvironmentAPIResponse
func (er *EnvironmentAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(er)
}

func (r *CreateEnvironmentAPIRequest) ToEnvironment() *Environment {
	return &Environment{
		ID:             uuid.New(),
		Title:          r.Title,
		CreatedBy:      r.CreatedBy,
		CreatedAt:      time.Now(),
		UpdateEvents:   []byte("[]"),
		LiveDockerCIDs: []byte("[]"),
		File:           r.File,
	}
}

func NewCreateEnvironmentAPIResponse(environment *Environment) *EnvironmentAPIResponse {
	return &EnvironmentAPIResponse{
		Message: environment,
	}
}

type ListEnvironmentsAPIResponse struct {
	Message []*Environment
}

// Implement the Write method for ListenvironmentsAPIResponse
func (pr *ListEnvironmentsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListEnvironmentsAPIResponse(environments []*Environment) *ListEnvironmentsAPIResponse {
	return &ListEnvironmentsAPIResponse{
		Message: environments,
	}
}

type UpdateEnvironmentAPIRequest struct {
	ID             uuid.UUID       `json:"id" validate:"required"`
	Title          string          `json:"title"`
	UpdateEvents   json.RawMessage `json:"update_events"`
	LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids"`
	File           []byte          `json:"file"`
}

func (r *UpdateEnvironmentAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

func (r *UpdateEnvironmentAPIRequest) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}

	// Todo : add custom validations
	return nil
}

func (r *UpdateEnvironmentAPIRequest) ToEnvironment(environment Environment) *Environment {
	updatedEnvironment := &Environment{
		ID:        environment.ID,
		CreatedBy: environment.CreatedBy,
		CreatedAt: environment.CreatedAt,
	}

	setField(&updatedEnvironment.Title, r.Title, environment.Title)
	setField(&updatedEnvironment.UpdateEvents, r.UpdateEvents, environment.UpdateEvents)
	setField(&updatedEnvironment.LiveDockerCIDs, r.LiveDockerCIDs, environment.LiveDockerCIDs)
	setField(&updatedEnvironment.File, r.File, environment.File)

	return updatedEnvironment
}
