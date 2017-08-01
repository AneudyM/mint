package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cmdGet = &cobra.Command{
	Use:   "get <Library Name>",
	Short: "Gets a dependency library remotely",
	Long: `Download a dependency library for your project
from the NPM registry.`,
	Run: cmdGetLibrary,
}

func cmdGetLibrary(cmd *cobra.Command, args []string) {
	libraryName := args[0]
	fmt.Println("Downloading " + libraryName + "...")
}

func init() {
	RootCmd.AddCommand(cmdGet)
}