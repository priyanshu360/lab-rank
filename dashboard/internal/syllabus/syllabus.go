package syllabus

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type SyllabusService interface {
	Create(context.Context, *models.Syllabus) (*models.Syllabus, models.AppError)
}

type syllabusService struct {
	repo repository.SyllabusRepository
}

func NewSyllabusService(repo repository.SyllabusRepository) *syllabusService {
	return &syllabusService{
		repo: repo,
	}
}

func (s *syllabusService) Create(ctx context.Context, syllabus *models.Syllabus) (*models.Syllabus, models.AppError) {
	syllabus.ID = uuid.New()

	if err := s.repo.CreateSyllabus(ctx, *syllabus); err != models.NoError {
		return nil, err
	}

	return syllabus, models.NoError
}
