package models

import (
	"github.com/google/uuid"
	dashboard_models "github.com/priyanshu360/lab-rank/dashboard/models"
)

type SubmissionQueue []SubmissionData

type SubmissionData struct {
	ID           uuid.UUID
	SubmissionId uuid.UUID
	ProblemID    uuid.UUID
	Link         string
	Status       dashboard_models.Status
	Score        float64
	RunTime      string
	Language     dashboard_models.ProgrammingLanguageEnum
}

func NewSubmissionData(submission dashboard_models.Submission) SubmissionData {
	return SubmissionData{
		ID:           uuid.New(),
		SubmissionId: submission.ID,
		ProblemID:    submission.ProblemID,
		Status:       dashboard_models.Queued,
		Language:     submission.Language,
	}
}
