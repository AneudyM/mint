package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"path/filepath"
)

var cmdBuild = &cobra.Command{
	Use: "build",
	Short: "Build the project.",
	Long: `Build the complete project and places the files in the project's /build directory`,
	Run: cmdBuildProject,
}


func cmdBuildProject(cmd *cobra.Command, args []string){
	// Compile files in "src" directory
	files, err := ioutil.ReadDir("src/")
	checkError(err)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		sourceFile := file.Name()
		compileHtmlFile(sourceFile)
	}
}

func init() {
	RootCmd.AddCommand(cmdBuild)
}

func compileHtmlFile(filename string) {
	var newString []string
	sourcePath := filepath.Join("src/", filename)
	targetPath := filepath.Join("build/", filename)
	// Tokenize input source file
	fileData := readFile(sourcePath)
	r := strings.NewReader(fileData)
	// Create a new tokenizer to tokenize the file data
	tokenizer := html.NewTokenizer(r)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if token.Data == "include" {
			if len(token.Attr) > 1 || len(token.Attr) < 1 {
				log.Fatal("Wrong number of attributes")
			}
			if token.Attr[0].Key != "src" {
				log.Fatal("Invalid attribute name")
			}
			if token.Attr[0].Val == "" {
				log.Fatal("Empty attribute 'src'")
			}
			// Validate include file
			includeFile := token.Attr[0].Val
			if checkFile(includeFile) == false {
				log.Fatal("Could not open include file.")
				return
			}
			includeFileData := readFile(includeFile)
			newString = append(newString, includeFileData)
			continue
		}
		newString = append(newString, token.String())
	}
	finalString := strings.Join(newString, "")
	newFile, err := os.Create(targetPath)
	checkError(err)
	newFile.WriteString(finalString)
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

func checkFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}