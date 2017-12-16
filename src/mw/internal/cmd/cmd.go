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
}

func (c *Command) Name() string {
	return c.Name()
}

func (c *Command) Usage() {
	fmt.Println(c.CmdUsage)
}
