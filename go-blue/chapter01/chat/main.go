package main

import (
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
	t.templ.Execute(w, nil)
}

func main() {

	r := NewRoom()
	// root
	http.Handle("/", &Templ{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going, running the room in a separate goroutine.
	go r.Run()

	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
