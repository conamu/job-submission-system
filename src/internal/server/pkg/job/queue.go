package job

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/constant"
	"sync"
)

type Queue interface {
	Place(job Job) string
	GetStatus(id string) Status
}
type queue struct {
	jobStatusMap *sync.Map
	jobs         chan Job
}

func CreateQueue() Queue {
	return &queue{
		jobStatusMap: &sync.Map{},
		jobs:         make(chan Job, 10),
	}
}

func (q queue) Place(job Job) string {
	//TODO implement me
	panic("implement me")
}

func (q queue) GetStatus(id string) Status {
	//TODO implement me
	panic("implement me")
}

func FromContext(ctx context.Context) Queue {
	return ctx.Value(constant.CTX_QUEUE).(*queue)
}
