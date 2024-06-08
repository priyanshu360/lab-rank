// college.go

package models

import (
	"encoding/json"
	"net/http"
)

// College struct
type College struct {
	ID           int    `json:"id" validate:"required" gorm:"column:id;primaryKey;autoIncrement"`
	Title        string `json:"title" validate:"required" gorm:"column:title"`
	UniversityID int    `json:"university_id" validate:"required" gorm:"column:university_id"`
	Description  string `json:"description" validate:"required" gorm:"column:description"`
}

// CreateCollegeAPIRequest struct
type CreateCollegeAPIRequest struct {
	Title        string `json:"title" validate:"required"`
	UniversityID int    `json:"university_id" validate:"required"`
	Description  string `json:"description" validate:"required"`
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

type CollegeIdName struct {
	ID    int
	Title string
}

func NewCollegeIdName(id int, name string) *CollegeIdName {
	return &CollegeIdName{
		ID:    id,
		Title: name,
	}
}

type ListCollegesIdNamesAPIResponse struct {
	Message []*CollegeIdName
}

// Implement the Write method for ListCollegesAPIResponse
func (pr *ListCollegesIdNamesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListCollegesIdNamesAPIResponse(colleges []*CollegeIdName) *ListCollegesIdNamesAPIResponse {
	return &ListCollegesIdNamesAPIResponse{
		Message: colleges,
	}
}

type ListCollegesAPIResponse struct {
	Message []*College
}

// Implement the Write method for ListcollegesAPIResponse
func (pr *ListCollegesAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListCollegesAPIResponse(colleges []*College) *ListCollegesAPIResponse {
	return &ListCollegesAPIResponse{
		Message: colleges,
	}
}
