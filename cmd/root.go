package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "mw <COMMAND> <ARGS>",
	Short: "A simple web front-end builder tool",
	Long: `Usage: mw <COMMAND> <ARGS...>	`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
