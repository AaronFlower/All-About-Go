package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloHandler)
	f, err := os.Create("./foo.txt")
	if err != nil {
		log.Fatalf("Create file failed: %v", err)
	}
	defer func () {
		err := f.Close()
		log.Fatalf("Close file failed: %v", err)
	}()

	str := "foo bar"
	_, _ = f.WriteString(str)

	err = http.ListenAndServe(":9097", nil)
	if err != nil {
		log.Fatal("Starting server failed")
	}
	fmt.Println("The server is listening at: 9097")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	str := "Hello world!"
	_, _ = w.Write([]byte(str))
}
