package job

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"math/rand/v2"
	"strconv"
	"time"
)

type Processor interface {
	Process(job *Job) error
}

type stringProcessor struct {
	ctx context.Context
}

func NewStringProcessor(ctx context.Context) Processor {
	return &stringProcessor{
		ctx: ctx,
	}
}

func (p *stringProcessor) Process(job *Job) error {
	l := logger.FromContext(p.ctx)

	if job.TaskData != nil {
		l.With("data", string(job.TaskData)).Info("processing job")
	}

	n := rand.IntN(30)
	if n < 5 {
		n = 5
	}

	job.Status = constants.JOB_PROCESSING

	d, _ := time.ParseDuration(strconv.Itoa(n) + "s")
	time.Sleep(d)
	job.Status = constants.JOB_COMPLETED
	return nil
}
