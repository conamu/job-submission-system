package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/conamu/job-submission-system/src/internal/pkg/logger"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/constant"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/handler"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/job"
	"github.com/conamu/job-submission-system/src/internal/server/pkg/worker"
	"github.com/spf13/viper"
	"log/slog"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

type Application interface {
	Run()
}
type application struct {
	ctx   context.Context
	log   *slog.Logger
	wg    *sync.WaitGroup
	pool  worker.Pool
	queue *job.Queue
}

func Create() Application {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	l := logger.New(slog.LevelDebug, "JobServer")
	ctx = context.WithValue(ctx, constant.CTX_LOGGER, l)
	wg := &sync.WaitGroup{}
	ctx = context.WithValue(ctx, constant.CTX_WG, wg)
	queue := job.CreateQueue()
	ctx = context.WithValue(ctx, constant.CTX_QUEUE, queue)
	pool := worker.CreatePool(ctx, viper.GetInt("server.workers"))
	ctx = context.WithValue(ctx, constant.CTX_POOL, pool)

	return &application{
		ctx:   ctx,
		log:   l,
		wg:    wg,
		pool:  pool,
		queue: queue,
	}
}

func (a *application) Run() {
	fmt.Println("Hello World!")

	h := handler.RegisterRoutes()

	srv := http.Server{
		Addr:        ":8080",
		Handler:     h,
		BaseContext: a.getBaseCtxFn,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.log.With("error", err.Error()).Error("error in http server")
		}
	}()

	<-a.ctx.Done()
	a.log.Info("Shutting down application...")
	err := srv.Shutdown(context.Background())
	if err != nil {
		a.log.With("error", err.Error()).Error("error when shutting down http server")
	}
	a.log.Info("waiting on remaining processes to finish...")
	a.wg.Wait()
}

func (a *application) getBaseCtxFn(config net.Listener) context.Context {
	return a.ctx
}
