package syllabus

import (
	"context"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type SyllabusService interface {
	Create(context.Context, *models.Syllabus) (*models.Syllabus, models.AppError)
	Fetch(context.Context, string) (*models.Syllabus, models.AppError)
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

func(s *syllabusService) Fetch(ctx context.Context, id string) (*models.Syllabus, models.AppError){
	if syllabusID,err := uuid.Parse(id); err != nil{
		return nil,models.InternalError.Add(err)
	}else{
		if syllabus,err := s.repo.GetSyllabusByID(ctx,syllabusID); err != models.NoError{
			return nil, err
		}else{
			return &syllabus,models.NoError
		}
	}
}