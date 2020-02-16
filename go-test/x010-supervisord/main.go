package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var name, password string
	flag.StringVar(&name, "name", "k", "The user name")
	flag.StringVar(&password, "password", "b", "The user password")
	flag.Parse()
	fmt.Printf("name = %+v\n", name)
	fmt.Printf("password = %+v\n", password)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The name is %s and password is %s ", name, password)
	})

	fmt.Println("The server is listening at :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
