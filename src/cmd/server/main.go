package main

import (
	"github.com/conamu/job-submission-system/src/internal/pkg/config"
	"github.com/conamu/job-submission-system/src/internal/server/app"
)

func main() {
	config.Init()
	a := app.Create()
	a.Run()
}
