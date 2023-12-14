package syllabus

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Syllabus) (*models.Syllabus, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Syllabus, models.AppError)
	AutoGenerateFromCollege(context.Context, *models.College) models.AppError
	AutoGenerateFromSubject(context.Context, *models.Subject) models.AppError
}

type service struct {
	repo repository.SyllabusRepository
}

func New(repo repository.SyllabusRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, syllabus *models.Syllabus) (*models.Syllabus, models.AppError) {
	syllabus.ID = uuid.New()

	if err := s.repo.CreateSyllabus(ctx, *syllabus); err != models.NoError {
		return nil, err
	}

	return syllabus, models.NoError
}

func (s *service) AutoGenerateFromCollege(ctx context.Context, college *models.College) models.AppError {
	subjects, err := s.repo.GetSubjectsByUniversityID(ctx, college.ID)
	if err != models.NoError {
		return err
	}

	for _, subject := range subjects {
		syllabus := models.Syllabus{
			ID:            uuid.New(),
			SubjectID:     subject.ID,
			SyllabusLevel: models.SyllabusLevelCollege,
			UniCollegeID:  college.ID,
		}
		if _, err := s.Create(ctx, &syllabus); err != models.NoError {
			log.Println("AutoGenerateFromSubject ", err)
		}
	}
	return models.NoError
}

func (s *service) AutoGenerateFromSubject(ctx context.Context, subject *models.Subject) models.AppError {
	collegeIDs, err := s.repo.GetCollegeIDsForUniversityID(ctx, subject.UniversityID)
	if err != models.NoError {
		return err
	}

	for _, cID := range collegeIDs {
		syllabus := []models.Syllabus{
			{
				ID:            uuid.New(),
				SubjectID:     subject.ID,
				SyllabusLevel: models.SyllabusLevelCollege,
				UniCollegeID:  cID,
			},
			{
				ID:            uuid.New(),
				SubjectID:     subject.ID,
				SyllabusLevel: models.SyllabusLevelUniversity,
				UniCollegeID:  subject.UniversityID,
			},
		}
		if _, err := s.Create(ctx, &syllabus[0]); err != models.NoError {
			log.Println("AutoGenerateFromSubject ", err)
		}
		if _, err := s.Create(ctx, &syllabus[1]); err != models.NoError {
			log.Println("AutoGenerateFromSubject ", err)
		}
	}
	return models.NoError
}

func (s *service) Fetch(ctx context.Context, id, limit string) ([]*models.Syllabus, models.AppError) {
	var syllabuss []*models.Syllabus
	switch {
	case id != "":
		if syllabusID, err := uuid.Parse(id); err != nil {
			return syllabuss, models.InternalError.Add(err)
		} else {
			if syllabus, err := s.repo.GetSyllabusByID(ctx, syllabusID); err != models.NoError {
				return nil, err
			} else {
				syllabuss = append(syllabuss, &syllabus)
				return syllabuss, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetSyllabusListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetSyllabusListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetSyllabusListByLimit(ctx, 1, 10)
	}
}
