package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cmdNewProject = &cobra.Command{
	Use:   "new <Project Name>",
	Short: "Creates a new Web project directory structure",
	Long:  `Creates a new Web Project in the current directory`,
	Run:   newProject,
}

var perm os.FileMode = 0755
var indexContent string = `<!doctype html>
<html class="no-js" lang="">
	<head>
		<meta charset="utf-8">
        <meta http-equiv="x-ua-compatible" content="ie=edge">
		<title>Mint Page</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="css/main.css">
	</head>
	<body>
		<h1>Hello, Mint Web!</h1>

	<script src="js/main.js"></script>
	</body>
</html>
`

type project struct {
	projectName string
	buildDir    string
	srcDir      string
	libsDir     string
	cssDir      string
	imgDir      string
	includesDir string
	jsDir       string
}

func newProject(cmd *cobra.Command, args []string) {
	var newProject project
	if len(args) <= 0 {
		fmt.Println("You need to specify a project name.")
		os.Exit(1)
	}
	newProject.projectName = args[0]
	newProject.buildDir = newProject.projectName + "/build" // built files for publication
	newProject.srcDir = newProject.projectName + "/src"     // Source files
	newProject.libsDir = newProject.projectName + "/libs"   // Downloaded libraries land here
	newProject.cssDir = newProject.srcDir + "/css"
	newProject.imgDir = newProject.srcDir + "/img"
	newProject.includesDir = newProject.srcDir + "/includes"
	newProject.jsDir = newProject.srcDir + "/js"
	err := os.Mkdir(newProject.projectName, perm)
	if os.IsExist(err) {
		fmt.Println("The project " + "\"" + newProject.projectName + "\"" + " already exists.")
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.buildDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.libsDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.srcDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.cssDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.imgDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.includesDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(newProject.jsDir, perm)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(newProject.jsDir + "/main.js")
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(newProject.cssDir + "/main.css")
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Create(newProject.srcDir + "/index.html")
	if err != nil {
		log.Fatal(err)
	}
	indexFile, err := os.OpenFile(newProject.srcDir+"/index.html", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = indexFile.WriteString(indexContent)
}

func init() {
	RootCmd.AddCommand(cmdNewProject)
}
