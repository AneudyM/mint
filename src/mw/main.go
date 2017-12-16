package main

import (
	"flag"
	"fmt"
	"mw/internal/build"
	"mw/internal/cmd"
	"mw/internal/get"
	"mw/internal/new"
	"mw/internal/run"
)

func init() {
	cmd.Commands = []*cmd.Command{
		run.CmdRun,
		build.CmdBuild,
		get.CmdGet,
		new.CmdNew,
	}
}

func main() {

	// Parse the command-line into defined flags
	flag.Parse()

	// Store the CLI args
	args := flag.Args()

	// If no command is specified print the Usage
	if len(args) < 1 {
		usage()
	}

	// Evaluate CLI's arguments against list of
	// available commands
	for _, c := range cmd.Commands {
		if c.CmdName == args[0] {
			c.Run(c)
		}
	}

}

func usage() {
	fmt.Println("usage: mw [command] [arguments...]")
}
