package main

import (
	"fmt"
	"log"
	"net/http"
)

func uploadHanlder(w http.ResponseWriter, r *http.Request) {
	var maxMemory int64
	maxMemory = 32 << 20 // 32 MB
	r.ParseMultipartForm(maxMemory)
	fmt.Println("-->")
	fmt.Printf("r = %+v\n", r.MultipartForm)
	fmt.Println("<--")
	w.WriteHeader(204)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/upload", uploadHanlder)
	fmt.Println("The server is listening at :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
