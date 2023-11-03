package college

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type CollegeService interface {
	Create(context.Context, *models.College) (*models.College, models.AppError)
	Fetch(context.Context, string) (*models.College, models.AppError)
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

func(s *collegeService) Fetch(ctx context.Context, id string) (*models.College, models.AppError){
	if collegeID,err := uuid.Parse(id); err != nil{
		return nil,models.InternalError.Add(err)
	}else{
		if college,err := s.repo.GetCollegeByID(ctx,collegeID);  err != models.NoError{
			return nil, err
		}else{
			return &college, models.NoError
		}
	}
}