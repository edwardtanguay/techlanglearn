package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "parsedata",
		Short: "Reads the .md files and parses them into JSON files that contain data from the md files, plus extended HTML to be displayed on the site.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("This script will parse these files:\n\n")

			mdPathAndFileNames, _ := getFilesFromDirectory("../../static/data", "md")
			for _, mdPathAndFileName := range mdPathAndFileNames {
				lines := getLinesFromFile(mdPathAndFileName)
				fileName := filepath.Base(mdPathAndFileName)
				fmt.Printf("%s - %d lines\n", fileName, len(lines))
				fmt.Print("=============================================================\n")
				for _, line := range lines {
					fmt.Printf("%s\n", line)
				}
				fmt.Print("=============================================================\n\n")
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
