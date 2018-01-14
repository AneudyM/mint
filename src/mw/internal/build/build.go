package build

import (
	"fmt"
	"log"
	"mw/internal/cmd"
	"os"
	"path/filepath"
)

var CmdBuild = &cmd.Command{
	CmdName:  "build",
	CmdUsage: "usage: mw build",
	Run:      buildProject,
}

var srcDir = "src"

func buildProject(c *cmd.Command, args []string) {
	_, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		fmt.Println("Project's src directory not found.")
		os.Exit(2)
	}

	err = filepath.Walk(srcDir, compile)
	if err != nil {
		log.Fatal(err)
	}
}

func compile(path string, info os.FileInfo, err error) error {

	return err
}

/*
func init() {
	RootCmd.AddCommand(cmdBuild)
}

func htmlCompiler(path string, fileInfo os.FileInfo, err error) error {
	//(1) Get the current directory's absolute path
	cwd, _ := os.Getwd()
	absPath := filepath.Join(cwd, path)
	buildPath := filepath.Join(cwd, "build")
	//(2) Process files according to their extension
	//(2.1) Get file extensionfunc cmdBuildProject(cmd *cobra.Command, args []string) {
	//(1) Make sure the user runs the command from its project's root directory
	_, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		log.Fatal("Asegurate de correr el build desde la raíz del proyecto.")
	}
	//(2) If user is in the root of the project, then run htmlCompiler
	filepath.Walk(srcDir, htmlCompiler)
} if the file is not a directory
	file, _ := os.Stat(absPath)
	if file.IsDir() {
		if file.Name() == "includes" {
			return filepath.SkipDir
		}
	}
	fileExtension := filepath.Ext(absPath)
	switch fileExtension {
	case ".html":
		//fmt.Println("The extension of this file", "'"+file.Name()+"'", "is", fileExtension+".")
		buildHTMLFile(absPath, buildPath)
	case ".scss":
		// Call SCSS file handler
	case ".js":
		// Call JS file handler
	}
	return nil
}

func buildHTMLFile(file string, target string) {
	var compiledString []string
	cwd, _ := os.Getwd()
	srcFile := readFile(file)
	r := strings.NewReader(srcFile)
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			// EOF is considered an error.
			// Therefore, break ends iterator whenever EOF is found
			break
		}
		token := tokenizer.Token()
		// Replace the include statement with include file contents
		if token.Data == "include" {
			// Validate attribute criteria
			if len(token.Attr) != 1 {
				log.Fatal("Has especificado más de un atributo.")
			}
			if token.Attr[0].Key != srcDir {
				log.Fatal("Atributo", "'"+token.Attr[0].Key+"'", "no reconocido.")
			}
			if token.Attr[0].Val == "" {
				log.Fatal("El atributp 'src' está vacio.")
			}
			includeFileName := token.Attr[0].Val
			// Retrive include file
			// But first make sure it exists:
			includeFilePath := filepath.Join(cwd, srcDir, "includes", includeFileName)
			_, err := os.Stat(includeFilePath)
			if os.IsNotExist(err) {
				log.Fatal(err)
			}
			includeFileContent := readFile(includeFilePath)
			// Adds contents of file to the constructed string instead of token
			compiledString = append(compiledString, includeFileContent)
			continue
		}
		compiledString = append(compiledString, token.String())
	}
	// Concatenate the built string
	joinedString := strings.Join(compiledString, "")
	// Now, create a new file in the build directory
	buildPath := filepath.Join(target, filepath.Base(file))
	buildFile, err := os.Create(buildPath)
	checkError(err)
	buildFile.WriteString(joinedString)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	return string(data)
}
*/
