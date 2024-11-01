package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "parsedata",
		Short: "Reads the .md files and parses them into JSON files that contain data from the md files, plus extended HTML to be displayed on the site.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("This script will parse these files:\n\n")
			mdPathAndFileNames, _ := getMdFiles("../../static/data")
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

func getMdFiles(dirPath string) ([]string, error) {
	var fileList []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func getLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	return lines
}
