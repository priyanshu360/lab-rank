// university.go

package models

import (
	"encoding/json"
	"net/http"
)

// University struct
type University struct {
	ID          int    `json:"id" validate:"required" gorm:"column:id;primaryKey;autoIncrement"`
	Title       string `json:"title" validate:"required" gorm:"column:title"`
	Description string `json:"description" validate:"required" gorm:"column:description"`
}

// CreateUniversityAPIRequest struct
type CreateUniversityAPIRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
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

func (r *CreateUniversityAPIRequest) ToUniversity() *University {
	return &University{
		Title:       r.Title,
		Description: r.Description,
	}
}

func NewCreateUniversityAPIResponse(university *University) *UniversityAPIResponse {
	return &UniversityAPIResponse{
		Message: university,
	}
}

type UniversityIdName struct {
	ID    int
	Title string
}

func NewUniversityIdName(id int, name string) *UniversityIdName {
	return &UniversityIdName{
		ID:    id,
		Title: name,
	}
}

type ListUniversitiesIdNamesAPIResponse struct {
	Message []*UniversityIdName
}

type ListUniversitiesAPIResponse struct {
	Message []*University
}

// Implement the Write method for ListUniversitiesAPIResponse
func (pr *ListUniversitiesIdNamesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

// Implement the Write method for ListUniversitiesAPIResponse
func (pr *ListUniversitiesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}
func NewListUniversitiesIdNamesAPIResponse(universities []*UniversityIdName) *ListUniversitiesIdNamesAPIResponse {
	if universities == nil {
		universities = []*UniversityIdName{} // Initialize an empty slice if it's nil
	}
	return &ListUniversitiesIdNamesAPIResponse{
		Message: universities,
	}
}

func NewListUniversitiesAPIResponse(universities []*University) *ListUniversitiesAPIResponse {
	return &ListUniversitiesAPIResponse{
		Message: universities,
	}
}
