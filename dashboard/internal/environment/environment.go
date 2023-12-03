package environment

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type EnvironmentService interface {
	Create(context.Context, *models.Environment) (*models.Environment, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Environment, models.AppError)
	Update(context.Context,*models.UpdateEnvironmentAPIRequest) (*models.Environment, models.AppError)
}

type environmentService struct {
	repo repository.EnvironmentRepository
	fs   repository.FileSystem
}

func NewEnvironmentService(repo repository.EnvironmentRepository, fs repository.FileSystem) *environmentService {
	return &environmentService{
		repo: repo,
		fs:   fs,
	}
}

func (s *environmentService) Create(ctx context.Context, environment *models.Environment) (*models.Environment, models.AppError) {
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

func (s *environmentService) Fetch(ctx context.Context, id, limit string) ([]*models.Environment, models.AppError) {
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

func (s *environmentService) Update(ctx context.Context, request *models.UpdateEnvironmentAPIRequest) (*models.Environment, models.AppError) {

	defaultEnvironment, err := s.repo.GetEnvironmentByID(ctx, request.ID)
	if err != models.NoError {
		return nil, err
	}
	updatedEnvironment := request.ToEnvironment(defaultEnvironment)
	if err := s.repo.UpdateEnvironment(ctx, request.ID, *updatedEnvironment); err != models.NoError {
		return nil, err
	}

	return updatedEnvironment, models.NoError
}
