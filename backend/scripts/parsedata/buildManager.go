package main

import "fmt"

func build(directory string) {
	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "tts.txt")
	buildTutorials(mdPathAndFileNames)
}

func buildTutorials(mdPathAndFileNames []string) error {
	for _, mdPathAndFileName := range mdPathAndFileNames {
		lines := getLinesFromFile(mdPathAndFileName)
		for _, line := range lines {
			fmt.Printf("LINE: %s\n", line)
		}
	}
	return nil
}