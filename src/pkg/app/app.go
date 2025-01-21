package app

import "fmt"

type Application interface {
	Run()
}

type Config struct {
}
type application struct {
}

func Create(c *Config) Application {
	return &application{}
}

func (a *application) Run() {
	fmt.Println("Hello World!")
}
