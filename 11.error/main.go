package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func testFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.FormValue("code")
	switch code {
	case "500":
		http.Error(w, "Internal Server Error!", 500)
	case "404":
		http.Error(w, "Resources not found!", 404)
	default:
		http.Error(w, "Unknown Error", 500)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/test", testFunc)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServ:", err)
		os.Exit(1)
	}
	fmt.Println("The server is listening at :9090")
}
