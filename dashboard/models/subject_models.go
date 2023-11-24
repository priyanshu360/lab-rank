// subject.go

package models

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Subject struct
type Subject struct {
	ID           uuid.UUID       `json:"id" validate:"required"`
	Title        string          `json:"title" validate:"required"`
	Description  json.RawMessage `json:"description" validate:"required"`
	UniversityID uuid.UUID       `json:"university_id" validate:"required"`
}

// CreateSubjectAPIRequest struct
type CreateSubjectAPIRequest struct {
	Title        string          `json:"title" validate:"required"`
	Description  json.RawMessage `json:"description" validate:"required"`
	UniversityID uuid.UUID       `json:"university_id" validate:"required"`
}

// SubjectAPIResponse struct
type SubjectAPIResponse struct {
	Message *Subject
}

// Implement the Parse method for POST request for CreateSubjectAPIRequest
func (r *CreateSubjectAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for SubjectAPIResponse
func (sr *SubjectAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(sr)
}

func (r *CreateSubjectAPIRequest) ToSubject() *Subject {
	return &Subject{
		ID:           uuid.New(),
		Title:        r.Title,
		Description:  r.Description,
		UniversityID: r.UniversityID,
	}
}

func NewCreateSubjectAPIResponse(subject *Subject) *SubjectAPIResponse {
	return &SubjectAPIResponse{
		Message: subject,
	}
}

type ListSubjectsAPIResponse struct {
	Message []*Subject
}

// Implement the Write method for ListSubjectsAPIResponse
func (pr *ListSubjectsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSubjectsAPIResponse(subjects []*Subject) *ListSubjectsAPIResponse {
	return &ListSubjectsAPIResponse{
		Message: subjects,
	}
}
