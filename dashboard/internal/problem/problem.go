package problem

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type ProblemService interface {
	Create(context.Context, *models.Problem) (*models.Problem, models.AppError)
}

type problemService struct {
	repo repository.ProblemRepository
}

func NewProblemService(repo repository.ProblemRepository) *problemService {
	return &problemService{
		repo: repo,
	}
}

func (s *problemService) Create(ctx context.Context, problem *models.Problem) (*models.Problem, models.AppError) {
	problem.ID = uuid.New()

	if err := s.repo.CreateProblem(ctx, *problem); err != models.NoError {
		return nil, err
	}

	return problem, models.NoError
}
