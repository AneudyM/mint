package main

import (
	"flag"
	"fmt"
)

func main() {

	// Parse the command-line into defined flags
	flag.Parse()

	// Store the CLI args
	args := flag.Args()

	fmt.Println(args)

}
