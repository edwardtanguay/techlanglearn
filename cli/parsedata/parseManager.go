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

		fileName := filepath.Base(mdPathAndFileName)
		fmt.Printf("FILE: %s\n", fileName)

		lines := getLinesFromFile(mdPathAndFileName)
		flashcards, _ := getFlashcards(lines)

		fmt.Println("=====================================================")
		fmt.Printf("Number of flashcards: %v\n", len(flashcards))
		fmt.Println("=====================================================")
		// for _, flashcard := range flashcards {
		// 	fmt.Printf("FRONT: %v", flashcard.Front)
		// }
		fmt.Println("=====================================================")

	}
}

func getFlashcards(lines []string) ([]Flashcard, error) {

	foundFirstBackticks := false
	foundVocab := false
	flashcards := []Flashcard{}
	flashcardLines := []string{}

	// get slice of all lines that have to do with flashcards
	for _, line := range lines {
		if foundVocab && foundFirstBackticks {
			flashcardLines = append(flashcardLines, line)
		}
		if strings.Contains(line, "## VOCAB") {
			foundVocab = true
		}
		if strings.Contains(line, "```") {
			foundFirstBackticks = true
		}
	}

	// create slice of flashcard instances
	processingLineType := "front"
	front := ""
	back := ""
	fmt.Printf("going through %d flashcards \n", len(flashcardLines))
	for _, flashcardLine := range flashcardLines {
		switch processingLineType {
		case "front":
			front = flashcardLine
			processingLineType = "back"
		case "back":
			back = flashcardLine
			processingLineType = "BLANK"
		case "BLANK":
			flashcards = append(flashcards, Flashcard{front, back})
			fmt.Printf("FRONTINSIDE: %s\n", front)
			processingLineType = "front"
		}
	}

	return flashcards, nil
}
