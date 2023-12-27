package subject

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Subject) (*models.Subject, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Subject, models.AppError)
	FetchByUniversityID(context.Context, uuid.UUID) ([]*models.Subject, models.AppError)
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
	subject.ID = uuid.New()

	if err := s.repo.CreateSubject(ctx, *subject); err != models.NoError {
		return nil, err
	}

	go s.syllabus.AutoGenerateFromSubject(context.Background(), subject)

	return subject, models.NoError
}

func (s *service) Fetch(ctx context.Context, id, limit string) ([]*models.Subject, models.AppError) {
	var subjects []*models.Subject
	switch {
	case id != "":
		if subjectID, err := uuid.Parse(id); err != nil {
			return subjects, models.InternalError.Add(err)
		} else {
			if subject, err := s.repo.GetSubjectByID(ctx, subjectID); err != models.NoError {
				return nil, err
			} else {
				subjects = append(subjects, &subject)
				return subjects, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetSubjectsListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetSubjectsListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetSubjectsListByLimit(ctx, 1, 10)
	}
}
func (s *service) FetchByUniversityID(ctx context.Context, universityId uuid.UUID) ([]*models.Subject, models.AppError) {
	return s.repo.GetSubjectsByUniversityID(ctx, universityId)
}
