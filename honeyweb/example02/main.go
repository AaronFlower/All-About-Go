package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/aaronflower/ago/honeyweb/example02/route"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/aaronflower/honey"
)

var env = viper.New()
var log = logrus.New()

func init() {
	env.SetConfigName("env")
	env.AddConfigPath("./config/")
	err := env.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

type barController struct {
	honey.Controller
}

// func (c *barController) Init() {
// 	fmt.Println("Your should create a log file for bar!")
// }

func (c *barController) Get() {
	fmt.Fprintf(c.Ct.ResponseWriter, "Hello Again")
	r := c.Ct.Request
	fmt.Printf("c = %+v\n", r)
	fmt.Printf("c = %+v\n", r.URL)
	fmt.Printf("c = %+v\n", r.RequestURI)
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
	honey.Run(env, config)
}
