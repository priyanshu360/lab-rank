package subject

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type SubjectService interface {
	Create(context.Context, *models.Subject) (*models.Subject, models.AppError)
}

type subjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) *subjectService {
	return &subjectService{
		repo: repo,
	}
}

func (s *subjectService) Create(ctx context.Context, subject *models.Subject) (*models.Subject, models.AppError) {
	subject.ID = uuid.New()

	if err := s.repo.CreateSubject(ctx, *subject); err != models.NoError {
		return nil, err
	}

	return subject, models.NoError
}
