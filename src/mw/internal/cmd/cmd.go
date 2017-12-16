package cmd

import (
	"flag"
	"fmt"
)

type Command struct {
	// Name of the command
	CmdName string

	// Single line usage of command
	CmdUsage string

	// Flags for this command
	CmdFlag flag.FlagSet

	// Function to run for this command
	CmdRun func(c *Command, args []string)
}

// Name returns the name of this command.
func (c *Command) Name() string {
	return c.Name()
}

// Usage prints the single line usage for
// this command.
func (c *Command) Usage() {
	fmt.Println(c.CmdUsage)
}
