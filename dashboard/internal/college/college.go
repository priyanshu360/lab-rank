package college

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type CollegeService interface {
	Create(context.Context, *models.College) (*models.College, models.AppError)
}

type collegeService struct {
	repo repository.CollegeRepository
}

func NewCollegeService(repo repository.CollegeRepository) *collegeService {
	return &collegeService{
		repo: repo,
	}
}

func (s *collegeService) Create(ctx context.Context, college *models.College) (*models.College, models.AppError) {
	college.ID = uuid.New()

	if err := s.repo.CreateCollege(ctx, *college); err != models.NoError {
		return nil, err
	}

	return college, models.NoError
}
