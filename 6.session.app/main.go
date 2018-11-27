package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"html/template"
	textTemplate "text/template"

	session "github.com/aaronflower/ago/6.session"
	_ "github.com/aaronflower/ago/6.session/providers"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_log"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello http!")
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")

	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < time.Now().Unix() {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()

	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")

		data := struct {
			UserName string
			Md5      string
		}{
			"",
			token,
		}
		if name, ok := sess.Get("username").(string); ok {
			data.UserName = name
		}
		t.Execute(w, data)
	} else {
		sess.Set("username", r.FormValue("username"))
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password:", r.FormValue("password"))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))

		t, err := textTemplate.New("Foo").Parse(`{{define "T"}} Hello, {{.}}!{{end}}`)
		if err != nil {
			log.Fatal(" pass error")
		}
		err = t.ExecuteTemplate(w, "T", r.Form.Get("username"))
	}
}

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "go-session-id", 3600)
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)

	// t := time.Date(2018, time.November, 10, 26, 0, 0, 0, time.UTC)
	// fmt.Printf("Go launched = %+v\n", t.Local())

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
