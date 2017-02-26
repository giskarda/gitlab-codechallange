package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var top_numbersFlag int
var filenameFlag string

func init() {
	flag.IntVar(&top_numbersFlag, "top-numbers", 0, "topN wanted to retrieve")
	flag.StringVar(&filenameFlag, "filename", "numbers.txt", "file to search for top numbers")
}

func check_number(final []int, number int) []int {
	for i, v := range final {
		if number > v {
			// https://github.com/golang/go/wiki/SliceTricks
			// pop
			_, final := final[len(final)-1], final[:len(final)-1]
			// insert
			final = append(final[:i], append([]int{number}, final[i:]...)...)
			break
		}
	}
	return final
}

func build_slice(top_numbers int) []int {
	log.Printf("Building slice of %d elements", top_numbers)
	var final []int

	for v := top_numbers; v > 1; v-- {
		final = append(final, v)
	}

	return final
}

func main() {
	flag.Parse()

	final := build_slice(top_numbersFlag)

	start := time.Now()

	file, err := os.Open(filenameFlag) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting to read file %s", filenameFlag)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		final = check_number(final, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	end := time.Now()
	fmt.Printf("The call took %v to run.\n", end.Sub(start))

	for _, v := range final {
		log.Println(v)
	}

}
