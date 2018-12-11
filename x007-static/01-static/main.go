package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)
	fmt.Println("The server is listening at: 8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
