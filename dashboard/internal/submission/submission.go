package submission

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
	queue_models "github.com/priyanshu360/lab-rank/queue/models"
)

type SubmissionService interface {
	Create(context.Context, *models.Submission) (*models.Submission, models.AppError)
	Fetch(context.Context, string, string) ([]*models.Submission, models.AppError)
}

type submissionService struct {
	repo repository.SubmissionRepository
	fs   repository.FileSystem
	// Todo : use interface of msgq
	msgq *queue_models.RabbitMQ
}

func NewSubmissionService(repo repository.SubmissionRepository, fs repository.FileSystem, msgq *queue_models.RabbitMQ) *submissionService {
	return &submissionService{
		repo: repo,
		fs:   fs,
		msgq: msgq,
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

	go s.addToQueue(ctx, submission)

	return submission, models.NoError
}

func (s *submissionService) addToQueue(ctx context.Context, submission *models.Submission) {
	// Todo : how to handle failures
	queueObj, err := s.repo.GetQueueData(ctx, *submission)
	if err != models.NoError {
		slog.Error(err.Error())
	}

	message, _ := json.Marshal(queueObj)
	s.msgq.Publish(message)
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
