package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"html/template"
)

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	if info == nil || info.IsDir() {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Fatal(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()

	// StripPrefix returns a handler that serves HTTP request by removing
	// the given prefix from the request URL's Path and invoking the handler.
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", serveTemplate)

	fmt.Println("The serve is listen at :8082")
	http.ListenAndServe(":8082", mux)
}
