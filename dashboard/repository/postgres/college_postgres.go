package postgres

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type collegePostgres struct {
	db *gorm.DB
}

// NewCollegePostgresRepo creates a new PostgreSQL repository for colleges.
func NewCollegePostgresRepo(db *gorm.DB) *collegePostgres {
	if err := db.AutoMigrate(models.College{}); err != nil {
		panic(err)
	}
	return &collegePostgres{db}
}

// CreateCollege creates a new college.
func (psql *collegePostgres) CreateCollege(ctx context.Context, college models.College) models.AppError {
	result := psql.db.WithContext(ctx).Create(&college)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetCollegeByID retrieves a college by its ID.
func (psql *collegePostgres) GetCollegeByID(ctx context.Context, collegeID int) (models.College, models.AppError) {
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

func (psql *collegePostgres) GetCollegesListByLimit(ctx context.Context, page int, pageSize int) ([]*models.College, models.AppError) {
	var colleges []*models.College

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch colleges with the specified pagination
	result := psql.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Find(&colleges)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return colleges, models.NoError
}

// Add other repository methods for colleges as needed.
func (psql *collegePostgres) GetCollegesByUniversityID(ctx context.Context, universityID int) ([]*models.CollegeIdName, models.AppError) {
	var colleges []*models.CollegeIdName

	// Fetch college names and IDs based on the university ID
	result := psql.db.WithContext(ctx).
		Select("id, title").
		Where("university_id = ?", universityID).
		Find(&colleges)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// No colleges found for the given university ID
			return nil, models.CollegeNotFoundError
		}
		return nil, models.InternalError.Add(result.Error)
	}

	return colleges, models.NoError
}
