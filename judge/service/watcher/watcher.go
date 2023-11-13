package watcher

import (
	"context"

	"github.com/priyanshu360/lab-rank/judge/service/executer"
	"github.com/priyanshu360/lab-rank/judge/service/queue"
)

type Watcher struct {
	svc   executer.Executer
	queue queue.Queue
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
			w.svc.Run(ctx, w.queue.Front(ctx))
		}
	}
}
