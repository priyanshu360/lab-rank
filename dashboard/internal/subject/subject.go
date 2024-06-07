package subject

import (
	"context"

	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Subject) (*models.Subject, models.AppError)
	Fetch(context.Context, int) (*models.Subject, models.AppError)
	FetchByUniversityID(context.Context, int) ([]*models.Subject, models.AppError)
}

type service struct {
	syllabus syllabus.Service
	repo     repository.SubjectRepository
}

func New(repo repository.SubjectRepository, syllabus syllabus.Service) *service {
	return &service{
		syllabus: syllabus,
		repo:     repo,
	}
}

func (s *service) Create(ctx context.Context, subject *models.Subject) (*models.Subject, models.AppError) {

	if err := s.repo.CreateSubject(ctx, *subject); err != models.NoError {
		return nil, err
	}

	go s.syllabus.AutoGenerateFromSubject(context.Background(), subject)

	return subject, models.NoError
}

func (s *service) Fetch(ctx context.Context, id int) (*models.Subject, models.AppError) {
	var subject models.Subject
	var err models.AppError

	if subject, err = s.repo.GetSubjectByID(ctx, id); err != models.NoError {
		return nil, err
	}
	return &subject, err
}
func (s *service) FetchByUniversityID(ctx context.Context, universityId int) ([]*models.Subject, models.AppError) {
	return s.repo.GetSubjectsByUniversityID(ctx, universityId)
}
