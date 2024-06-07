// environment.go

package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Environment struct
type Environment struct {
	ID             int       `json:"id" validate:"required" gorm:"column:id;primaryKey;autoIncrement"`
	Title          string    `json:"title" validate:"required" gorm:"column:title"`
	Link           string    `json:"link" validate:"required" gorm:"column:link"`
	CreatedBy      uuid.UUID `json:"created_by" validate:"required" gorm:"column:created_by"`
	CreatedAt      time.Time `json:"created_at" validate:"required" gorm:"column:created_at"`
	UpdateEvents   string    `json:"update_events" validate:"required" gorm:"column:update_events"`
	LiveDockerCIDs string    `json:"live_dockerc_ids" validate:"required" gorm:"column:live_dockerc_ids"`
	File           []byte    `json:"file" gorm:"-"`
}

// CreateEnvironmentAPIRequest struct
type CreateEnvironmentAPIRequest struct {
	Title          string    `json:"title" validate:"required"`
	CreatedBy      uuid.UUID `json:"created_by" validate:"required"`
	UpdateEvents   string    `json:"update_events"`
	LiveDockerCIDs string    `json:"live_dockerc_ids"`
	File           []byte    `json:"file" validate:"required"`
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
		Title:     r.Title,
		CreatedBy: r.CreatedBy,
		CreatedAt: time.Now(),

		File: r.File,
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
