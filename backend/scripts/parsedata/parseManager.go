package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
Parses the md files into JSON for the website to consume

parse("../../static/data")

- use relative path
*/
func parse(directory string) {
	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "md")
	parseFlashcards(mdPathAndFileNames)
	parseTimes(mdPathAndFileNames)
}

func parseTimes(mdPathAndFileNames []string) error {
	config, _ := LoadConfig()
	timeUnits := []TimeUnit{}
	for _, mdPathAndFileName := range mdPathAndFileNames {
		lines := getLinesFromFile(mdPathAndFileName)
		fileTimeUnits, _ := getTimeUnitsFromFile(lines)
		timeUnits = append(timeUnits, fileTimeUnits...)
	}
	devlog(fmt.Sprintf("%d time units were created.", len(timeUnits)))

	jsonData, _ := json.MarshalIndent(timeUnits, "", "\t")
	writeTextFile("../../../"+config.WebDataDirectory+"/times.json", string(jsonData))
	return nil
}

func parseFlashcards(mdPathAndFileNames []string) error {
	config, _ := LoadConfig()
	flashcards := []Flashcard{}
	for _, mdPathAndFileName := range mdPathAndFileNames {
		lines := getLinesFromFile(mdPathAndFileName)
		fileFlashcards, _ := getFlashcardsFromFile(lines)
		flashcards = append(flashcards, fileFlashcards...)
	}
	devlog(fmt.Sprintf("%d flashcards were created.", len(flashcards)))

	jsonData, err := json.MarshalIndent(flashcards, "", "	")
	if err != nil {
		println("could not convert to JSON text")
	}
	writeTextFile("../../../"+config.WebDataDirectory+"/flashcards.json", string(jsonData))
	return nil
}

func getTimeUnitsFromFile(lines []string) ([]TimeUnit, error) {
	for _, rawLine := range lines {
		if strings.HasPrefix(rawLine, "##") {
			calendarDate := getDateFromLine(rawLine)
			if calendarDate != "" {
				devlog(rawLine)
			}
		}
	}
	timeUnits := []TimeUnit{
		{
			CalendarDate: "ccc",
			Duration:     "ddd",
		},
	}
	return timeUnits, nil
}

func getFlashcardsFromFile(lines []string) ([]Flashcard, error) {

	foundFirstBackticks := false
	foundVocab := false
	flashcards := []Flashcard{}
	flashcardLines := []string{}

	// get slice of all lines that have to do with flashcards
	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if foundVocab && foundFirstBackticks {
			if (len(flashcardLines) == 0 && line == "") || line == "```" {
				// don't save
			} else {
				flashcardLines = append(flashcardLines, line)
			}
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
			processingLineType = "front"
		}
	}

	// if there were no cards but a marker was saved, then fix it
	if len(flashcards) == 1 && flashcards[0].Back == "```" {
		println(111, "delete this")
	}

	return flashcards, nil
}
