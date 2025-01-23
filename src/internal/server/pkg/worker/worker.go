package worker

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/job"
	"log/slog"
	"sync"
)

type Worker struct {
	ctx          context.Context
	wg           *sync.WaitGroup
	processor    job.Processor
	queueChannel chan *job.Job
	log          *slog.Logger
}

func Create(ctx context.Context) *Worker {

	l := logger.FromContext(ctx)
	l = l.With("type", "worker")
	wg := ctx.Value(constants.CTX_WG).(*sync.WaitGroup)
	queue := job.FromContext(ctx)

	w := &Worker{
		ctx:          ctx,
		wg:           wg,
		processor:    job.NewStringProcessor(ctx),
		queueChannel: queue.GetJobQueue(),
		log:          l,
	}
	go w.start()

	return w
}

func (w *Worker) start() {
	w.log.Debug("starting worker")
	w.wg.Add(1)

	for {
		select {
		case <-w.ctx.Done():
			w.log.Debug("shutting down worker")
			w.wg.Done()
			return
		case j := <-w.queueChannel:
			err := w.processor.Process(j)
			if err != nil {
				w.log.
					With("error", err.Error()).
					With("job", j.Id).
					Error("error processing job")
			}
		}
	}
}
