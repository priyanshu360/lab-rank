package syllabus

import (
	"context"
	"log"

	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type Service interface {
	Create(context.Context, *models.Syllabus) (*models.Syllabus, models.AppError)
	Fetch(context.Context, int) (*models.Syllabus, models.AppError)
	AutoGenerateFromCollege(context.Context, *models.College) models.AppError
	AutoGenerateFromSubject(context.Context, *models.Subject) models.AppError
	FetchBySubjectID(context.Context, int) ([]*models.Syllabus, models.AppError)
	// UpdateAccessIDsForUser(context.Context, *models.User) models.AppError
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

	if err := s.repo.CreateSyllabus(ctx, *syllabus); err != models.NoError {
		return nil, err
	}

	return syllabus, models.NoError
}

func (s *service) AutoGenerateFromCollege(ctx context.Context, college *models.College) models.AppError {
	subjects, err := s.repo.GetSubjectsByUniversityID(ctx, college.UniversityID)
	if err != models.NoError {
		return err
	}

	for _, subject := range subjects {
		syllabus := subject.ToSyllabus(college.ID, models.SyllabusLevelCollege)
		if _, err := s.Create(ctx, syllabus); err != models.NoError {
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
		syllabus := []*models.Syllabus{
			subject.ToSyllabus(cID, models.SyllabusLevelCollege),
			subject.ToSyllabus(subject.UniversityID, models.SyllabusLevelUniversity),
		}
		if _, err := s.Create(ctx, syllabus[0]); err != models.NoError {
			log.Println("AutoGenerateFromSubject ", err)
		}
		if _, err := s.Create(ctx, syllabus[1]); err != models.NoError {
			log.Println("AutoGenerateFromSubject ", err)
		}
	}
	return models.NoError
}

func (s *service) Fetch(ctx context.Context, id int) (*models.Syllabus, models.AppError) {
	var syllabus models.Syllabus
	var err models.AppError

	if syllabus, err = s.repo.GetSyllabusByID(ctx, id); err != models.NoError {
		return nil, err
	}

	return &syllabus, models.NoError
}

func (s *service) FetchBySubjectID(ctx context.Context, subjectID int) ([]*models.Syllabus, models.AppError) {
	// You may need to convert subjectID to uuid.UUID if necessary
	// For example, subjectUUID, err := uuid.Parse(subjectID)

	syllabuss, err := s.repo.GetSyllabusBySubjectID(ctx, subjectID)
	if err != models.NoError {
		return nil, err
	}

	return syllabuss, models.NoError
}
