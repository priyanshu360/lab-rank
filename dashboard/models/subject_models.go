// subject.go

package models

import (
	"encoding/json"
	"net/http"
)

// Subject struct
type Subject struct {
	ID           int    `json:"id" validate:"required" gorm:"column:id;primaryKey;autoIncrement"`
	Title        string `json:"title" validate:"required" gorm:"column:title"`
	Description  string `json:"description" validate:"required" gorm:"column:description"`
	UniversityID int    `json:"university_id" validate:"required" gorm:"column:university_id"`
}

// CreateSubjectAPIRequest struct
type CreateSubjectAPIRequest struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description" validate:"required"`
	UniversityID int    `json:"university_id" validate:"required"`
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
