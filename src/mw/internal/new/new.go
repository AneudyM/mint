package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"mw/internal/cmd"
)

var CmdNew = &cmd.Command{
	CmdName:    "new",
	CmdUsage:   "usage: mw new <project-name>",
	HasNoFlags: true,
	Run:        createProject,
}

var pStruct = [...]string{"src", "build", "img", "js", "css", "templates"}
var srcDir = pStruct[0]
var buildDir = pStruct[1]

type IndexTemplate struct {
	Title       string
	MintVersion string
}

const indexTemplate = `<!DOCTYPE html>
	<html class="no-js" lang="en">
		<head>
			<meta charset="utf-8">
       		<meta http-equiv="x-ua-compatible" content="ie=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<meta name="generator" content="Mint Web 0.0.2" />
			<title>Mint Web Page</title>
			<link rel="stylesheet" href="css/main.css">
		</head>
		<body>
			<h1>Hello, Mint Web!</h1>

			<script src="js/main.js"></script>
		</body>
	</html>
`

func createProject(c *cmd.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("No project name specified.")
	}

	if len(args) > 1 {
		fmt.Printf(c.CmdUsage + "\n")
		os.Exit(2)
	}

	projectName := args[0]
	err := os.Mkdir(projectName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(filepath.Join(projectName, srcDir), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(filepath.Join(projectName, buildDir), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range pStruct[2:] {
		err := os.Mkdir(filepath.Join(projectName, srcDir, dir), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	file := filepath.Join(projectName, srcDir, "index.html")
	cf, err := os.Create(file)
	defer cf.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = cf.WriteString(indexTemplate)
	if err != nil {
		log.Fatal(err)
	}
}
