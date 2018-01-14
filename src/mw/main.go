package main

import (
	"flag"
	"fmt"
	"mw/internal/build"
	"mw/internal/cmd"
	"mw/internal/get"
	"mw/internal/new"
	"mw/internal/run"
	"os"
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

	flag.Usage = cmd.Usage

	// Parse the command-line into defined flags
	flag.Parse()

	// Store the CLI args
	args := flag.Args()

	// If no command is specified print the Usage
	if len(args) < 1 {
		os.Exit(1)
	}

	if args[0] == "help" {
		mwUsage()
		return
	}
	// Evaluate CLI's arguments against list of
	// available commands
	for _, cmd := range cmd.Commands {
		fmt.Println(cmd)
	}
}

func init() {
	cmd.Usage = mwUsage
}

func mwUsage() {
	fmt.Println("usage: mw [command] [arguments...]")
	os.Exit(1)
}
