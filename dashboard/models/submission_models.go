// submissions.go

package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Accepted            Status = "Accepted"              // normal
	MemoryLimitExceeded Status = "Memory Limit Exceeded" // mle
	TimeLimitExceeded   Status = "Time Limit Exceeded"   // tle
	OutputLimitExceeded Status = "Output Limit Exceeded" // ole
	FileError           Status = "File Error"            // fe
	NonzeroExitStatus   Status = "Nonzero Exit Status"
	Signalled           Status = "Signalled"
	InternalErrorStatus Status = "Internal Error" // system error
	Queued              Status = "Queued"
	Running             Status = "Running"
)

// Submission struct
type Submission struct {
	ID           uuid.UUID               `json:"id" validate:"required"`
	SubmissionID uuid.UUID               `json:"Submission_id" validate:"required"`
	Link         string                  `json:"link" validate:"required"`
	CreatedBy    uuid.UUID               `json:"created_by" validate:"required"`
	CreatedAt    time.Time               `json:"created_at" validate:"required"`
	Score        float64                 `json:"score" validate:"required,min=0,max=100"`
	RunTime      string                  `json:"run_time" validate:"required"`
	Metadata     json.RawMessage         `json:"metadata" validate:"required"`
	Lang         ProgrammingLanguageEnum `json:"lang" validate:"required"`
	Status       Status                  `json:"status"`
	Solution     []byte                  `json:"solution" gorm:"-"`
}

// CreateSubmissionAPIRequest struct
// Todo ; change Lang to Language / add status in sql

type CreateSubmissionAPIRequest struct {
	SubmissionID uuid.UUID               `json:"Submission_id" validate:"required"`
	Solution     []byte                  `json:"solution" validate:"required"`
	CreatedBy    uuid.UUID               `json:"created_by" validate:"required"`
	Metadata     json.RawMessage         `json:"metadata"`
	Lang         ProgrammingLanguageEnum `json:"lang" validate:"required"`
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
		ID:           uuid.New(),
		SubmissionID: r.SubmissionID,
		Solution:     r.Solution,
		CreatedBy:    r.CreatedBy,
		CreatedAt:    time.Now(),
		Metadata:     r.Metadata,
		Lang:         r.Lang,
		Status:       Queued,
	}
}

func NewCreateSubmissionAPIResponse(submission *Submission) *SubmissionAPIResponse {
	return &SubmissionAPIResponse{
		Message: submission,
	}
}

type ListSubmissionsAPIResponse struct {
	Message []*Submission
}

// Implement the Write method for ListSubmissionsAPIResponse
func (pr *ListSubmissionsAPIResponse) Write(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(pr)
}

func NewListSubmissionsAPIResponse(submissions []*Submission) *ListSubmissionsAPIResponse {
	return &ListSubmissionsAPIResponse{
		Message: submissions,
	}
}
