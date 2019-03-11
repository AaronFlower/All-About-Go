package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/AaronFlower/All-About-Go/go-programming-qiniu/02-sorter/src/algorithms/bubblesort"
	"github.com/AaronFlower/All-About-Go/go-programming-qiniu/02-sorter/src/algorithms/qsort"
)

// read values for infile
func readValues(infile string) ([]int, error) {
	values := []int{}
	file, err := os.Open(infile) // Open 是对 OpenFile 的另一个包装，打开的是只读文件。
	if err != nil {
		return values, err
	}
	defer file.Close()

	// file 实现了 Read 方法就可以用来被初始化为一个 Reader
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return values, err
		}
		values = append(values, v)
	}

	return values, err
}

func writeValues(outfile string, values []int) error {
	file, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, v := range values {
		str := strconv.Itoa(v)
		file.WriteString(str + "\n")
	}

	return nil
}

// 你应该熟悉 `os, io, bufio, strconv, flag` 的标准库的用法，会大大提高你的效率。

func main() {
	var (
		infile  string
		outfile string
		algo    string
	)
	flag.StringVar(&infile, "i", "infile", "File Contains values for sorting")
	flag.StringVar(&outfile, "o", "outfile", "File to receive sorted values")
	flag.StringVar(&algo, "a", "qsort", "Sort algorightm")

	flag.Parse()

	fmt.Printf("infile = %+v\n", infile)
	fmt.Printf("outfile = %+v\n", outfile)
	fmt.Printf("algo = %+v\n", algo)

	values, err := readValues("./" + infile)

	if err != nil {
		log.Fatal(err)
	}

	t1 := time.Now()
	switch algo {
	case "qsort":
		qsort.QuickSort(values)
	case "bubble":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("Sorting alogrithm ", algo, " is either unknown or unsupported.")
	}
	t2 := time.Now()
	fmt.Println("The sorting process costs", t2.Sub(t1), " to complete.")

	fmt.Printf("Input values = %+v\n", values)
	fmt.Printf("Sorted files = %+v\n", values)

	writeValues(outfile, values)
}
