package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace traces log
	Trace *log.Logger
	// Info traces log
	Info *log.Logger
	// Warning  traces log
	Warning *log.Logger
	// Error traces log
	Error *log.Logger
)

// Init initializes the log system.
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// to config the default log.
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	log.Println("Hello log!")

	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	defer f.Close()
	// You can also have the logger write to multiple destinations at the same time.
	multi := io.MultiWriter(f, os.Stdout)
	Init(ioutil.Discard, f, multi, f)
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
