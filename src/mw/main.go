package main

import (
	"flag"
	"fmt"
	"mw/internal/cmd"
)

func main() {

	// Parse the command-line into defined flags
	flag.Parse()

	// Store the CLI args
	args := flag.Args()

	// If no command is specified print the Usage
	if len(args) < 1 {
		usage()
	}

}

func usage() {
	fmt.Println("usage: mw [command] [arguments...]")
}
