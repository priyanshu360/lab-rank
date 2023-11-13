package watcher

import (
	"context"

	dashboard_models "github.com/priyanshu360/lab-rank/dashboard/models"
	"github.com/priyanshu360/lab-rank/judge/repository"
	"github.com/priyanshu360/lab-rank/judge/service/executer"
	"github.com/priyanshu360/lab-rank/judge/service/queue"
)

type Watcher struct {
	svc   executer.Executer
	queue queue.Queue
	repo  repository.SubmissionRepository
}

func NewWatcher(svc executer.Executer, queue queue.Queue) Watcher {
	return Watcher{
		svc:   svc,
		queue: queue,
	}
}

func (w Watcher) Run(ctx context.Context) {
	for {
		w.queue.Refresh(ctx)
		for !w.queue.IsEmpty(ctx) {
			candidate := w.queue.Front(ctx)
			candidate.Status = dashboard_models.Accepted
			w.repo.Update(ctx, candidate)
			w.svc.Run(ctx, w.queue.Front(ctx))
		}
	}
}
