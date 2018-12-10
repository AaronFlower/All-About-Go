package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"html/template"
)

// Templ represents a single template
type Templ struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request
func (t *Templ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()
	// root
	r := NewRoom()
	// r.tracer = trace.New(os.Stdout)
	http.Handle("/", &Templ{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going, running the room in a separate goroutine.
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
