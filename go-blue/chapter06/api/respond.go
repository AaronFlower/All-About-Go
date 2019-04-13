package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	// 创建一个 json Decoder 把 r.Body reader 中读取内容解析到 v 中
	return json.NewDecoder(r.Body).Decode(v)
}

func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// 创建一个 json Encoder 把 v 的内容编码到 w writer 中.
	return json.NewEncoder(w).Encode(v)
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		encodeBody(w, r, data)
	}
}

func respondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func respondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	respondErr(w, r, status, http.StatusText(status))
}
