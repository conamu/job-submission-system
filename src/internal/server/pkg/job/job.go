package job

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/google/uuid"
)

type Job struct {
	ctx      context.Context
	Id       string
	TaskData []byte
	Status   constants.Status
}

func CreateJob(ctx context.Context, data []byte) *Job {
	return &Job{
		ctx:      ctx,
		Id:       uuid.New().String(),
		TaskData: data,
		Status:   constants.JOB_PENDING,
	}
}
