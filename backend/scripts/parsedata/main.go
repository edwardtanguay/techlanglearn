package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	config, _ := LoadConfig()
	var rootCmd = &cobra.Command{
		Use:   "parsedata",
		Short: "Reads the .md files and parses them into JSON files that contain data from the md files, plus extended HTML to be displayed on the site.",
		Run: func(cmd *cobra.Command, args []string) {
			build("../../../" + config.TutorialFilesDirectory)
			parse("../../../" + config.TutorialFilesDirectory)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
