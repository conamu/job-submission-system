package job

import (
	"context"
	"errors"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/constant"
	"sync"
)

type Queue struct {
	jobQueue     chan *Job
	jobStatusMap *sync.Map
}

var ErrQueueFull = errors.New("job queue is full")
var ErrStatusNotFound = errors.New("job status was not found")

func CreateQueue() *Queue {
	return &Queue{
		jobQueue:     make(chan *Job, 10),
		jobStatusMap: &sync.Map{},
	}
}

func (q *Queue) Place(job *Job) (string, error) {
	if len(q.jobQueue) == 10 {
		return "", ErrQueueFull
	}
	q.jobStatusMap.Store(job.Id, job)
	q.jobQueue <- job
	return job.Id, nil
}

func (q *Queue) GetStatus(id string) (Status, error) {
	if data, ok := q.jobStatusMap.Load(id); ok {
		j := data.(*Job)
		return j.Status, nil
	}
	return "", ErrStatusNotFound
}

func (q *Queue) GetJobQueue() chan *Job {
	return q.jobQueue
}

func FromContext(ctx context.Context) *Queue {
	return ctx.Value(constant.CTX_QUEUE).(*Queue)
}
