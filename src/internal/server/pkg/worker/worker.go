package worker

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"log/slog"
)

type Worker struct {
	ctx context.Context
	log *slog.Logger
}

func Create(ctx context.Context) *Worker {

	l := logger.FromContext(ctx)
	l = l.With("type", "worker")

	w := &Worker{
		ctx: ctx,
		log: l,
	}
	w.start()

	return w
}

func (w *Worker) start() {
	w.log.Debug("started worker")
}
