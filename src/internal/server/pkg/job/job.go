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
	ctx      context.Context
	Id       string
	TaskData []byte
	Status   Status
}

func CreateJob(ctx context.Context, data []byte) *Job {
	return &Job{
		ctx:      ctx,
		Id:       uuid.New().String(),
		TaskData: data,
		Status:   JOB_PENDING,
	}
}
