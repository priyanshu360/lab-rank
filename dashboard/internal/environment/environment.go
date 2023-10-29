package environment

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type EnvironmentService interface {
	Create(context.Context, *models.Environment) (*models.Environment, models.AppError)
}

type environmentService struct {
	repo repository.EnvironmentRepository
}

func NewEnvironmentService(repo repository.EnvironmentRepository) *environmentService {
	return &environmentService{
		repo: repo,
	}
}

func (s *environmentService) Create(ctx context.Context, environment *models.Environment) (*models.Environment, models.AppError) {
	environment.ID = uuid.New()

	if err := s.repo.CreateEnvironment(ctx, *environment); err != models.NoError {
		return nil, err
	}

	return environment, models.NoError
}
