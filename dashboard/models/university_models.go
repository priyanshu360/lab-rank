// university.go

package models

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// University struct
type University struct {
    ID          uuid.UUID `json:"id" validate:"required"`
    Title       string    `json:"title" validate:"required"`
    Description json.RawMessage `json:"description" validate:"required"`
}

// CreateUniversityAPIRequest struct
type CreateUniversityAPIRequest struct {
    Title       string        `json:"title" validate:"required"`
    Description json.RawMessage `json:"description" validate:"required"`
}

// UniversityAPIResponse struct
type UniversityAPIResponse struct {
    Message *University
}

// Implement the Parse method for POST request for CreateUniversityAPIRequest
func (r *CreateUniversityAPIRequest) Parse(req *http.Request) error {
    if err := json.NewDecoder(req.Body).Decode(r); err != nil {
        return err
    }
    return validate.Struct(r)
}

// Implement the Write method for UniversityAPIResponse
func (ur *UniversityAPIResponse) Write(w http.ResponseWriter) error {
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(ur)
}


func (r *CreateUniversityAPIRequest) ToUniversity() University {
    return University{
        ID:          uuid.New(),
        Title:       r.Title,
        Description: r.Description,
    }
}

func NewCreateUniversityAPIResponse(university *University) *UniversityAPIResponse {
    return &UniversityAPIResponse{
        Message: university,
    }
}