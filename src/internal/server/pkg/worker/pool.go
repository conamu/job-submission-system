package worker

import (
	"context"
)

type Pool []*Worker

func CreatePool(ctx context.Context, workerAmount int) Pool {

	p := make([]*Worker, workerAmount)

	for idx, _ := range p {
		p[idx] = Create(ctx)
	}

	return p
}
