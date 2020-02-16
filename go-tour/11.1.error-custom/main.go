package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (f appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := f(w, r); e != nil { // e is *appError, not os.Error.
		http.Error(w, e.Message, e.Code)
	}
}

// type HandlerFunc func(ResponseWriter, *Request)

// // ServeHTTP calls f(w, r).
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }

func testFunc(w http.ResponseWriter, r *http.Request) *appError {
	r.ParseForm()
	code := r.FormValue("code")
	switch code {
	case "500":
		return &appError{nil, "Internal Server Error!", 500}
	case "404":
		return &appError{nil, "Resources not found!", 404}
	default:
		return &appError{nil, "Unknown Error", 500}
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	h := appHandler(testFunc)

	http.Handle("/test", h)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServ:", err)
		os.Exit(1)
	}
	fmt.Println("The server is listening at :9090")
}
