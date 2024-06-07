package postgres

import (
	"context"
	"fmt"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"gorm.io/gorm"
)

type problemPostgres struct {
	db *gorm.DB
}

// NewProblemPostgresRepo creates a new PostgreSQL repository for problems.
func NewProblemPostgresRepo(db *gorm.DB) *problemPostgres {
	if err := db.AutoMigrate(models.Problem{}); err != nil {
		panic(err)
	}
	return &problemPostgres{db}
}

// CreateProblem creates a new problem.
func (psql *problemPostgres) CreateProblem(ctx context.Context, problem models.Problem) models.AppError {
	result := psql.db.WithContext(ctx).Create(&problem)
	if result.Error != nil {
		return models.InternalError.Add(result.Error)
	}
	return models.NoError
}

// GetProblemByID retrieves a problem by its ID.
func (psql *problemPostgres) GetProblemByID(ctx context.Context, problemID int) (models.Problem, models.AppError) {
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

func (psql *problemPostgres) GetProblemsListByLimit(ctx context.Context, page int, pageSize int) ([]*models.Problem, models.AppError) {
	var problems []*models.Problem

	// Calculate the offset
	offset := (page - 1) * pageSize

	// Fetch problems with the specified pagination
	result := psql.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Find(&problems)
	if result.Error != nil {
		return nil, models.InternalError.Add(result.Error)
	}

	return problems, models.NoError
}

// Add other repository methods for problems as needed.

func (psql *problemPostgres) GetProblemsForSubject(ctx context.Context, subjectID, collegeID int) ([]*models.Problem, models.AppError) {
	var problems []*models.Problem

	// Todo : this need to be changed
	// Fetch syllabus IDs based on the given subject ID
	syllabusQuery := `
        WITH selected_syllabus AS (
            SELECT id
            FROM lab_rank.syllabus
            WHERE subject_id = ?
        )
        SELECT p.*
        FROM lab_rank.problems p
        JOIN selected_syllabus s ON p.syllabus_id = s.id
        WHERE p.syllabus_id IS NOT NULL AND p.problem_link IS NOT NULL;
    `

	// Use Gorm's Raw method to execute the raw SQL query
	result := psql.db.WithContext(ctx).Raw(syllabusQuery, subjectID).Scan(&problems)

	if result.Error != nil {
		return nil, models.InternalError.Add(fmt.Errorf("Failed to execute syllabus query %s", result.Error))
	}

	return problems, models.NoError
}
