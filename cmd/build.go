package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var cmdBuild = &cobra.Command{
	Use: "build",
	Short: "Build the project.",
	Long: `Build the complete project and places the files in the project's /build directory`,
	Run: cmdBuildProject,
}


func cmdBuildProject(cmd *cobra.Command, args []string){

}

func init() {
	RootCmd.AddCommand(cmdBuild)
}

func CopyDir(origin string, dest string){
	if origin == "" {
		log.Fatal()
	}
}