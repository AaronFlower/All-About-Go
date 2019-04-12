package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bitly/go-nsq"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func doCount(countsLock *sync.Mutex, counts *map[string]int, pollData *mgo.Collection) {
	countsLock.Lock()
	defer countsLock.Unlock()

	if len(*counts) == 0 {
		log.Println("No new votes, skipping database update")
		return
	}

	log.Println("Updating database... ")
	log.Println(*counts)
	ok := true
	for option, count := range *counts {
		sel := bson.M{"options": bson.M{"$sin": []string{option}}}
		up := bson.M{"$inc": bson.M{"results." + option: count}}
		if _, err := pollData.UpdateAll(sel, up); err != nil {
			log.Println("failed to update:", err)
			ok = false
		}
	}
	if ok {
		log.Println("Finished updating database...")
		*counts = nil // reset counts
	}

}

const updateDuration = 1 * time.Second

func main() {
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("Connection to database...")
	db, err := mgo.Dial("mongodb://127.0.0.1:27017/localhost")
	if err != nil {
		fatal(err)
		return
	}

	defer func() {
		log.Println("Closing database connection...")
		db.Close()
	}()

	pollData := db.DB("ballots").C("polls")

	var counts map[string]int
	var countsLock sync.Mutex

	log.Println("Connecting to nsq...")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}

	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
		defer countsLock.Unlock()
		if counts == nil {
			counts = make(map[string]int)
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))

	// 我们能过链接 http 服务 nsqlookupd 而不是直接去连接 nsqd
	if err := q.ConnectToNSQLookupd("localhost:4161"); err != nil {
		fatal(err)
		return
	}

	// Responding to Ctrl+C 处理 Ctrl+C
	ticker := time.NewTicker(updateDuration)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case <-ticker.C:
			doCount(&countsLock, &counts, pollData)
		case <-termChan:
			ticker.Stop()
			q.Stop()
		case <-q.StopChan:
			// finished
			return
		}
	}
}
