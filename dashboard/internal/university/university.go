package university

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type UniversityService interface {
	Create(context.Context, *models.University) (*models.University, models.AppError)
	Fetch(context.Context, string) (*models.University, models.AppError)
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

func(s *universityService) Fetch(ctx context.Context, id string) (*models.University, models.AppError){
	if universityID,err := uuid.Parse(id); err != nil{
		return nil,models.InternalError.Add(err)
	}else{
		if university,err := s.repo.GetUniversityByID(ctx,universityID); err != models.NoError{
			return nil, err
		}else{
			return &university,models.NoError
		}
	}
}