package main

import (
	"log"
	"net/http"
)

const message = "Hello SSL"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	})
	err := http.ListenAndServe(":8085", mux)
	if err != nil {
		log.Fatal(err)
	}

}
