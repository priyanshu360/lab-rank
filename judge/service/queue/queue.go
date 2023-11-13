package queue

import (
	"context"

	dashboard_models "github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/judge/models"
	"github.com/priyanshu360/lab-rank/judge/repository"
)

type Queue struct {
	repo  repository.SubmissionRepository
	queue repository.QueueRepository
}

func NewQueue(repo repository.SubmissionRepository, queue repository.QueueRepository) *Queue {
	return &Queue{
		repo:  repo,
		queue: queue,
	}
}

func (q *Queue) Refresh(ctx context.Context) *Queue {
	submissions_list := q.repo.GetNextBatch(ctx, dashboard_models.Queued, 10)
	for _, submission := range submissions_list {
		q.queue.Add(ctx, models.NewSubmissionData(submission))
	}
	return q
}

func (q *Queue) Front(ctx context.Context) models.SubmissionData {
	return q.queue.GetNext(ctx)
}

func (q *Queue) IsEmpty(ctx context.Context) bool {
	return q.queue.IsEmpty(ctx)
}
