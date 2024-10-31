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
			fmt.Printf("TODO: parse these files:\n\n")
			mdPathAndFileNames, _ := getMdFiles("../../static/data")
			for _, mdPathAndFileName := range mdPathAndFileNames {
				fileName := filepath.Base(mdPathAndFileName)
				fmt.Printf("%s\n", fileName)
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
