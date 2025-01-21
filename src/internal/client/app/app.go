package app

import "fmt"

type Application interface {
	Run()
}
type application struct {
}

func Create() Application {
	return &application{}
}

func (a *application) Run() {
	fmt.Println("Hello World!")
}
