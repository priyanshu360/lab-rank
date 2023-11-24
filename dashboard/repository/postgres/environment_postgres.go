package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type environmentPostgres struct {
	db *gorm.DB
}

// NewEnvironmentPostgresRepo creates a new PostgreSQL repository for environments.
func NewEnvironmentPostgresRepo(db *gorm.DB) *environmentPostgres {
	return &environmentPostgres{db}
}

// CreateEnvironment creates a new environment.
func (psql *environmentPostgres) CreateEnvironment(ctx context.Context, environment models.Environment) models.AppError {
	result := psql.db.WithContext(ctx).Table("lab_rank.environment").Create(environment)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetEnvironmentByID retrieves an environment by its ID.
func (psql *environmentPostgres) GetEnvironmentByID(ctx context.Context, environmentID uuid.UUID) (models.Environment, models.AppError) {
	var environment models.Environment
	result := psql.db.WithContext(ctx).Table("lab_rank.environment").First(&environment, environmentID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Environment not found
			return environment, models.EnvironmentNotFoundError
		}
		return environment, models.InternalError.Add(result.Error)
	}
	return environment, models.NoError
}

func (psql *environmentPostgres) GetEnvironmentsListByLimit(ctx context.Context, page int, pageSize int) ([]*models.Environment, models.AppError) {
	var environments []*models.Environment

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch environments with the specified pagination
	result := psql.db.Offset(offset).Table("lab_rank.environment").Limit(pageSize).Find(&environments)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return environments, models.NoError
}

// Add other repository methods for environments as needed.
