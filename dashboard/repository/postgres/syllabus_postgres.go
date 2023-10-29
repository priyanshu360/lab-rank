package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type syllabusPostgres struct {
	db *gorm.DB
}

// NewSyllabusPostgresRepo creates a new PostgreSQL repository for syllabi.
func NewSyllabusPostgresRepo(db *gorm.DB) *syllabusPostgres {
	return &syllabusPostgres{db}
}

// CreateSyllabus creates a new syllabus.
func (psql *syllabusPostgres) CreateSyllabus(ctx context.Context, syllabus models.Syllabus) models.AppError {
	result := psql.db.WithContext(ctx).Create(syllabus)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetSyllabusByID retrieves a syllabus by its ID.
func (psql *syllabusPostgres) GetSyllabusByID(ctx context.Context, syllabusID uuid.UUID) (models.Syllabus, models.AppError) {
	var syllabus models.Syllabus
	result := psql.db.WithContext(ctx).First(&syllabus, syllabusID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Syllabus not found
			return syllabus, models.SyllabusNotFoundError
		}
		return syllabus, models.InternalError.Add(result.Error)
	}
	return syllabus, models.NoError
}

// Add other repository methods for syllabi as needed.
