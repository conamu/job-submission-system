package worker

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/constant"
)

type Pool []*Worker

func CreatePool(ctx context.Context, workerAmount int) Pool {

	p := make([]*Worker, workerAmount)

	for idx, _ := range p {
		p[idx] = Create(ctx)
	}

	return p
}

func FromContext(ctx context.Context) Pool {
	return ctx.Value(constant.CTX_POOL).(Pool)
}
