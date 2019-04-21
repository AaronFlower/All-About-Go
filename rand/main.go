package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/**
 * rand [options] <num>
 */
func main() {
	const MinUint = 0
	const MaxUint = ^uint(MinUint)
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -(MaxInt - 1)

	// isUniquePtr := flag.Bool("u", false, "To generate unique numbers which are repeatable by default")
	numBeginPtr := flag.Int("b", 0, "To specify the begin of {begin, end} which is zero by default")
	numEndPtr := flag.Int("e", 65535, "To specify the end of {begin, end} which is 65535")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: rand [options] <num>")
		os.Exit(1)
	}

	total, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Usage: rand [options] <num>")
		os.Exit(1)
	}

	min, max := float64(*numBeginPtr), float64(*numEndPtr)
	scope := float64(max - min)

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	for i := 0; i < total; i++ {
		fmt.Println(int(min + r.Float64()*scope))
	}
}
