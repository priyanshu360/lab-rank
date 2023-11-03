package environment

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type EnvironmentService interface {
	Create(context.Context, *models.Environment) (*models.Environment, models.AppError)
	Fetch(context.Context, string) (*models.Environment, models.AppError)
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

func(s *environmentService) Fetch(ctx context.Context, id string) (*models.Environment, models.AppError){
	if envID,err := uuid.Parse(id); err != nil{
		return nil,models.InternalError.Add(err)
	}else{
		if environment,err := s.repo.GetEnvironmentByID(ctx,envID); err != models.NoError{
			return nil, err
		}else{
			return &environment,models.NoError
		}
	}
}