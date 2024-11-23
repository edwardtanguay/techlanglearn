package main

import (
	"fmt"
	"strings"
)

func build(directory string) {
	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "tts.txt")
	buildTutorials(mdPathAndFileNames)
}

func buildTutorials(mdPathAndFileNames []string) error {
	for _, mdPathAndFileName := range mdPathAndFileNames {
		rawLines := getLinesFromFile(mdPathAndFileName)
		for _, rawLine := range rawLines {
			if strings.HasPrefix(rawLine, ">>create>>") {
				line := strings.TrimPrefix(rawLine, ">>create>>")
				line = strings.TrimSpace(line)
				buildTutorial(line)
			}
		}
	}
	return nil
}

// process e.g.: kinde; en; 2024; 4.9; https://www.youtube.com/watch?v=_EjOHdRihjA
func buildTutorial(line string) error {
	fmt.Printf("BUILD: [%s]\n", line)
	return nil
}