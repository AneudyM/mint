/*
	Package cmd represents a command object.

*/
package cmd

import (
	"flag"
	"fmt"
)

// Commands holds the list of all the supported commands.
// This list is filled in main.go by referencing to it.
var Commands []*Command

var Usage func()

type Command struct {
	// Name of the command
	CmdName string

	// Single line showing usage
	CmdUsage string

	// Flags for the command
	CmdFlag flag.FlagSet

	// Operation the command performs
	Run func(c *Command, args []string)
}

// Name returns the name of the command.
func (c *Command) Name() string {
	return c.CmdName
}

// Usage prints the single line usage for the command.
func (c *Command) Usage() {
	fmt.Println(c.CmdUsage)
}
