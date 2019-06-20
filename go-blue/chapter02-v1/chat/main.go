package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Templ parse the template file and return
type Templ struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (tpl *Templ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl.once.Do(func() {
		tpl.templ = template.Must(template.ParseFiles(filepath.Join("template", tpl.filename)))
	})
	tpl.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8089", "The addr fo the application")
	flag.Parse()

	r := mux.NewRouter()
	r.Handle("/", &Templ{filename: "chat.html"})

	room := NewRoom()
	// room.tracer = trace.New(os.Stdout)
	r.Handle("/room", room)

	go room.Run()

	server := &http.Server{
		Handler:      r,
		Addr:         *addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
