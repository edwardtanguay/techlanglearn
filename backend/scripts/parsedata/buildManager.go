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

// process e.g.: kinde; en; 00:09:37; 2024; 4.9; https://www.youtube.com/watch?v=_EjOHdRihjA; quick video showing how to build Kinde into a React site
func buildTutorial(line string) error {
	tutorial := parseTutorialLine(line)
	fmt.Printf("the title is [%s]\n", tutorial.Title)
	return nil
}

func parseTutorialLine(line string) Tutorial {
	parts := strings.Split(line, ";")

	if len(parts) < 8 {
		panic("Invalid input: not enough fields in the line")
	}

	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return Tutorial{
		Topics:      strings.TrimSpace(parts[0]),
		Language:    strings.TrimSpace(parts[1]),
		Duration:    strings.TrimSpace(parts[2]),
		Year:        strings.TrimSpace(parts[3]),
		Rank:        strings.TrimSpace(parts[4]),
		Url:         strings.TrimSpace(parts[5]),
		Title:       strings.TrimSpace(parts[6]),
		Description: strings.TrimSpace(parts[7]),
	}
}
