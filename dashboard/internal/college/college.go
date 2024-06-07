package college

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.College) (*models.College, models.AppError)
	Fetch(context.Context, int) (*models.College, models.AppError)
	GetCollegeIDsForUniversityID(context.Context, int) ([]*models.CollegeIdName, models.AppError)
}

type service struct {
	syllabus syllabus.Service
	repo     repository.CollegeRepository
}

func New(repo repository.CollegeRepository, syllabus syllabus.Service) *service {
	return &service{
		syllabus: syllabus,
		repo:     repo,
	}
}

func (s *service) Create(ctx context.Context, college *models.College) (*models.College, models.AppError) {

	if err := s.repo.CreateCollege(ctx, *college); err != models.NoError { // Todo: Check if university id exists
		return nil, err
	}

	go s.syllabus.AutoGenerateFromCollege(context.Background(), college)

	return college, models.NoError
}

func (s *service) Fetch(ctx context.Context, id int) (*models.College, models.AppError) {

	var college models.College
	var err models.AppError

	if college, err = s.repo.GetCollegeByID(ctx, id); err != models.NoError {
		return nil, err
	}
	return &college, err
}

func (s *service) GetCollegeIDsForUniversityID(ctx context.Context, universityId int) ([]*models.CollegeIdName, models.AppError) {
	return s.repo.GetCollegesByUniversityID(ctx, universityId)
}
