// Package main is a program that takes a list of things
// from stdin and returns a list formatted for a splunk query
// also allows user to specify a prefix for each item in the list
// e.g. app_id=
package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
)

func main() {
	// initialize the command line flags
	prefixPtr := flag.String("prefix", "", "a field name that will be put before each item")
	flag.Parse()

	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeNamedPipe != 0) {
		values := make([]string, 0)
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			values = append(values, reader.Text())
		}
		fmt.Printf("(")
		for i:=0;i<len(values)-1;i++ {
			fmt.Printf("%s%s OR ", *prefixPtr, values[i])
		}
		fmt.Printf("%s%s)\n", *prefixPtr, values[len(values)-1])
	} else {
		fmt.Println("you gotta pipe in some shit, fool")
		fmt.Println("e.g. cat somefile | splunk_list")
		fmt.Println("e.g. pbpaste | cut -d\" \" -f2 | splunk_list -prefix=app_id")
	}
}