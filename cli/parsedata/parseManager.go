package main

import (
	"fmt"
	"path/filepath"
	"strings"
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

		flashcards, _ := getFlashcards(lines)

		// display data
		fileName := filepath.Base(mdPathAndFileName)
		fmt.Printf("%s - %d lines\n", fileName, len(lines))
		fmt.Print("=============================================================\n")
		for _, flashcard := range flashcards {
			fmt.Printf("FLASHCARD: %s\n", flashcard)
		}
		fmt.Print("=============================================================\n\n")
	}
}

func getFlashcards(lines []string) ([]string, error) {
	flashcards := []string{}
	for _, line := range lines {
		if strings.Contains(line, "the") {
			flashcards = append(flashcards, line)
		}
	}
	return flashcards, nil
}
