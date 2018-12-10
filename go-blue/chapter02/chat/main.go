package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"html/template"

	"github.com/aaronflower/ago/go-blue/chapter02/trace"
)

type templ struct {
	filename string
	once     sync.Once
	tmpl     *template.Template
}

func (t *templ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles("templates/" + t.filename))
	})
	t.tmpl.Execute(w, r)
}

func main() {
	room := newRoom()
	mux := http.NewServeMux()
	mux.Handle("/", &templ{filename: "chat.html"})
	mux.Handle("/room", room)
	room.tracer = trace.New(os.Stdout)
	go room.run()
	fmt.Println("The server is listening at ")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
