package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type subjectPostgres struct {
	db *gorm.DB
}

// NewSubjectPostgresRepo creates a new PostgreSQL repository for subjects.
func NewSubjectPostgresRepo(db *gorm.DB) *subjectPostgres {
	return &subjectPostgres{db}
}

// CreateSubject creates a new subject.
func (psql *subjectPostgres) CreateSubject(ctx context.Context, subject models.Subject) models.AppError {
	result := psql.db.WithContext(ctx).Table("lab_rank.subject").Create(subject)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetSubjectByID retrieves a subject by its ID.
func (psql *subjectPostgres) GetSubjectByID(ctx context.Context, subjectID uuid.UUID) (models.Subject, models.AppError) {
	var subject models.Subject
	result := psql.db.WithContext(ctx).First(&subject, subjectID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Subject not found
			return subject, models.SubjectNotFoundError
		}
		return subject, models.InternalError.Add(result.Error)
	}
	return subject, models.NoError
}

// Add other repository methods for subjects as needed.
