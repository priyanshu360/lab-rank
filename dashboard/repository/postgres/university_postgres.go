package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type universityPostgres struct {
	db *gorm.DB
}

// NewUniversityPostgresRepo creates a new PostgreSQL repository for universities.
func NewUniversityPostgresRepo(db *gorm.DB) *universityPostgres {
	return &universityPostgres{db}
}

// CreateUniversity creates a new university.
func (psql *universityPostgres) CreateUniversity(ctx context.Context, university models.University) models.AppError {
	result := psql.db.WithContext(ctx).Create(university)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetUniversityByID retrieves a university by its ID.
func (psql *universityPostgres) GetUniversityByID(ctx context.Context, universityID uuid.UUID) (models.University, models.AppError) {
	var university models.University
	result := psql.db.WithContext(ctx).First(&university, universityID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// University not found
			return university, models.UniversityNotFoundError
		}
		return university, models.InternalError.Add(result.Error)
	}
	return university, models.NoError
}

// Add other repository methods for universities as needed.
