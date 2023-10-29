// problems.go

package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type DifficultyEnum string

const (
    DifficultyEasy   DifficultyEnum = "EASY"
    DifficultyMedium DifficultyEnum = "MEDIUM"
    DifficultyHard   DifficultyEnum = "HARD"
)

// Problem struct
type Problem struct {
    ID           uuid.UUID `json:"id" validate:"required"`
    Title        string    `json:"title" validate:"required"`
    CreatedBy    uuid.UUID `json:"created_by" validate:"required"`
    CreatedAt    time.Time `json:"created_at" validate:"required"`
    Environment  json.RawMessage `json:"environment" validate:"required"`
    ProblemLink  string    `json:"problem_link" validate:"required"`
    Difficulty   DifficultyEnum `json:"difficulty" validate:"required"`
    SyllabusID   uuid.UUID `json:"syllabus_id" validate:"required"`
    TestLink     string    `json:"test_link" validate:"required"`
}

// CreateProblemAPIRequest struct
type CreateProblemAPIRequest struct {
    Title       string    `json:"title" validate:"required"`
    CreatedBy   uuid.UUID `json:"created_by" validate:"required"`
    Environment json.RawMessage `json:"environment" validate:"required"`
    ProblemLink string    `json:"problem_link" validate:"required"`
    Difficulty  DifficultyEnum `json:"difficulty" validate:"required"`
    SyllabusID  uuid.UUID `json:"syllabus_id" validate:"required"`
    TestLink    string    `json:"test_link" validate:"required"`
}

// ProblemAPIResponse struct
type ProblemAPIResponse struct {
    Message *Problem
}

// Implement the Parse method for POST request for CreateProblemAPIRequest
func (r *CreateProblemAPIRequest) Parse(req *http.Request) error {
    if err := json.NewDecoder(req.Body).Decode(r); err != nil {
        return err
    }
    return validate.Struct(r)
}

// Implement the Write method for ProblemAPIResponse
func (pr *ProblemAPIResponse) Write(w http.ResponseWriter) error {
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(pr)
}

func (r *CreateProblemAPIRequest) ToProblem() *Problem {
    return &Problem{
        ID:           uuid.New(),
        Title:        r.Title,
        CreatedBy:    r.CreatedBy,
        CreatedAt:    time.Now(),
        Environment:  r.Environment,
        ProblemLink:  r.ProblemLink,
        Difficulty:   r.Difficulty,
        SyllabusID:   r.SyllabusID,
        TestLink:     r.TestLink,
    }
}
