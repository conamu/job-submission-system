package simulation

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/client/pkg/client"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"log/slog"
	"sync"
	"time"
)

func SimulateClient(ctx context.Context, client client.Client, wg *sync.WaitGroup) {
	wg.Add(1)
	simulatedData := `
{
    "payload": "some cool payload"
}
`
	l := logger.New(slog.LevelInfo, "client")
	jobId := ""
	var err error

	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			jobId, err = client.CreateJob(ctx, simulatedData)
			if err != nil {
				l.With("error", err.Error()).Error("error when creating job")
			}

			status := ""

			for status != "COMPLETED" {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					time.Sleep(time.Second)

					status, err = client.GetJobStatus(ctx, jobId)
					if err != nil {
						l.With("error", err.Error()).Error("error when creating job")
					}
				}
			}
		}

	}

}
