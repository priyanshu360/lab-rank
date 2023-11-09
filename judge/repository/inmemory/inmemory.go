package inmemory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/judge/models"
)

type InMemoryQueue struct {
	mu          sync.Mutex
	submissions []models.SubmissionData
}

// Add adds a new submission to the in-memory queue.
func (q *InMemoryQueue) Add(ctx context.Context, submission models.SubmissionData) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.submissions = append(q.submissions, submission)
	return true
}

// Remove removes a submission from the in-memory queue.
func (q *InMemoryQueue) Remove(ctx context.Context, ID uuid.UUID) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	for i, sub := range q.submissions {
		if sub.ID == ID {
			q.submissions = append(q.submissions[:i], q.submissions[i+1:]...)
			return true
		}
	}
	return false
}

// GetNext returns the next submission from the in-memory queue.
func (q *InMemoryQueue) GetNext(ctx context.Context) models.SubmissionData {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.submissions) > 0 {
		nextSubmission := q.submissions[0]
		q.submissions = q.submissions[1:]
		return nextSubmission
	}

	// If the queue is empty, return an empty SubmissionData or handle as needed.
	return models.SubmissionData{}
}

// Update updates the information of a submission in the in-memory queue.
func (q *InMemoryQueue) Update(ctx context.Context, updatedSubmission models.SubmissionData) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for i, sub := range q.submissions {
		if sub.ID == updatedSubmission.ID {
			q.submissions[i] = updatedSubmission
			return
		}
	}
}
