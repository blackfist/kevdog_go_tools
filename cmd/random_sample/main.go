// Package main is a program that takes a list of things
// from stdin and returns a randomly-selected sublist
package main

import (
	"os"
	"fmt"
	"bufio"
	"math/rand"
	"time"
	"flag"
)

func minInt(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}

func printResults(values []string, count int) {
	if count > 0 {
		for i:=0; i < minInt(count, len(values)); i++ {
			fmt.Println(values[i])
		}
	} else {
		for i:=0; i<len(values); i++ {
			fmt.Println(values[i])
		}
	}
}

func main() {
	// initialize the command line flags
	numbPtr := flag.Int("n", -1, "number of results desired. Default all")
	flag.Parse()

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeNamedPipe != 0) {
		values := make([]string, 0)
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			values = append(values, reader.Text())
		}
		
		// initialize random and get an array of random indexes
		r := rand.New(rand.NewSource(time.Now().Unix()))
		shuffled := make([]string, 0)

		for _, index := range r.Perm(len(values)) {
			shuffled = append(shuffled, values[index])
		}

		printResults(shuffled, *numbPtr)

	} else {
		fmt.Println("you gotta pipe in some shit, fool")
		fmt.Println("e.g. cat somefile | random_sample")
		fmt.Println("e.g. pbpaste | cut -d\" \" -f2 | random_sample")
	}
}