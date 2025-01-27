package job

import (
	"context"
	"errors"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/spf13/viper"
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
		jobQueue:     make(chan *Job, viper.GetInt("server.queueSize")),
		jobStatusMap: &sync.Map{},
	}
}

func (q *Queue) Place(job *Job) (string, error) {
	if len(q.jobQueue) == viper.GetInt("server.queueSize") {
		return "", ErrQueueFull
	}
	q.jobStatusMap.Store(job.Id, job)
	q.jobQueue <- job
	return job.Id, nil
}

func (q *Queue) GetStatus(id string) (constants.Status, error) {
	if data, ok := q.jobStatusMap.Load(id); ok {
		j := data.(*Job)
		return j.Status, nil
	}
	return "", ErrStatusNotFound
}

func (q *Queue) GetJobQueue() chan *Job {
	return q.jobQueue
}

func (q *Queue) GetJobStatuses() *sync.Map {
	return q.jobStatusMap
}

func FromContext(ctx context.Context) *Queue {
	return ctx.Value(constants.CTX_QUEUE).(*Queue)
}
