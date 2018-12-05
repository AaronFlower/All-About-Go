package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log for json.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	// log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Warn("The ice breaks!")

	// 通过日志语句重用字段
	// logrus.Entry返回自WithFields()
	// WithFields 可以公用。在下一次记录时也会使用。
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I'll be logged with common and other field")
	contextLogger.Warn("Me too")
}
