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
	result := psql.db.WithContext(ctx).Table("lab_rank.syllabus").Create(syllabus)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetSyllabusByID retrieves a syllabus by its ID.
func (psql *syllabusPostgres) GetSyllabusByID(ctx context.Context, syllabusID uuid.UUID) (models.Syllabus, models.AppError) {
	var syllabus models.Syllabus
	result := psql.db.WithContext(ctx).Table("lab_rank.syllabus").First(&syllabus, syllabusID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Syllabus not found
			return syllabus, models.SyllabusNotFoundError
		}
		return syllabus, models.InternalError.Add(result.Error)
	}
	return syllabus, models.NoError
}

func (psql *syllabusPostgres) GetSyllabusListByLimit(ctx context.Context, page int, pageSize int) ([]*models.Syllabus, models.AppError) {
	var syllabuss []*models.Syllabus

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch syllabuss with the specified pagination
	result := psql.db.Offset(offset).Table("lab_rank.syllabus").Limit(pageSize).Find(&syllabuss)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return syllabuss, models.NoError
}

// UpdateSyllabus updates a Syllabus's information.
func (psql *syllabusPostgres) UpdateSyllabus(ctx context.Context, syllabusID uuid.UUID, syllabus models.Syllabus) models.AppError {
	// Check if the Syllabus with the provided ID exists before updating
	var existingSyllabus models.Syllabus
	result := psql.db.WithContext(ctx).First(&existingSyllabus, syllabusID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Syllabus not found
			return models.SyllabusNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the update
	result = psql.db.WithContext(ctx).Model(&syllabus).Where("id = ?", syllabusID).Updates(syllabus)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// Add other repository methods for syllabus as needed.
