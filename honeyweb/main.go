package main

import (
	"fmt"
	"time"

	"github.com/aaronflower/honey"
)

type fooController struct {
	honey.Controller
}

func (c *fooController) Get() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Hello World!")
}

func main() {
	honey.MyApp.Handlers.Add("/", &fooController{})
	config := &honey.Config{
		HTTPAddr:     "localhost",
		HTTPPort:     9090,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("MyApp = %+v\n", honey.MyApp)
	honey.Run(config)
}
