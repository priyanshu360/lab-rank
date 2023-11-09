package queue

import (
	"context"

	"github.com/priyanshu360/lab-rank/judge/models"
	"github.com/priyanshu360/lab-rank/judge/repository"
)

type Queue struct {
	repo  repository.SubmissionRepository
	queue repository.QueueRepository
}

func NewQueue(repo repository.SubmissionRepository, queue repository.QueueRepository) Queue {
	return Queue{
		repo:  repo,
		queue: queue,
	}
}

func (q Queue) Refresh(ctx context.Context) Queue {
	return q
}

func (q Queue) Front(ctx context.Context) models.SubmissionData {
	return q.queue.GetNext(ctx)
}
