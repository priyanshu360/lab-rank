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
	InitCode []byte                  `json:"init_code" validate:"required"`
}

type ProblemEnvironmentType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	Id       int                     `json:"id" validate:"required"`
}

// TODO : instead of list maybe have map
type EnvironmentJSON []ProblemEnvironmentType

type TestLinkType struct {
	Language ProgrammingLanguageEnum `json:"language" validate:"required"`
	Link     string                  `json:"link" validate:"required"`
	Title    string                  `json:"title" validate:"required"`
}

// TODO : instead of list maybe have map
type TestLinkJSON []TestLinkType

// Problem struct
type Problem struct {
	ID          int             `json:"id" validate:"required" gorm:"column:id;primaryKey;autoIncrement"`
	Title       string          `json:"title" validate:"required" gorm:"column:title"`
	CreatedBy   uuid.UUID       `json:"created_by" validate:"required" gorm:"column:created_by"`
	CreatedAt   time.Time       `json:"created_at" validate:"required" gorm:"column:created_at"`
	Environment EnvironmentJSON `json:"environment" validate:"required" gorm:"column:environment"`
	ProblemLink string          `json:"problem_link" validate:"required" gorm:"column:problem_link"`
	Difficulty  DifficultyEnum  `json:"difficulty" validate:"required" gorm:"column:difficulty"`
	SyllabusID  int             `json:"syllabus_id" validate:"required" gorm:"column:syllabus_id"`
	TestLinks   TestLinkJSON    `json:"test_links" validate:"required" gorm:"column:test_links"`
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
	SyllabusID  int             `json:"syllabus_id" validate:"required"`
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

func (e *ProblemEnvironmentType) Scan(value interface{}) error {
	if value == nil {
		*e = ProblemEnvironmentType{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var env []ProblemEnvironmentType
		if err := json.Unmarshal(v, &env); err != nil {
			return err
		}
		// Todo: code might break in case of empty
		*e = env[0]
		return nil
	default:
		return errors.New("unsupported type for ProblemEnvironmentType")
	}
}

// Value implements the driver.Valuer interface
func (tl ProblemEnvironmentType) Value() (driver.Value, error) {
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
	return json.NewEncoder(w).Encode(pr)
}

type ListProblemsAPIResponse struct {
	Message []*Problem
}

// Implement the Write method for ListProblemsAPIResponse
func (pr *ListProblemsAPIResponse) Write(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(pr)
}

func NewListProblemsAPIResponse(problems []*Problem) *ListProblemsAPIResponse {
	return &ListProblemsAPIResponse{
		Message: problems,
	}
}

// todo : rename
type InitProblemCode struct {
	Message []byte
}

// Implement the Write method for ListProblemsAPIResponse
func (pr *InitProblemCode) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewInitProblemCode(code []byte) *InitProblemCode {
	return &InitProblemCode{
		Message: code,
	}
}

func (r *CreateProblemAPIRequest) ToProblem() *Problem {
	return &Problem{
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

func (e *TestLinkType) Scan(value interface{}) error {
	if value == nil {
		*e = TestLinkType{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		var t []TestLinkType
		if err := json.Unmarshal(v, &t); err != nil {
			return err
		}
		// Todo: code might break in case of empty
		*e = t[0]
		return nil
	default:
		return errors.New("unsupported type for TestLinkType")
	}
}

// Value implements the driver.Valuer interface
func (tl TestLinkType) Value() (driver.Value, error) {
	return json.Marshal(tl)
}
