package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	userID := req.FormValue("userID")
	file, header, err := req.FormFile("avatarFile")
	log.Printf("userID = %+v\n", userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := path.Join("avatars", userID+path.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, "Successful")
}
