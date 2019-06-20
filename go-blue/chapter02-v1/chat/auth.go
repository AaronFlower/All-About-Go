package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check code
	h.next.ServeHTTP(w, r)
}

// MustAuth wraps a auth check handler for the next handler.
func MustAuth(next http.Handler) http.Handler {
	return &authHandler{
		next: next,
	}
}
