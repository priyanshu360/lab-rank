package college

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/internal/syllabus"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.College) (*models.College, models.AppError)
	Fetch(context.Context, string, string) ([]*models.College, models.AppError)
	GetCollegeIDsForUniversityID(context.Context, uuid.UUID) ([]*models.CollegeIdName, models.AppError)
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
	college.ID = uuid.New()

	if err := s.repo.CreateCollege(ctx, *college); err != models.NoError { // Todo: Check if university id exists
		return nil, err
	}

	go s.syllabus.AutoGenerateFromCollege(context.Background(), college)

	return college, models.NoError
}

func (s *service) Fetch(ctx context.Context, id, limit string) ([]*models.College, models.AppError) {
	var colleges []*models.College
	switch {
	case id != "":
		if collegeID, err := uuid.Parse(id); err != nil {
			return colleges, models.InternalError.Add(err)
		} else {
			if college, err := s.repo.GetCollegeByID(ctx, collegeID); err != models.NoError {
				return nil, err
			} else {
				colleges = append(colleges, &college)
				return colleges, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetCollegesListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetCollegesListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetCollegesListByLimit(ctx, 1, 10)
	}
}

func (s *service) GetCollegeIDsForUniversityID(ctx context.Context, universityId uuid.UUID) ([]*models.CollegeIdName, models.AppError) {
	return s.repo.GetCollegesByUniversityID(ctx, universityId)
}
