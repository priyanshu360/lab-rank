package submission

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/dashboard/repository"
	queue_models "github.com/priyanshu360/lab-rank/queue/models"
)

type Service interface {
	Create(context.Context, *models.Submission) (*models.Submission, models.AppError)
	Fetch(context.Context, int) (*models.Submission, models.AppError)
	Update(context.Context, int, *models.Submission) (*models.Submission, models.AppError)
	FetchForUserID(context.Context, uuid.UUID) ([]*models.SubmissionWithProblemTitle, models.AppError)
}

type service struct {
	repo repository.SubmissionRepository
	fs   repository.FileSystem
	// Todo : use interface of msgq
	msgq *queue_models.RabbitMQ
}

func New(repo repository.SubmissionRepository, fs repository.FileSystem, msgq *queue_models.RabbitMQ) *service {
	return &service{
		repo: repo,
		fs:   fs,
		msgq: msgq,
	}
}

func (s *service) Create(ctx context.Context, submission *models.Submission) (*models.Submission, models.AppError) {

	var err models.AppError
	if submission.Link, err = s.fs.StoreFile(ctx, []byte(submission.Solution), fmt.Sprintf("%d", submission.ID), models.SOLUTION, submission.Lang.GetExtension()); err != models.NoError {
		return nil, err
	}

	if err := s.repo.CreateSubmission(ctx, *submission); err != models.NoError {
		return nil, err
	}

	go s.addToQueue(ctx, submission)

	return submission, models.NoError
}

func (s *service) Update(ctx context.Context, id int, updatedSubmission *models.Submission) (*models.Submission, models.AppError) {
	submission, err := s.repo.GetSubmissionByID(ctx, id)
	if err != models.NoError {
		return nil, err
	}

	submission.UpdateFrom(*updatedSubmission)

	err = s.repo.UpdateSubmission(ctx, id, submission)
	return &submission, err
}

func (s *service) addToQueue(ctx context.Context, submission *models.Submission) {
	// Todo : how to handle failures
	queueObj, err := s.repo.GetQueueData(ctx, *submission)
	if err != models.NoError {
		slog.Error(err.Error())
	}

	log.Print(queueObj)

	message, _ := json.Marshal(queueObj)
	s.msgq.Publish(message)
}

func (s *service) Fetch(ctx context.Context, id int) (*models.Submission, models.AppError) {
	var submission models.Submission
	var err models.AppError
	if submission, err = s.repo.GetSubmissionByID(ctx, id); err != models.NoError {
		return nil, err
	}
	return &submission, err
}

func (s *service) FetchForUserID(ctx context.Context, userID uuid.UUID) ([]*models.SubmissionWithProblemTitle, models.AppError) {
	return s.repo.GetSubmissionsWithTitleByUserID(ctx, userID)
}
