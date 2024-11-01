package main

import (
	"fmt"
	"path/filepath"
)

/*
Parses the md files into JSON for the website to consume

parse("../../static/data")

- use relative path
*/
func parse(directory string) {
	fmt.Printf("This script will parse these files:\n\n")

	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "md")
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
}
