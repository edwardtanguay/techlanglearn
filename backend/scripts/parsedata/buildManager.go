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
				fmt.Printf("[%v]\n", line)
				println(len(line))
				line = "kinde; en; 2024; 4.9; yt; https://www.youtube.com/watch?v=_EjOHdRihjA"
				fmt.Printf("[%v]\n", line)
				println(len(line))
				// buildTutorial(line)
			}
		}
	}
	return nil
}

// process e.g.: kinde; en; 2024; 4.9; yt; https://www.youtube.com/watch?v=_EjOHdRihjA
func buildTutorial(line string) error {
	// fmt.Printf("BUILD\n");
	// fmt.Printf("BUILD: |%s|", line)
	return nil
}