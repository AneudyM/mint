package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Serve the current project",
	Long: `Serve the page in the build directory of the current project's
working directory.`,
	Run: servePage,
}

func servePage(cmd *cobra.Command, args []string) {
	fmt.Println("Serving site at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./build"))))
}

func init() {
	RootCmd.AddCommand(cmdServe)
}
