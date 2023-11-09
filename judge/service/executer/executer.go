package executer

import (
	"github.com/priyanshu360/lab-rank/judge/models"
	"github.com/priyanshu360/lab-rank/judge/repository"
)

type Executer struct {
	queue repository.QueueRepository
}

func (e Executer) Run(submission models.SubmissionData) {

}

// psql db
// queue
// logic
// watcher
