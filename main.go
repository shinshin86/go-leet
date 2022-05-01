package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/shinshin86/go-leet/leet"
)

func flagUsage() {
	fmt.Println(`go-leet is leetspeak converter.

Usage:
go-leet {input text}
go-leet help
`)
}

func main() {
	flag.Usage = flagUsage

	var input string
	if len(os.Args) == 2 {
		input = os.Args[1]

		if input == "help" {
			flag.Usage()
			os.Exit(0)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	leetStr := Leet(input)
	fmt.Println(leetStr)
}
