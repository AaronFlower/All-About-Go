package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"html/template"

	session "github.com/aaronflower/ago/6.session"
	_ "github.com/aaronflower/ago/6.session/providers"
)

func index(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	username := sess.Get("username")

	data := struct {
		UserName    string
		AccessTimes int
	}{
		"",
		0,
	}

	if name, ok := username.(string); ok {
		data.UserName = name
		data.AccessTimes = countAccessTimes(w, r)
	}

	t, _ := template.ParseFiles("index.gtpl")
	w.Header().Set("Content-Type", "text/html")

	t.Execute(w, data)
}

func countAccessTimes(w http.ResponseWriter, r *http.Request) int {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")

	if (createtime.(int64) + 3600) < time.Now().Unix() {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
		http.Redirect(w, r, "/", 302)
	}

	count := 1
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		count = ct.(int) + 1
		sess.Set("countnum", count)
	}
	return count
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
		sess.Set("createtime", time.Now().Unix())
		http.Redirect(w, r, "/", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionDestroy(w, r)
	http.Redirect(w, r, "/", 302)
}

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "go-session-id", 3600)
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	// t := time.Date(2018, time.November, 10, 26, 0, 0, 0, time.UTC)
	// fmt.Printf("Go launched = %+v\n", t.Local())

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
