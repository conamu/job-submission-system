package job

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/google/uuid"
)

type Job struct {
	ctx      context.Context
	Id       string
	TaskData string
	Status   constants.Status
}

func CreateJob(ctx context.Context, data string) *Job {
	return &Job{
		ctx:      ctx,
		Id:       uuid.New().String(),
		TaskData: data,
		Status:   constants.JOB_PENDING,
	}
}
