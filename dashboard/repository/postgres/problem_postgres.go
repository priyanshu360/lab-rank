package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type problemPostgres struct {
	db *gorm.DB
}

// NewProblemPostgresRepo creates a new PostgreSQL repository for problems.
func NewProblemPostgresRepo(db *gorm.DB) *problemPostgres {
	return &problemPostgres{db}
}

// CreateProblem creates a new problem.
func (psql *problemPostgres) CreateProblem(ctx context.Context, problem models.Problem) models.AppError {
	result := psql.db.WithContext(ctx).Create(problem)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetProblemByID retrieves a problem by its ID.
func (psql *problemPostgres) GetProblemByID(ctx context.Context, problemID uuid.UUID) (models.Problem, models.AppError) {
	var problem models.Problem
	result := psql.db.WithContext(ctx).First(&problem, problemID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Problem not found
			return problem, models.ProblemNotFoundError
		}
		return problem, models.InternalError.Add(result.Error)
	}
	return problem, models.NoError
}

// Add other repository methods for problems as needed.
