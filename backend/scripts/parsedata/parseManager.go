package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

/*
Parses the md files into JSON for the website to consume

parse("../../static/data")

- use relative path
*/
func parse(directory string) {
	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "md")
	parseFlashcards(mdPathAndFileNames)
	parseStats(mdPathAndFileNames)
}

func parseStats(mdPathAndFileNames []string) error {
	config, _ := LoadConfig()
	timeUnits := []TimeUnit{}
	for _, mdPathAndFileName := range mdPathAndFileNames {
		lines := getLinesFromFile(mdPathAndFileName)

		// LinkedInLearning files
		linkedInLearningFileTimeUnits, _ := getLinkedInLearningTimeUnitsFromFile(lines)
		timeUnits = append(timeUnits, linkedInLearningFileTimeUnits...)

		// other files (YouTube videos, other videos, articles)
		otherFileTimeUnits, _ := getOtherTimeUnitsFromFile(lines)
		timeUnits = append(timeUnits, otherFileTimeUnits...)
	}

	totalTimeUnits := computeDurationPerDay(timeUnits)
	filledTimeUnits := fillInZeroDays(totalTimeUnits)
	statsInfo := calculateStatsInfo(filledTimeUnits)
	reversedstatsInfo := reverseOrderOfTimeUnits(statsInfo)

	jsonData, _ := json.MarshalIndent(reversedstatsInfo, "", "\t")

	writeTextFile("../../../"+config.WebDataDirectory+"/stats.json", string(jsonData))
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

func getLinkedInLearningTimeUnitsFromFile(lines []string) ([]TimeUnit, error) {
	timeUnits := []TimeUnit{}
	for _, rawLine := range lines {
		if strings.HasPrefix(rawLine, "##") {
			calendarDate := getDateFromLine(rawLine)
			duration := getDurationFromLine(rawLine)
			if calendarDate != "" {
				timeUnits = append(timeUnits, TimeUnit{
					CalendarDate: calendarDate,
					Duration:     duration,
				})
			}
		}
	}
	return timeUnits, nil
}

func getOtherTimeUnitsFromFile(lines []string) ([]TimeUnit, error) {
	timeUnits := []TimeUnit{}
	recordingLines := false
	for _, rawLine := range lines {
		if recordingLines && strings.HasPrefix(rawLine, "##") {
			recordingLines = false
		}

		if recordingLines {
			calendarDate := getDateFromLine(rawLine)
			duration := getDurationFromLine(rawLine)
			if calendarDate != "" {
				timeUnits = append(timeUnits, TimeUnit{
					CalendarDate: calendarDate,
					Duration:     duration,
				})
			}
		}

		if strings.HasPrefix(rawLine, "## watchlog") {
			recordingLines = true
		}
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

func fillInZeroDays(timeUnits []TimeUnit) []TimeUnit {

	layout := "2006-01-02"
	sort.Slice(timeUnits, func(i, j int) bool {
		date1, _ := time.Parse(layout, timeUnits[i].CalendarDate)
		date2, _ := time.Parse(layout, timeUnits[j].CalendarDate)
		return date1.Before(date2)
	})

	startDate, _ := time.Parse(layout, timeUnits[0].CalendarDate)
	endDate, _ := time.Parse(layout, timeUnits[len(timeUnits)-1].CalendarDate)

	dateMap := make(map[string]TimeUnit)
	for _, unit := range timeUnits {
		dateMap[unit.CalendarDate] = unit
	}

	var filledTimeUnits []TimeUnit
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		dateStr := date.Format(layout)
		if unit, exists := dateMap[dateStr]; exists {
			filledTimeUnits = append(filledTimeUnits, unit)
		} else {
			filledTimeUnits = append(filledTimeUnits, TimeUnit{
				CalendarDate: dateStr,
				Duration:     "00:00:00",
			})
		}
	}

	return filledTimeUnits
}

func calculateStatsInfo(timeUnits []TimeUnit) TimesInfo {

	dateSet := make(map[string]bool)
	var totalDuration time.Duration

	for _, unit := range timeUnits {
		duration, err := parseDuration(unit.Duration)
		if err != nil {
			fmt.Printf("Error parsing duration for date %s: %v\n", unit.CalendarDate, err)
			continue
		}
		totalDuration += duration
		dateSet[unit.CalendarDate] = true
	}

	totalDays := len(dateSet)

	var averageDuration time.Duration
	if totalDays > 0 {
		averageDuration = totalDuration / time.Duration(totalDays)
	}

	return TimesInfo{
		TotalDays:             totalDays,
		AverageDurationPerDay: formatDuration(averageDuration),
		TotalDuration:         formatDuration(totalDuration),
		TimeUnits:             timeUnits,
	}
}

func reverseOrderOfTimeUnits(timesInfo TimesInfo) TimesInfo {
	reversedTimeUnits := make([]TimeUnit, len(timesInfo.TimeUnits))

	for i, unit := range timesInfo.TimeUnits {
		reversedTimeUnits[len(timesInfo.TimeUnits)-1-i] = unit
	}

	return TimesInfo{
		TotalDays:             timesInfo.TotalDays,
		AverageDurationPerDay: timesInfo.AverageDurationPerDay,
		TotalDuration:         timesInfo.TotalDuration,
		TimeUnits:             reversedTimeUnits,
	}
}