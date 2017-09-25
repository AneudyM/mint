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

var srcDir string = "src/"

func cmdBuildProject(cmd *cobra.Command, args []string) {
	// Read "src" directory
	err := filepath.Walk(srcDir, printFile)
	checkError(err)
}

func init() {
	RootCmd.AddCommand(cmdBuild)
}

func printFile(path string, file os.FileInfo, err error) error {

	/*
		buildDir := "build"
		if file.IsDir() {
			if file.Name() != "src"{
				dirName := filepath.Join(buildDir, file.Name())
				os.Mkdir(dirName, 0666)
			}
		}
	*/
	return nil
}

func compileHtmlFile(filename string) {
	var newString []string
	sourcePath := filename
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
				log.Fatal("Could not open include file. " + includeFile)
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
