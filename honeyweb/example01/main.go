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

func (c *fooController) Post() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Your post has been received!")
}

func main() {
	honey.MyApp.Handlers.Add("/", &fooController{})
	config := &honey.Config{
		HTTPAddr:     "localhost",
		HTTPPort:     9090,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	honey.Run(config)
}
