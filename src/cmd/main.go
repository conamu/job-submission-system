package main

import "gitlab.ho-me.zone/conamu/job-submission-system/src/pkg/app"

func main() {
	c := &app.Config{}

	a := app.Create(c)

	a.Run()
}
