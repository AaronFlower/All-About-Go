package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details.
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world\n"))
}

func currentHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %s", curTime)))
}

func main() {
	addr := ":8080"
	fmt.Println("Run....")
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", helloHandler)
	mux.HandleFunc("/v1/time", currentHandler)

	wrappedMux := NewLogger(mux)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is listening at %s", addr)
	log.Printf("Server is listening at %s", addr)
}
