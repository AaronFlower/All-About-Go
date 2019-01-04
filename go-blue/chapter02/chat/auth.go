package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"

	"strings"

	"github.com/stretchr/gomniauth"
	gomniauthcommon "github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

// ChatUser defines a chat user in the room
type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || err != nil || cookie.Value == "" {
		// no authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

// MustAuth wraps a handler which needs authorization.
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s:%s", provider, err), http.StatusBadRequest)
			return
		}

		loginURL, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s", provider, err),
				http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", loginURL)
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error when trying to get provider %s:%s", provider, err),
				http.StatusBadRequest,
			)
			return
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error when trying to complete auth for %s:%s", provider, err),
				http.StatusBadRequest,
			)
			return
		}
		user, err := provider.GetUser(creds)
		fmt.Printf("user = %+v\n", user)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error when trying to GetUser for %s:%s", provider, err),
				http.StatusBadRequest,
			)
			return
		}

		chatUser := &chatUser{User: user}
		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))
		chatUser.uniqueID = fmt.Sprintf("%x", m.Sum(nil))
		avatarURL, err := avatars.GetAvatarURL(chatUser)
		if err != nil {
			log.Fatalln("Error when trying to GetAvatarURL", "-", err)
		}
		authCookieValue := objx.New(map[string]interface{}{
			"userID":     chatUser.uniqueID,
			"name":       user.Name(),
			"avatar_url": avatarURL,
			"email":      user.Email(),
		}).MustBase64()
		fmt.Printf("authCookieValue = %+v\n", authCookieValue)

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth aciton %s not supported", action)
	}

}
