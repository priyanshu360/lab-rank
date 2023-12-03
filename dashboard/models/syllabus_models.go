// syllabus.go

package models

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type SyllabusLevelEnum string

const (
	SyllabusLevelUniversity SyllabusLevelEnum = "UNIVERSITY"
	SyllabusLevelCollege    SyllabusLevelEnum = "COLLEGE"
	SyllabusLevelGlobal     SyllabusLevelEnum = "GLOBAL"
)

// Syllabus struct
type Syllabus struct {
	ID            uuid.UUID         `json:"id" validate:"required"`
	SubjectID     uuid.UUID         `json:"subject_id" validate:"required"`
	UniCollegeID  uuid.UUID         `json:"uni_college_id" validate:"required"`
	SyllabusLevel SyllabusLevelEnum `json:"syllabus_level" validate:"required"`
}

// CreateSyllabusAPIRequest struct
type CreateSyllabusAPIRequest struct {
	SubjectID     uuid.UUID         `json:"subject_id" validate:"required"`
	UniCollegeID  uuid.UUID         `json:"uni_college_id" validate:"required"`
	SyllabusLevel SyllabusLevelEnum `json:"syllabus_level" validate:"required"`
}

// SyllabusAPIResponse struct
type SyllabusAPIResponse struct {
	Message *Syllabus
}

// Implement the Parse method for POST request for CreateSyllabusAPIRequest
func (r *CreateSyllabusAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// Implement the Write method for SyllabusAPIResponse
func (sr *SyllabusAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(sr)
}

func (r *CreateSyllabusAPIRequest) ToSyllabus() *Syllabus {
	return &Syllabus{
		ID:            uuid.New(),
		SubjectID:     r.SubjectID,
		UniCollegeID:  r.UniCollegeID,
		SyllabusLevel: r.SyllabusLevel,
	}
}

func NewCreateSyllabusAPIResponse(syllabus *Syllabus) *SyllabusAPIResponse {
	return &SyllabusAPIResponse{
		Message: syllabus,
	}
}

type ListSyllabusAPIResponse struct {
	Message []*Syllabus
}

// Implement the Write method for ListSyllabussAPIResponse
func (pr *ListSyllabusAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSyllabusAPIResponse(syllabus []*Syllabus) *ListSyllabusAPIResponse {
	return &ListSyllabusAPIResponse{
		Message: syllabus,
	}
}

type UpdateSyllabusAPIRequest struct {
	ID            uuid.UUID         `json:"id" validate:"required"`
	SubjectID     uuid.UUID         `json:"subject_id"`
	UniCollegeID  uuid.UUID         `json:"uni_college_id"`
	SyllabusLevel SyllabusLevelEnum `json:"syllabus_level"` // validate:"oneof= UNIVERSITY COLLEGE GLOBAL`
}

func (r *UpdateSyllabusAPIRequest) Parse(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return err
	}
	return r.validate()
}

func (r *UpdateSyllabusAPIRequest) validate() error {
	if err := validate.Struct(r); err != nil {
		return err.(validator.ValidationErrors)
	}

	// Todo : add custom validations
	return nil
}

func (r *UpdateSyllabusAPIRequest) ToSyllabus(syllabus Syllabus) *Syllabus {
	newSyllabus := &Syllabus{
		ID: syllabus.ID,
	}

	setField(&newSyllabus.SubjectID, r.SubjectID, syllabus.SubjectID)
	setField(&newSyllabus.UniCollegeID, r.UniCollegeID, syllabus.UniCollegeID)
	setField(&newSyllabus.SyllabusLevel, r.SyllabusLevel, syllabus.SyllabusLevel)

	return newSyllabus
}
