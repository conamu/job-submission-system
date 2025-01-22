package app

import (
	"context"
	"fmt"
	"github.com/conamu/job-submission-system/src/internal/client/pkg/client"
	"github.com/conamu/job-submission-system/src/internal/client/pkg/simulation"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sync"
)

type Application interface {
	Run()
}
type application struct {
	wg     *sync.WaitGroup
	client client.Client
}

func Create() Application {
	return &application{
		wg:     &sync.WaitGroup{},
		client: client.NewClient(),
	}
}

func (a *application) Run() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	simulatedWorkers := viper.GetInt("client.instances")
	for i := 0; i < simulatedWorkers; i++ {
		fmt.Println("started a simulated client")
		go simulation.SimulateClient(ctx, a.client, a.wg)
	}

	<-ctx.Done()
	a.wg.Wait()
}
