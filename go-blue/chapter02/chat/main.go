package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"html/template"

	"github.com/aaronflower/ago/go-blue/chapter02/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	yaml "gopkg.in/yaml.v2"
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
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

// Provider defines a app provider for OAuth2
type Provider struct {
	Key      string `json:"key"`
	Secret   string `json:"secret"`
	Callback string `json:"callback"`
}

// ProviderMap defines a app provider map
type ProviderMap map[string]Provider

func loadConfig(providers *ProviderMap) error {
	configSource, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configSource, &providers)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	var providers ProviderMap
	err := loadConfig(&providers)
	if err != nil {
		log.Fatal(err)
		return
	}

	// setup gomniauth
	gomniauth.SetSecurityKey("I am the security key")
	gomniauth.WithProviders(
		google.New(
			providers["google"].Key,
			providers["google"].Secret,
			providers["google"].Callback,
		),
		facebook.New(
			providers["facebook"].Key,
			providers["facebook"].Secret,
			providers["facebook"].Callback,
		),
	)
	// root
	r := NewRoom()
	r.tracer = trace.New(os.Stdout)

	mux := http.NewServeMux()
	mux.Handle("/", MustAuth(&Templ{filename: "chat.html"}))
	mux.Handle("/login", &Templ{filename: "login.html"})
	mux.HandleFunc("/auth/", loginHandler)
	mux.Handle("/room", r)

	// get the room going, running the room in a separate goroutine.
	go r.Run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
