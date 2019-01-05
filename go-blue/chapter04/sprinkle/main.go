package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord + " app",
	otherWord + " site",
	otherWord + " time",
	"get " + otherWord,
	"go " + otherWord,
	"lets " + otherWord,
	otherWord + " hq",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	length := len(transforms)
	for s.Scan() {
		t := transforms[rand.Intn(length)]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
