// problems.go

package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

type TestFilesType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	File     []byte                  `json:"file" validate:"required"`
	Title    string                  `json:"title" validate:"requierd"`
}

type ProblemEnvironmentType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	Id       uuid.UUID               `json:"id" validate:"required"`
}

type EnvironmentJSON []ProblemEnvironmentType

type TestLinkType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	Link     string                  `json:"link" validate:"required"`
	Title    string                  `json:"title"`
}

type TestLinkJSON []TestLinkType

// Problem struct
type Problem struct {
	ID          uuid.UUID       `json:"id" validate:"required"`
	Title       string          `json:"title" validate:"required"`
	CreatedBy   uuid.UUID       `json:"created_by" validate:"required"`
	CreatedAt   time.Time       `json:"created_at" validate:"required"`
	Environment EnvironmentJSON `json:"environment" validate:"required"`
	ProblemLink string          `json:"problem_link" validate:"required"`
	Difficulty  DifficultyEnum  `json:"difficulty" validate:"required"`
	SyllabusID  uuid.UUID       `json:"syllabus_id" validate:"required"`
	TestLinks   TestLinkJSON    `json:"test_links" validate:"required"`
	ProblemFile []byte          `json:"problem_file" validate:"required" gorm:"-"`
	TestFiles   []TestFilesType `json:"test_files" validate:"required" gorm:"-"`
}

// CreateProblemAPIRequest struct
type CreateProblemAPIRequest struct {
	Title       string          `json:"title" validate:"required"`
	CreatedBy   uuid.UUID       `json:"created_by" validate:"required"`
	Environment EnvironmentJSON `json:"environment" validate:"required"`
	ProblemFile []byte          `json:"problem_file" validate:"required"`
	Difficulty  DifficultyEnum  `json:"difficulty" validate:"required"`
	SyllabusID  uuid.UUID       `json:"syllabus_id" validate:"required"`
	TestFiles   []TestFilesType `json:"test_files" validate:"required"`
}

// Value implements the driver.Valuer interface
func (e EnvironmentJSON) Value() (driver.Value, error) {
	return json.Marshal(e)
}

// Scan implements the sql.Scanner interface
func (e *EnvironmentJSON) Scan(value interface{}) error {
	if value == nil {
		*e = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var env []ProblemEnvironmentType
		if err := json.Unmarshal(v, &env); err != nil {
			return err
		}
		*e = EnvironmentJSON(env)
		return nil
	default:
		return errors.New("unsupported type for EnvironmentJSON")
	}
}

// Value implements the driver.Valuer interface
func (tl TestLinkJSON) Value() (driver.Value, error) {
	return json.Marshal(tl)
}

// Scan implements the sql.Scanner interface
func (tl *TestLinkJSON) Scan(value interface{}) error {
	if value == nil {
		*tl = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var testLinks []TestLinkType
		if err := json.Unmarshal(v, &testLinks); err != nil {
			return err
		}
		*tl = TestLinkJSON(testLinks)
		return nil
	default:
		return errors.New("unsupported type for TestLinkJSON")
	}
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

type ListProblemsAPIResponse struct {
	Message []*Problem
}

// Implement the Write method for ListProblemsAPIResponse
func (pr *ListProblemsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListProblemsAPIResponse(problems []*Problem) *ListProblemsAPIResponse {
	return &ListProblemsAPIResponse{
		Message: problems,
	}
}

func (r *CreateProblemAPIRequest) ToProblem() *Problem {
	return &Problem{
		ID:          uuid.New(),
		Title:       r.Title,
		CreatedBy:   r.CreatedBy,
		CreatedAt:   time.Now(),
		Environment: r.Environment,
		ProblemFile: r.ProblemFile,
		Difficulty:  r.Difficulty,
		SyllabusID:  r.SyllabusID,
		TestFiles:   r.TestFiles,
	}
}

func NewCreateProblemAPIResponse(problem *Problem) *ProblemAPIResponse {
	return &ProblemAPIResponse{
		Message: problem,
	}
}
