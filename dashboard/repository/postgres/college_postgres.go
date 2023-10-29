package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type collegePostgres struct {
	db *gorm.DB
}

// NewCollegePostgresRepo creates a new PostgreSQL repository for colleges.
func NewCollegePostgresRepo(db *gorm.DB) *collegePostgres {
	return &collegePostgres{db}
}

// CreateCollege creates a new college.
func (psql *collegePostgres) CreateCollege(ctx context.Context, college models.College) models.AppError {
	result := psql.db.WithContext(ctx).Create(college)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetCollegeByID retrieves a college by its ID.
func (psql *collegePostgres) GetCollegeByID(ctx context.Context, collegeID uuid.UUID) (models.College, models.AppError) {
	var college models.College
	result := psql.db.WithContext(ctx).First(&college, collegeID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// College not found
			return college, models.CollegeNotFoundError
		}
		return college, models.InternalError.Add(result.Error)
	}
	return college, models.NoError
}

// Add other repository methods for colleges as needed.
