// Package main is a program that takes a list of things
// from stdin and downcases all the items and creates a
// list suitable for pasting into an sql statement
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	info, _ := os.Stdin.Stat()
	if (info.Mode() & os.ModeNamedPipe != 0) {  
		values := make([]string, 0)
		reader := bufio.NewScanner(os.Stdin)
		fmt.Print("(")
		for reader.Scan() {
			var buffer bytes.Buffer
			buffer.WriteString("'")
			buffer.WriteString(strings.ToLower(reader.Text()))
			buffer.WriteString("'")
			values = append(values, buffer.String())
		}
		
		fmt.Print(strings.Join(values, ","))
		fmt.Print(")")
		fmt.Print("\n")
	} else {
		fmt.Println("you gotta pipe in some shit, fool")
		fmt.Println("e.g. cat somefile | sql_list")
		fmt.Println("e.g. pbpaste | cut -d\" \" -f2 | sql_list")
	}
	
}
