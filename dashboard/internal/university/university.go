package university

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.University) (*models.University, models.AppError)
	Fetch(context.Context, string, string) ([]*models.University, models.AppError)
	GetAllUniversityNames(context.Context) ([]*models.UniversityIdName, models.AppError)
}

type service struct {
	repo repository.UniversityRepository
}

func New(repo repository.UniversityRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, university *models.University) (*models.University, models.AppError) {
	university.ID = uuid.New()

	if err := s.repo.CreateUniversity(ctx, *university); err != models.NoError {
		return nil, err
	}

	return university, models.NoError
}

func (s *service) Fetch(ctx context.Context, id, limit string) ([]*models.University, models.AppError) {
	var universities []*models.University
	switch {
	case id != "":
		if universityID, err := uuid.Parse(id); err != nil {
			return universities, models.InternalError.Add(err)
		} else {
			if university, err := s.repo.GetUniversityByID(ctx, universityID); err != models.NoError {
				return nil, err
			} else {
				universities = append(universities, &university)
				return universities, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetUniversitiesListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetUniversitiesListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetUniversitiesListByLimit(ctx, 1, 10)
	}
}

func (s *service) GetAllUniversityNames(ctx context.Context) ([]*models.UniversityIdName, models.AppError) {

	universities, err := s.repo.GetAllUniversityNames(ctx)
	return universities, err
}
