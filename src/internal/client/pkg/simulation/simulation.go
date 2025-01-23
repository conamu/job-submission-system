package simulation

import (
	"context"
	"github.com/conamu/job-submission-system/src/internal/client/pkg/client"
	"github.com/conamu/job-submission-system/src/internal/pkg/constants"
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
	jobIds := make(map[string]string, 3)

	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			for i := 0; i < 3; i++ {
				jobId, err := client.CreateJob(ctx, simulatedData)
				if err != nil {
					l.With("error", err.Error()).Error("error when creating job")
				}
				jobIds[jobId] = string(constants.JOB_PENDING)
			}

			allDone := false

			for !allDone {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					for id, _ := range jobIds {
						time.Sleep(time.Second)

						status, err := client.GetJobStatus(ctx, id)
						if err != nil {
							l.With("error", err.Error()).Error("error when creating job")
						}
						jobIds[id] = status
						if status == string(constants.JOB_COMPLETED) {
							delete(jobIds, id)
						}

						if len(jobIds) == 0 {
							allDone = true
						}
					}
				}
			}
		}

	}

}
