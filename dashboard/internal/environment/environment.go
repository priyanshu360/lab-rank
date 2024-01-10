package environment

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Environment) (*models.Environment, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Environment, models.AppError)
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
	environment.ID = uuid.New()

	var err models.AppError
	if environment.Link, err = s.fs.StoreFile(ctx, environment.File, environment.ID, models.ENVIRONMENT, models.YAML.GetExtension()); err != models.NoError {
		return nil, err
	}

	if err := s.repo.CreateEnvironment(ctx, *environment); err != models.NoError {
		return nil, err
	}

	return environment, models.NoError
}

// Todo : fix single value response for list api
func (s *service) Fetch(ctx context.Context, id, limit string) ([]*models.Environment, models.AppError) {
	var environments []*models.Environment
	switch {
	case id != "":
		if environmentID, err := uuid.Parse(id); err != nil {
			return environments, models.InternalError.Add(err)
		} else {
			if environment, err := s.repo.GetEnvironmentByID(ctx, environmentID); err != models.NoError {
				return nil, err
			} else {
				environments = append(environments, &environment)
				return environments, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetEnvironmentsListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetEnvironmentsListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetEnvironmentsListByLimit(ctx, 1, 10)
	}
}
