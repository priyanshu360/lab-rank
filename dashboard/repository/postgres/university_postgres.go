package postgres

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type universityPostgres struct {
	db *gorm.DB
}

// NewUniversityPostgresRepo creates a new PostgreSQL repository for universities.
func NewUniversityPostgresRepo(db *gorm.DB) *universityPostgres {
	if err := db.AutoMigrate(models.University{}); err != nil {
		panic(err)
	}
	return &universityPostgres{db}
}

// CreateUniversity creates a new university.
func (psql *universityPostgres) CreateUniversity(ctx context.Context, university *models.University) models.AppError {
	result := psql.db.WithContext(ctx).Create(university)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetUniversityByID retrieves a university by its ID.
func (psql *universityPostgres) GetUniversityByID(ctx context.Context, universityID int) (models.University, models.AppError) {
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

func (psql *universityPostgres) GetUniversitiesListByLimit(ctx context.Context, page int, pageSize int) ([]*models.University, models.AppError) {
	var universities []*models.University

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch universities with the specified pagination
	result := psql.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&universities)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return universities, models.NoError
}

// Add other repository methods for universities as needed.
func (psql *universityPostgres) GetAllUniversityNames(ctx context.Context) ([]*models.UniversityIdName, models.AppError) {
	var universities []*models.UniversityIdName

	// Fetch all university names with their IDs
	result := psql.db.WithContext(ctx).
		Select("id, title").
		Find(&universities)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return universities, models.NoError
}
