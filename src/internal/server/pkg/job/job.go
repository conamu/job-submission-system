package job

import (
	"context"
	"github.com/google/uuid"
)

type Status string

const (
	JOB_PENDING    Status = "PENDING"
	JOB_PROCESSING Status = "PROCESSING"
	JOB_COMPLETED  Status = "COMPLETED"
)

type Job struct {
	ctx    context.Context
	id     string
	taskFn func(data any)
	status Status
}

func CreateJob(ctx context.Context, fn func(data any)) *Job {
	return &Job{
		ctx:    ctx,
		id:     uuid.New().String(),
		taskFn: fn,
		status: JOB_PENDING,
	}
}
