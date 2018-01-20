package new

import (
	"fmt"
	"log"
	"mw/internal/cmd"
	"os"
	"path/filepath"
)

var CmdNew = &cmd.Command{
	CmdName:    "new",
	CmdUsage:   "usage: mw new <project-name>",
	HasNoFlags: true,
	Run:        createProject,
}

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

	err = os.Mkdir(filepath.Join(projectName, "src"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(filepath.Join(projectName, "build"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	srcStructure := []string{"js", "css", "img"}
	for _, dir := range srcStructure {
		err := os.Mkdir(filepath.Join(projectName, "src", dir), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*
var indexContent string = `<!DOCTYPE html>
	<html class="no-js" lang="en">
		<head>
			<meta charset="utf-8">
       		<meta http-equiv="x-ua-compatible" content="ie=edge">
			<title>Mint Page</title>
			<meta name="viewport" content="width=device-width, initial-scale=1">
			<meta name="generator" content="MintWeb 0.0.1" />
			<link rel="stylesheet" href="css/main.css">
		</head>
		<body>
			<h1>Hello, Mint Web!</h1>

			<script src="js/main.js"></script>
		</body>
	</html>
`

var perm os.FileMode = os.ModePerm

func newProject(c *cmd.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("Nombre de proyecto no especificado.")
	}
	projectName, err := filepath.Abs(filepath.Clean(args[0]))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(projectName, perm)
	if err != nil {
		log.Fatal("El proyecto ya existe.")
	}
	baseStructure := []string{
		filepath.Join(projectName, "src"),
		filepath.Join(projectName, "build"),
		filepath.Join(projectName, "libs"),
	}
	for _, dir := range baseStructure {
		err = os.Mkdir(dir, perm)
		if err != nil {
			log.Fatal(err)
		}
	}
	srcDir, _ := filepath.Abs(filepath.Clean(baseStructure[0]))
	srcStructure := []string{
		filepath.Join(srcDir, "pages"),
		filepath.Join(srcDir, "includes"),
		filepath.Join(srcDir, "img"),
		filepath.Join(srcDir, "css"),
		filepath.Join(srcDir, "js"),
	}
	for _, dir := range srcStructure {
		err = os.Mkdir(dir, perm)
		if err != nil {
			log.Fatal(err)
		}
	}
	srcTemplates := []string{
		filepath.Join(srcStructure[0], "index.html"),
		filepath.Join(srcStructure[3], "style.css"),
		filepath.Join(srcStructure[4], "main.js"),
	}
	for _, file := range srcTemplates {
		newFile, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Base(file) == "index.html" {
			newFile.WriteString(indexContent)
		}
	}
}

func init() {
	RootCmd.AddCommand(cmdNewProject)
}
*/
