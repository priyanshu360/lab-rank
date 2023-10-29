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
    ID            uuid.UUID `json:"id" validate:"required"`
    Title         string    `json:"title" validate:"required"`
    Link          string    `json:"link" validate:"required"`
    CreatedBy     uuid.UUID `json:"created_by" validate:"required"`
    CreatedAt     time.Time `json:"created_at" validate:"required"`
    UpdateEvents  json.RawMessage `json:"update_events" validate:"required"`
    LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids" validate:"required"`
}

// CreateEnvironmentAPIRequest struct
type CreateEnvironmentAPIRequest struct {
    Title        string          `json:"title" validate:"required"`
    Link         string          `json:"link" validate:"required"`
    CreatedBy    uuid.UUID       `json:"created_by" validate:"required"`
    UpdateEvents json.RawMessage `json:"update_events" validate:"required"`
    LiveDockerCIDs json.RawMessage `json:"live_dockerc_ids" validate:"required"`
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
        ID:            uuid.New(),
        Title:         r.Title,
        Link:          r.Link,
        CreatedBy:     r.CreatedBy,
        CreatedAt:     time.Now(),
        UpdateEvents:  r.UpdateEvents,
        LiveDockerCIDs: r.LiveDockerCIDs,
    }
}
