package repository

import (
	"context"

	"github.com/google/uuid"
	dashboard_models "github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/judge/models"
)

type QueueRepository interface {
	Add(context.Context, models.SubmissionData) bool
	Remove(context.Context, uuid.UUID) bool
	GetNext(context.Context) models.SubmissionData
	Update(context.Context, models.SubmissionData)
	IsEmpty(context.Context) bool
}

type SubmissionRepository interface {
	GetNextBatch(context.Context, dashboard_models.Status, int) []dashboard_models.Submission
	Update(context.Context, dashboard_models.Submission)
	GetEnvFromPIDAndLang(context.Context, uuid.UUID, dashboard_models.ProgrammingLanguageEnum) dashboard_models.Environment
}
