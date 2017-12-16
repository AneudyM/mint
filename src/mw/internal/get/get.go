package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var cmdGet = &cobra.Command{
	Use:   "get <Library Name>",
	Short: "Gets a dependency library remotely",
	Long: `Download a dependency library for your project
from the NPM registry.`,
	Run: cmdGetLibrary,
}

func cmdGetLibrary(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		fmt.Println("You need to specify a project name.")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("Specify only one file.")
		os.Exit(1)
	}
	libraryName := args[0]
	fmt.Println(libraryName)
	htmlFile, err := ioutil.ReadFile(libraryName)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("netFile.html", htmlFile, 0666)
	fmt.Println(string(htmlFile), "Hello")
}

func init() {
	RootCmd.AddCommand(cmdGet)
}
