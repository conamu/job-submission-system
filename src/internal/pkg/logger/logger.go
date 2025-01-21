package logger

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/constant"
	"log/slog"
	"os"
)

func New(level slog.Level, app string) *slog.Logger {
	lh := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     level,
	})

	l := slog.New(lh)
	l = l.With("app", app)

	return l
}

func FromContext(ctx context.Context) *slog.Logger {
	return ctx.Value(constant.CTX_LOGGER).(*slog.Logger)
}
