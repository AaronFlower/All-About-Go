package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type file struct {
	chunkSize int
	name      string
	complete  chan bool
	blobs     map[int]string
}

var filePath = "./test"

func decodeString(s string) ([]byte, error) {
	needle := "base64,"
	i := strings.Index(s, needle)
	if i == -1 {
		return nil, errors.New("Wrong Raw data")
	}
	s = s[i+len(needle):]
	return base64.StdEncoding.DecodeString(s)
}

func (f *file) create() {
	defer delete(filesMap, f.name)
	for ok := range f.complete {
		t := time.Now()
		name := strconv.FormatInt(t.UnixNano()/int64(time.Millisecond), 16) + f.name
		newFile, err := os.OpenFile(filePath+"/"+name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer newFile.Close()
		if err != nil {
			fmt.Println("The file created failed!")
			log.Fatal(err)
			return
		}
		for i := 0; i < f.chunkSize; i++ {
			data, err := decodeString(f.blobs[i])
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = newFile.Write(data)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		fmt.Println("The file has been upload successfully!")
		fmt.Println(ok)
		return
	}
}

var filesMap = make(map[string]*file)

func uploadHanlder(w http.ResponseWriter, r *http.Request) {
	var maxMemory int64
	maxMemory = 32 << 20 // 32 MB
	r.ParseMultipartForm(maxMemory)
	filename := r.MultipartForm.Value["filename"][0]
	data := r.MultipartForm.Value["file_data"][0]
	timestamp := r.MultipartForm.Value["timestamp"][0]
	chunkSize, err := strconv.Atoi(r.MultipartForm.Value["chunk_size"][0])
	if err != nil {
		log.Fatal(err)
	}
	index, err := strconv.Atoi(r.MultipartForm.Value["index"][0])
	if err != nil {
		log.Fatal(err)
	}
	var f *file
	var ok bool
	key := filename + timestamp
	f, ok = filesMap[key]
	if ok {
		fmt.Println(key, "The file has been created!")
		f.blobs[index] = data
	} else {
		fmt.Println(key, "Create the file placeholder!")
		f = &file{
			chunkSize: chunkSize,
			name:      filename,
			complete:  make(chan bool),
			blobs:     make(map[int]string),
		}
		filesMap[key] = f
		f.blobs[index] = data
		go f.create()
	}
	if len(f.blobs) == f.chunkSize {
		f.complete <- true
	}
	w.WriteHeader(204)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/upload", uploadHanlder)
	fmt.Println("The server is listening at :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
