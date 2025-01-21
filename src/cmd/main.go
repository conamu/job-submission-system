package main

import "github.com/conamu/job-submission-system/src/pkg/app"

func main() {
	c := &app.Config{}

	a := app.Create(c)

	a.Run()
}
