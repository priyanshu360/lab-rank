package environment

import (
	"context"
	"fmt"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Environment) (*models.Environment, models.AppError)
	Fetch(context.Context, int) (*models.Environment, models.AppError)
}

type service struct {
	repo repository.EnvironmentRepository
	fs   repository.FileSystem
}

func New(repo repository.EnvironmentRepository, fs repository.FileSystem) *service {
	return &service{
		repo: repo,
		fs:   fs,
	}
}

func (s *service) Create(ctx context.Context, environment *models.Environment) (*models.Environment, models.AppError) {
	var err models.AppError
	if environment.Link, err = s.fs.StoreFile(ctx, environment.File, fmt.Sprintf("%d", environment.ID), models.ENVIRONMENT, models.YAML.GetExtension()); err != models.NoError {
		return nil, err
	}

	if err := s.repo.CreateEnvironment(ctx, *environment); err != models.NoError {
		return nil, err
	}

	return environment, models.NoError
}

// Todo : fix single value response for list api
func (s *service) Fetch(ctx context.Context, id int) (*models.Environment, models.AppError) {
	var environment models.Environment
	var err models.AppError

	if environment, err = s.repo.GetEnvironmentByID(ctx, id); err != models.NoError {
		return nil, err
	}

	return &environment, err
}
