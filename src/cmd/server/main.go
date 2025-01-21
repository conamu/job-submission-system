package main

import (
	"github.com/conamu/job-submission-system/src/internal/pkg/config"
	"github.com/conamu/job-submission-system/src/internal/server/app"
)

func main() {
	config.Init()

	c := &app.Config{}

	a := app.Create(c)

	a.Run()
}
