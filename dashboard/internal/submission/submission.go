package submission

import (
	"context"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
)

type SubmissionService interface {
	Create(context.Context, *models.Submission) (*models.Submission, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Submission, models.AppError)
}

type submissionService struct {
	repo repository.SubmissionRepository
	fs   repository.FileSystem
}

func NewSubmissionService(repo repository.SubmissionRepository, fs repository.FileSystem) *submissionService {
	return &submissionService{
		repo: repo,
		fs:   fs,
	}
}

func (s *submissionService) Create(ctx context.Context, submission *models.Submission) (*models.Submission, models.AppError) {
	submission.ID = uuid.New()

	var err models.AppError
	if submission.Link, err = s.fs.StoreFile(ctx, []byte(submission.Solution), submission.ID, models.SOLUTION, submission.Lang.GetExtension()); err != models.NoError {
		return nil, err
	}

	if err := s.repo.CreateSubmission(ctx, *submission); err != models.NoError {
		return nil, err
	}

	return submission, models.NoError
}

func (s *submissionService) Fetch(ctx context.Context, id, limit string) ([]*models.Submission, models.AppError) {
	var submissions []*models.Submission
	switch {
	case id != "":
		if submissionID, err := uuid.Parse(id); err != nil {
			return submissions, models.InternalError.Add(err)
		} else {
			if submission, err := s.repo.GetSubmissionByID(ctx, submissionID); err != models.NoError {
				return nil, err
			} else {
				submissions = append(submissions, &submission)
				return submissions, models.NoError
			}
		}

	case limit != "":
		if limit, err := strconv.ParseInt(limit, 10, 64); err != nil {
			return s.repo.GetSubmissionsListByLimit(ctx, 1, 10)

		} else {
			return s.repo.GetSubmissionsListByLimit(ctx, 1, int(limit))
		}
	default:

		return s.repo.GetSubmissionsListByLimit(ctx, 1, 10)
	}
}
