// submissions.go

package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)


type ProgrammingLanguageEnum string

const (
    ProgrammingLanguageC        ProgrammingLanguageEnum = "C"
    ProgrammingLanguageCPlusPlus ProgrammingLanguageEnum = "C++"
    ProgrammingLanguageJava     ProgrammingLanguageEnum = "Java"
    ProgrammingLanguagePython   ProgrammingLanguageEnum = "Python"
    ProgrammingLanguageJavaScript ProgrammingLanguageEnum = "JavaScript"
    ProgrammingLanguageGo       ProgrammingLanguageEnum = "Go"
    ProgrammingLanguageRust     ProgrammingLanguageEnum = "Rust"
    // Add more programming languages as needed
)


// Submission struct
type Submission struct {
    ID           uuid.UUID `json:"id" validate:"required"`
    ProblemID    uuid.UUID `json:"problem_id" validate:"required"`
    Link         string    `json:"link" validate:"required"`
    CreatedBy    uuid.UUID `json:"created_by" validate:"required"`
    CreatedAt    time.Time `json:"created_at" validate:"required"`
    Score        float64   `json:"score" validate:"required,min=0,max=100"`
    RunTime      string    `json:"run_time" validate:"required"`
    Metadata     json.RawMessage `json:"metadata" validate:"required"`
    Language     ProgrammingLanguageEnum `json:"lang" validate:"required"`
}

// CreateSubmissionAPIRequest struct
type CreateSubmissionAPIRequest struct {
    ProblemID uuid.UUID `json:"problem_id" validate:"required"`
    Link      string    `json:"link" validate:"required"`
    CreatedBy uuid.UUID `json:"created_by" validate:"required"`
    Score     float64   `json:"score" validate:"required,min=0,max=100"`
    RunTime   string    `json:"run_time" validate:"required"`
    Metadata  json.RawMessage `json:"metadata" validate:"required"`
    Language  ProgrammingLanguageEnum `json:"lang" validate:"required"`
}

// SubmissionAPIResponse struct
type SubmissionAPIResponse struct {
    Message *Submission
}

// Implement the Parse method for POST request for CreateSubmissionAPIRequest
func (r *CreateSubmissionAPIRequest) Parse(req *http.Request) error {
    if err := json.NewDecoder(req.Body).Decode(r); err != nil {
        return err
    }
    return validate.Struct(r)
}

// Implement the Write method for SubmissionAPIResponse
func (sr *SubmissionAPIResponse) Write(w http.ResponseWriter) error {
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(sr)
}


func (r *CreateSubmissionAPIRequest) ToSubmissions() *Submission {
    return &Submission{
        ID:          uuid.New(),
        ProblemID:   r.ProblemID,
        Link:        r.Link,
        CreatedBy:   r.CreatedBy,
        CreatedAt:   time.Now(),
        Score:       r.Score,
        RunTime:     r.RunTime,
        Metadata:    r.Metadata,
        Language:        r.Language,
    }
}
