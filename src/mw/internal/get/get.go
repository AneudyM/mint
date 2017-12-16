package get

import (
	"fmt"
	"mw/internal/cmd"
)

var CmdGet = &cmd.Command{
	CmdName:  "get",
	CmdUsage: "usage: mw get [library]",
	Run:      getLibrary,
}

func getLibrary(c *cmd.Command) {
	fmt.Println("You invoked the 'get' command")
}

/*
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
*/
