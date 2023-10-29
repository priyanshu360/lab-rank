package university

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type UniversityService interface {
	Create(context.Context, *models.University) (*models.University, models.AppError)
}

type universityService struct {
	repo repository.UniversityRepository
}

func NewUniversityService(repo repository.UniversityRepository) *universityService {
	return &universityService{
		repo: repo,
	}
}

func (s *universityService) Create(ctx context.Context, university *models.University) (*models.University, models.AppError) {
	university.ID = uuid.New()

	if err := s.repo.CreateUniversity(ctx, *university); err != models.NoError {
		return nil, err
	}

	return university, models.NoError
}
