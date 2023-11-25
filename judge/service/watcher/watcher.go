package watcher

import (
	"context"
	"encoding/json"
	"log"

	"github.com/priyanshu360/lab-rank/judge/service/executer"
	// "github.com/priyanshu360/lab-rank/judge/service/queue"
	queue_models "github.com/priyanshu360/lab-rank/queue/models"
)

type Watcher struct {
	svc      executer.Executer
	consumer *queue_models.RabbitMQ
}

func NewWatcher(svc executer.Executer, consumer *queue_models.RabbitMQ) Watcher {
	return Watcher{
		svc:      svc,
		consumer: consumer,
	}
}

// Todo make things parallel
func (w Watcher) Run(ctx context.Context) {
	forever := make(chan bool)
	go func() {
		msgq := w.consumer.GetMsgCh()
		for d := range msgq {
			var queueObj queue_models.QueueObj
			msg := d.Body
			log.Printf("Received a message: %s", msg)
			json.Unmarshal(msg, &queueObj)
			// todo : manage number of cuncurrent running tasks
			go w.svc.Run(ctx, queueObj)
		}
	}()

	log.Println("[*] Waiting for message. To exit press CTRL+C")
	<-forever
}
