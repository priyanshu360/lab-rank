package university

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.University) (*models.University, models.AppError)
	Fetch(context.Context, int) (*models.University, models.AppError)
	GetAllUniversityNames(context.Context) ([]*models.UniversityIdName, models.AppError)
}

type service struct {
	repo repository.UniversityRepository
}

func New(repo repository.UniversityRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, university *models.University) (*models.University, models.AppError) {

	if err := s.repo.CreateUniversity(ctx, university); err != models.NoError {
		return nil, err
	}

	return university, models.NoError
}

func (s *service) Fetch(ctx context.Context, id int) (*models.University, models.AppError) {
	var university models.University
	var err models.AppError
	if university, err = s.repo.GetUniversityByID(ctx, id); err != models.NoError {
		return nil, err
	}
	return &university, models.NoError
}

func (s *service) GetAllUniversityNames(ctx context.Context) ([]*models.UniversityIdName, models.AppError) {

	universities, err := s.repo.GetAllUniversityNames(ctx)
	return universities, err
}
