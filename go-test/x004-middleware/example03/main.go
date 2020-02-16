package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Adapter defines a adapter function takes a handler as
// parameter and returns a handler.
// The Adapter gets its name from the adapter pattern --
// also known as the decorate pattern.
type Adapter func(http.Handler) http.Handler

// Headerit provides set headers
func Headerit(key, value string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(key, value)
			h.ServeHTTP(w, r)
		})
	}
}

// Logify writes log info for a request.
func Logify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Logger 1: before")
			defer logger.Println("Logger 1: after")
			h.ServeHTTP(w, r)
		})
	}
}

// Logtwo writes log info for a request.
func Logtwo(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Logger 2: before")
			defer logger.Println("Logger 2: after")
			h.ServeHTTP(w, r)
		})
	}
}

// Adapt takes adapters to the specified handler.
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, a := range adapters {
		h = a(h)
	}
	return h
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Write([]byte(fmt.Sprintf("The current time is %s ", start)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", helloFunc)
	mux.HandleFunc("/v1/test", testFunc)

	logger := log.New(os.Stdout, "server:", log.Lshortfile)

	// wrappedMux := Logify(logger)(mux)
	// 注意添加顺序和执行相反.
	wrappedMux := Adapt(mux,
		Logtwo(logger),
		Logify(logger),
		Headerit("Custom-Header-Key", "Foo"),
	)
	err := http.ListenAndServe(":8083", wrappedMux)
	if err != nil {
		log.Fatal(err)
	}
}
