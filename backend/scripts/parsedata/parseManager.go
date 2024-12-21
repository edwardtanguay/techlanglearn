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

parseCurrentTutorialFiles("../../static/data")

- use relative path
*/
func parseCurrentTutorialFiles(directory string) {
	mdPathAndFileNames, _ := getFilesFromDirectory(directory, "md")
	parseFlashcards(mdPathAndFileNames)
	parseStats(mdPathAndFileNames)
	createTutorialJsonFile(mdPathAndFileNames)
}

func createTutorialJsonFile(mdPathAndFileNames []string) error {
	config, _ := LoadConfig()
	tutorials := []Tutorial{}
	for _, mdPathAndFileName := range mdPathAndFileNames {
		lines := getLinesFromFile(mdPathAndFileName)
		fileName := getTutorialFileName(mdPathAndFileName)

		tutorial := Tutorial{}
		tutorial.Title = getRestOfLine(lines[0], "#")
		tutorial.FileIdCode = strings.TrimSuffix(fileName, ".md")
		tutorial.Url = lines[2]
		tutorial.Description = getFieldValueFromLines(lines, "description")
		tutorial.Rank = getFieldValueFromLines(lines, "rank")
		tutorial.Topics = getFieldValueFromLines(lines, "topics")
		tutorial.Language = getFieldValueFromLines(lines, "language")
		tutorial.Duration = getFieldValueFromLines(lines, "duration")
		tutorial.Year = getFieldValueFromLines(lines, "year")
		tutorial.Status = getFieldValueFromLines(lines, "status")
		tutorial.Platform = getPlatformFromUrl(tutorial.Url)

		tutorials = append(tutorials, tutorial)
		jsonData, _ := json.MarshalIndent(tutorials, "", "\t")
		writeTextFile("../../../"+config.WebDataDirectory+"/tutorials.json", string(jsonData))
	}
	return nil
}

func parseStats(mdPathAndFileNames []string) error {
	config, _ := LoadConfig()
	timeUnits := []TimeUnit{}
	for _, mdPathAndFileName := range mdPathAndFileNames {
		lines := getLinesFromFile(mdPathAndFileName)
		addLinkedInLearningTimeUnitsFromFile(lines, &timeUnits)
		getOtherTimeUnitsFromFile(lines, &timeUnits)
	}

	totalTimeUnits := computeDurationPerDay(timeUnits)
	filledTimeUnits := fillInZeroDays(totalTimeUnits)
	statsInfo := calculateStatsInfo(filledTimeUnits)
	reversedstatsInfo := reverseOrderOfTimeUnits(statsInfo)
	withAddedDayStatsInfo := addCurrentDayIfDoesntExist(reversedstatsInfo)

	jsonData, _ := json.MarshalIndent(withAddedDayStatsInfo, "", "\t")

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
	writeTextFile("C:\\edward\\learn2024\\trilingual-little-prince\\src\\data\\flashcards.json", string(jsonData))
	return nil
}

func addLinkedInLearningTimeUnitsFromFile(lines []string, addTimeUnits *[]TimeUnit) error {
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
	*addTimeUnits = append(*addTimeUnits, timeUnits...)
	return nil
}

func getOtherTimeUnitsFromFile(lines []string, allTimeUnits *[]TimeUnit) error {
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
	intervalUnits, _ := convertToIntervalTimes(timeUnits)
	*allTimeUnits = append(*allTimeUnits, intervalUnits...)
	return nil
}

func convertToIntervalTimes(timeUnits []TimeUnit) ([]TimeUnit, error) {
	intervalUnits := make([]TimeUnit, len(timeUnits))

	if len(timeUnits) > 0 {
		intervalUnits[0] = timeUnits[0]
	}

	for i := 1; i < len(timeUnits); i++ {
		currentDuration, _ := parseDuration(timeUnits[i].Duration)
		prevDuration, _ := parseDuration(timeUnits[i-1].Duration)
		diff := currentDuration - prevDuration
		intervalUnits[i] = TimeUnit{
			CalendarDate: timeUnits[i].CalendarDate,
			Duration:     formatDuration(diff),
		}
	}
	return intervalUnits, nil
}

func getFlashcardsFromFile(lines []string) ([]Flashcard, error) {

	config, _ := LoadConfig()

	flashcards := []Flashcard{}
	language := ""
	marker := "## VOCAB"

	// get all lines in the general vocab block (currently at end of file)
	vocabBlockLines := getLinesFromMarkerToEnd(lines, marker)

	// define the language from the first line
	if len(vocabBlockLines) == 0 {
		return flashcards, nil
	}
	restOfLine := getRestOfLine(vocabBlockLines[0], marker)
	if softIncludes(restOfLine, "spanish") {
		language = "es"
	}
	if softIncludes(restOfLine, "italian") {
		language = "it"
	}

	if softIncludes(restOfLine, "french") {
		language = "fr"
	}

	if softIncludes(restOfLine, "dutch") {
		language = "nl"
	}

	// define vocabLines (only the text of the flashcards themselves)
	vocabLines := vocabBlockLines[1:] // remove first line (language heading)
	vocabLines = removeAllLinesWithMarker(vocabLines, "```")

	lineBlocks := getLineBlocksFromLines(vocabLines)

	for _, lineBlock := range lineBlocks {
		lineBlock = padLineBlock(lineBlock, 4)
		flashcard := Flashcard{
			Language:    language,
			Front:       lineBlock[0],
			Back:        lineBlock[1],
			WhenCreated: "",
			Extras:      "",
		}

		// whenCreated
		if lineBlock[2] != "" {
			flashcard.WhenCreated = lineBlock[2]
		} else {
			flashcard.WhenCreated = config.DefaultWhenCreated
		}

		// transfer extras from back to extras
		if strings.Contains(flashcard.Back, ";") {
			flashcard.Extras = getRestOfLine(flashcard.Back, ";")
			flashcard.Back = deleteRestAtMarker(flashcard.Back, ";")
		}

		flashcards = append(flashcards, flashcard)
	}

	return flashcards, nil
}

func fillInZeroDays(timeUnits []TimeUnit) []TimeUnit {

	var filledTimeUnits []TimeUnit
	if len(timeUnits) > 0 {

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

func addCurrentDayIfDoesntExist(timesInfo TimesInfo) TimesInfo {
	now := time.Now()
	itsLateEnough := now.Hour() >= 5

	if !itsLateEnough {
		return timesInfo
	}

	dateSet := make(map[string]bool)
	for _, timeUnit := range timesInfo.TimeUnits {
		dateSet[timeUnit.CalendarDate] = true
	}

	var earliestDate time.Time
	if len(timesInfo.TimeUnits) > 0 {
		var err error
		earliestDate, err = time.Parse("2006-01-02", timesInfo.TimeUnits[len(timesInfo.TimeUnits)-1].CalendarDate)
		if err != nil {
			panic("Invalid date format in TimeUnits")
		}
	} else {
		earliestDate = now
	}

	missingDays := []TimeUnit{}
	for date := earliestDate; !date.After(now); date = date.AddDate(0, 0, 1) {
		dateStr := date.Format("2006-01-02")
		if !dateSet[dateStr] {
			missingDays = append(missingDays, TimeUnit{
				CalendarDate: dateStr,
				Duration:     "00:00:00",
			})
		}
	}

	timesInfo.TimeUnits = append(missingDays, timesInfo.TimeUnits...)

	sort.Slice(timesInfo.TimeUnits, func(i, j int) bool {
		return timesInfo.TimeUnits[i].CalendarDate > timesInfo.TimeUnits[j].CalendarDate
	})

	timesInfo = calculateStatsInfo(timesInfo.TimeUnits)

	return timesInfo
}
