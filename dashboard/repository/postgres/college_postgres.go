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
	result := psql.db.WithContext(ctx).Table("lab_rank.college").Create(college)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetCollegeByID retrieves a college by its ID.
func (psql *collegePostgres) GetCollegeByID(ctx context.Context, collegeID uuid.UUID) (models.College, models.AppError) {
	var college models.College
	result := psql.db.WithContext(ctx).Table("lab_rank.college").First(&college, collegeID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// College not found
			return college, models.CollegeNotFoundError
		}
		return college, models.InternalError.Add(result.Error)
	}
	return college, models.NoError
}

func (psql *collegePostgres) GetCollegesListByLimit(ctx context.Context, page int, pageSize int) ([]*models.College, models.AppError) {
	var colleges []*models.College

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch colleges with the specified pagination
	result := psql.db.Offset(offset).Table("lab_rank.college").Limit(pageSize).Find(&colleges)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return colleges, models.NoError
}

// UpdateCollege updates a College's information.
func (psql *collegePostgres) UpdateCollege(ctx context.Context, collegeID uuid.UUID, college models.College) models.AppError {
	// Check if the College with the provided ID exists before updating
	var existingCollege models.College
	result := psql.db.WithContext(ctx).First(&existingCollege, collegeID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// College not found
			return models.CollegeNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the update
	result = psql.db.WithContext(ctx).Model(&college).Where("id = ?", collegeID).Updates(college)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// Add other repository methods for colleges as needed.
