package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world\n"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %s", curTime)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", helloHandler)
	mux.HandleFunc("/v1/time", currentHandler)

	err := http.ListenAndServe("localhost:8091", mux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The serve is listen at :8091")
}
