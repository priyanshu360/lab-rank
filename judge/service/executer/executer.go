package executer

import (
	"context"
	"log"

	"github.com/priyanshu360/lab-rank/judge/models"
)

type Executer struct {
}

func (e Executer) Run(ctx context.Context, submission models.SubmissionData) {
	log.Println(submission)
}

func NewExecuter() Executer {
	return Executer{}
}
