// college.go

package models

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// College struct
type College struct {
    ID           uuid.UUID `json:"id" validate:"required"`
    Title        string    `json:"title" validate:"required"`
    UniversityID uuid.UUID `json:"university_id" validate:"required"`
    Description  json.RawMessage `json:"description" validate:"required"`
}

// CreateCollegeAPIRequest struct
type CreateCollegeAPIRequest struct {
    Title       string        `json:"title" validate:"required"`
    UniversityID uuid.UUID    `json:"university_id" validate:"required"`
    Description json.RawMessage `json:"description" validate:"required"`
}

// CollegeAPIResponse struct
type CollegeAPIResponse struct {
    Message *College
}

// Implement the Parse method for POST request for CreateCollegeAPIRequest
func (r *CreateCollegeAPIRequest) Parse(req *http.Request) error {
    if err := json.NewDecoder(req.Body).Decode(r); err != nil {
        return err
    }
    return validate.Struct(r)
}

// Implement the Write method for CollegeAPIResponse
func (cr *CollegeAPIResponse) Write(w http.ResponseWriter) error {
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(cr)
}

func (r *CreateCollegeAPIRequest) ToCollege() *College {
    return &College{
        ID:           uuid.New(),
        Title:        r.Title,
        UniversityID: r.UniversityID,
        Description:  r.Description,
    }
}

func NewCreateCollegeAPIResponse(college *College) *CollegeAPIResponse {
    return &CollegeAPIResponse{
        Message: college,
    }
}