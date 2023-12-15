// syllabus.go

package models

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type SyllabusLevelEnum string

const (
	SyllabusLevelUniversity SyllabusLevelEnum = "UNIVERSITY"
	SyllabusLevelCollege    SyllabusLevelEnum = "COLLEGE"
	SyllabusLevelGlobal     SyllabusLevelEnum = "GLOBAL"
)

type AccessLevel struct {
	ID         uuid.UUID           `gorm:"type:uuid;primaryKey" json:"id"`
	Mode       AccessLevelModeEnum `gorm:"type:access_level_mode_enum;not null" json:"mode"`
	SyllabusID uuid.UUID           `gorm:"type:uuid;not null;index" json:"syllabus_id"`
}

// AccessLevelModeEnum represents the lab_rank.access_level_mode_enum enum type.
type AccessLevelModeEnum string

const (
	AccessLevelAdmin   AccessLevelModeEnum = "ADMIN"
	AccessLevelTeacher AccessLevelModeEnum = "TEACHER"
	AccessLevelStudent AccessLevelModeEnum = "STUDENT"
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

func (r Subject) ToSyllabus(id uuid.UUID, level SyllabusLevelEnum) *Syllabus {
	return &Syllabus{
		ID:            uuid.New(),
		SubjectID:     r.ID,
		UniCollegeID:  id,
		SyllabusLevel: level,
	}
}

func (r Syllabus) ToAccessLevel(mode AccessLevelModeEnum) *AccessLevel {
	return &AccessLevel{
		ID:         uuid.New(),
		Mode:       mode,
		SyllabusID: r.ID,
	}
}
