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
	result := psql.db.WithContext(ctx).Table("lab_rank.problems").Create(problem)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetProblemByID retrieves a problem by its ID.
func (psql *problemPostgres) GetProblemByID(ctx context.Context, problemID uuid.UUID) (models.Problem, models.AppError) {
	var problem models.Problem
	result := psql.db.WithContext(ctx).Table("lab_rank.problems").First(&problem, problemID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Problem not found
			return problem, models.ProblemNotFoundError
		}
		return problem, models.InternalError.Add(result.Error)
	}
	return problem, models.NoError
}

func (psql *problemPostgres) GetProblemsListByLimit(ctx context.Context, page int, pageSize int) ([]*models.Problem, models.AppError) {
	var problems []*models.Problem

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch problems with the specified pagination
	result := psql.db.Offset(offset).Table("lab_rank.problems").Limit(pageSize).Find(&problems)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return problems, models.NoError
}

// UpdateProblem updates a Problem's information.
func (psql *problemPostgres) UpdateProblem(ctx context.Context, problemID uuid.UUID, problem models.Problem) models.AppError {
	// Check if the Problem with the provided ID exists before updating
	var existingProblem models.Problem
	result := psql.db.WithContext(ctx).First(&existingProblem, problemID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Problem not found
			return models.ProblemNotFoundError
		}
		return models.InternalError.Add(result.Error)
	}

	// Perform the update
	result = psql.db.WithContext(ctx).Model(&problem).Where("id = ?", problemID).Updates(problem)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// Add other repository methods for problems as needed.
