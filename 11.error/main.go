package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServ:", err)
		os.Exit(1)
	}
	fmt.Println("The server is listening at :9090")
}
