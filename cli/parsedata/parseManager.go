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

	// var flashcards []Flashcard

	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "md")
	for _, mdPathAndFileName := range mdPathAndFileNames {

		fileName := filepath.Base(mdPathAndFileName)
		fmt.Printf("FILE: %s\n", fileName)
		fmt.Println("=====================================================")

		lines := getLinesFromFile(mdPathAndFileName)
		flashcards, _ := getFlashcards(lines)

		fmt.Println("=====================================================")
		fmt.Printf("Number of lines: %v\n", len(lines))
		fmt.Printf("Number of flashcards: %v\n", len(flashcards))
		fmt.Printf("=====================================================\n\n")

	}
}

func getFlashcards(lines []string) ([]Flashcard, error) {
	foundFirstBackticks := false
	foundVocab := false
	flashcards := []Flashcard{}
	flashcardLines := []string{}
	for _, line := range lines {
		if foundVocab && foundFirstBackticks {
			flashcardLines = append(flashcardLines, line)
			// fmt.Printf("RECORDING FLASHCARD, there are %v flashcard lines now\n", len(flashcardLines))
		}
		if strings.Contains(line, "## VOCAB") {
			foundVocab = true
		}
		if strings.Contains(line, "```") {
			foundFirstBackticks = true
		}
	}
	fmt.Printf("FINISHED: there are %v flashcard lines now\n", len(flashcardLines))

	for _, flashcardLine := range flashcardLines {
		fmt.Printf("LINE: %s\n", flashcardLine)
	}

	// // create data
	// processingLineType := "front"
	// front := ""
	// back := ""
	// // fileName := filepath.Base(mdPathAndFileName)
	// for _, flashcardLine := range flashcardLines {
	// 	switch processingLineType {
	// 	case "front":
	// 		front = flashcardLine
	// 		processingLineType = "back"
	// 	case "back":
	// 		back = flashcardLine
	// 		processingLineType = "BLANK"
	// 	case "BLANK":
	// 		flashcards = append(flashcards, Flashcard{front, back})
	// 		processingLineType = "front"
	// 	}

	// }
	return flashcards, nil
}
