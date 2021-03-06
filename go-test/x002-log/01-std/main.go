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
	// ioutil.Discard is a null device where all write calls succeed without doing anything.
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
