package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	fmt.Println("The server is listening at :80")
	log.Fatal(http.ListenAndServe(":80", mux))
}
