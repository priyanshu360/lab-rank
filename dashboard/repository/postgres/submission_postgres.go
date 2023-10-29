package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type submissionPostgres struct {
	db *gorm.DB
}

// NewSubmissionPostgresRepo creates a new PostgreSQL repository for submissions.
func NewSubmissionPostgresRepo(db *gorm.DB) *submissionPostgres {
	return &submissionPostgres{db}
}

// CreateSubmission creates a new submission.
func (psql *submissionPostgres) CreateSubmission(ctx context.Context, submission models.Submission) models.AppError {
	result := psql.db.WithContext(ctx).Create(submission)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetSubmissionByID retrieves a submission by its ID.
func (psql *submissionPostgres) GetSubmissionByID(ctx context.Context, submissionID uuid.UUID) (models.Submission, models.AppError) {
	var submission models.Submission
	result := psql.db.WithContext(ctx).First(&submission, submissionID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Submission not found
			return submission, models.SubmissionNotFoundError
		}
		return submission, models.InternalError.Add(result.Error)
	}
	return submission, models.NoError
}

// Add other repository methods for submissions as needed.
