package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loginfo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer log.Printf("%s %s %v\n", r.Method, r.URL, time.Since(start))
		handler.ServeHTTP(w, r)
	})
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world\n"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %s", curTime)))
}

func main() {
	addr := ":8082"
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", helloHandler)
	mux.HandleFunc("/v1/time", currentHandler)

	wrapperMux := loginfo(mux)

	err := http.ListenAndServe(addr, wrapperMux)
	if err != nil {
		log.Fatal(err)
	}
}
