package main

import (
	"fmt"
	"time"

	"github.com/aaronflower/honey"
)

type barController struct {
	honey.Controller
}

func (c *barController) Get() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Hello Again")
}

func (c *barController) Post() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Your post has been handled.")
}

func main() {
	honey.MyApp.Handlers.Add("/", &barController{})
	config := &honey.Config{
		HTTPAddr:     "localhost",
		HTTPPort:     9090,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	honey.Run(config)
}
