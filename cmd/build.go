package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var cmdBuild = &cobra.Command{
	Use:   "build",
	Short: "Build the project.",
	Long:  `Build the complete project and places the files in the project's /build directory`,
	Run:   cmdBuildProject,
}

var srcDir string = "src"

func cmdBuildProject(cmd *cobra.Command, args []string) {
	//(1) Make sure the user runs the command from its project's root directory
	_, err := os.Stat(srcDir)
	if os.IsNotExist(err) {
		log.Fatal("Make sure you are in the root of your project.")
	}
	//(2) If user is in the root of the project, then run htmlCompiler
	filepath.Walk(srcDir, build)
}

func init() {
	RootCmd.AddCommand(cmdBuild)
}

func build(path string, fileInfo os.FileInfo, err error) error {
	//(1) Get the current directory's absolute path
	cwd, _ := os.Getwd()
	absPath := filepath.Join(cwd, path)
	buildPath := filepath.Join(cwd, "/build")
	//(2) Process files according to their extension
	//(2.1) Get file extension if the file is not a directory
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
				log.Fatal("Too many attributes.")
			}
			if token.Attr[0].Key != srcDir {
				log.Fatal("Unrecognized", "'"+token.Attr[0].Key+"'", "attribute.")
			}
			if token.Attr[0].Val == "" {
				log.Fatal("Empty attribute 'src'")
			}
			includeFileName := token.Attr[0].Val
			// Retrive include file
			// But first make sure it exists:
			includeFilePath := filepath.Join(cwd, srcDir, "includes", includeFileName)
			_, err := os.Stat(includeFilePath)
			if os.IsNotExist(err){
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