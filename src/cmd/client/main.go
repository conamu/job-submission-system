package main

import (
	"github.com/conamu/job-submission-system/src/internal/client/app"
	"github.com/conamu/job-submission-system/src/internal/pkg/config"
)

func main() {
	config.Init()
	a := app.Create()
	a.Run()
}
