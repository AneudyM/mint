package main

import (
	"context"
	"fmt"
	cdp "github.com/knq/chromedp"
	"github.com/knq/chromedp/runner"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

var cmdServe = &cobra.Command{
	Use:   "run",
	Short: "Sirve el proyecto en tu browser. ",
	Long: `El subcomando run lanza el ultimo build de tu proyecto desde el directorio "/build".
	Mint-Web automaticamente lanza el browser predefinido.`,
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

func myWait() {
	time.Sleep(200 * time.Second)
}
