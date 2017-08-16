package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"context"
	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/runner"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Serve the current project",
	Long: `			Serve the page in the build directory of the current project's
			working directory.`,
	Run: servePage,
}

func servePage(cmd *cobra.Command, args []string) {
	fmt.Println("Serving site at http://localhost:8080")

	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := cdp.New(ctxt, cdp.WithRunnerOptions(
		runner.Path("/opt/google/chrome/chrome"),
	))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Run(ctxt, goToSite())
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", http.FileServer(http.Dir("./build")))
}

func init() {
	RootCmd.AddCommand(cmdServe)
}

func goToSite() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate("localhost:8080"),
	}
}