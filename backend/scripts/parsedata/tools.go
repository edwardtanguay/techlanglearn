package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Get all files of a certain kind as a slice of strings

mdPathAndFileNames, _ := getFilesFromDirectory("../../static/data", "md")

- use relative path

- extension is mandatory
*/
func getFilesFromDirectory(dirPath, ext string) ([]string, error) {
	var fileList []string

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

/*
Get all lines from a file as a slice of strings

lines := getLinesFromFile("../../notes.txt")

- use relative path
*/
func getLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return lines
}

/*
Get content from a file as a string

lines := getContentFromFile("../../notes.txt")

- use relative path
*/
func getContentFromFile(fileName string) string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	return contents
}

/*
Output information in console in a uniform way

devlog("no files are locked")

devlog(fmt.Sprintf("There are %d flashcards.", len(flashcards)))
*/
func devlog(line string) {
	fmt.Printf(">>> %s\n", line)
}

/*
Write text to a file

writeTextFile("../../src/data/test.txt", "testcontent")

- use relative path
*/
func writeTextFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("File could not be created: %s\n", err.Error())
	}
	_, err2 := file.WriteString(content)
	if err2 != nil {
		fmt.Printf("Could not write to file: %s\n", err2.Error())
	}
}

/*
Get a date in form of 2024-11-01 from a line, otherwise empty string

calendarDate := getDateFromLine(rawLine)
*/
func getDateFromLine(line string) string {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	calendarDate := re.FindString(line)
	return calendarDate
}

func getDurationFromLine(line string) string {
	re := regexp.MustCompile(`^\d{2}:\d{2}:\d{2}$`)
	if re.MatchString(line) {
		return line
	}

	reTime := regexp.MustCompile(`(\d{1,2}):(\d{2})`)
	match := reTime.FindStringSubmatch(line)
	if match == nil {
		return ""
	}

	minutes, _ := strconv.Atoi(match[1])
	seconds, _ := strconv.Atoi(match[2])

	return fmt.Sprintf("00:%02d:%02d", minutes, seconds)
}

func durationToSeconds(duration string) int {
	parts := strings.Split(duration, ":")
	if len(parts) != 3 {
		return 0
	}
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	seconds, _ := strconv.Atoi(parts[2])

	return hours*3600 + minutes*60 + seconds
}

func secondsToDuration(seconds int) string {
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func computeDurationPerDay(timeUnits []TimeUnit) []TimeUnit {
	durationMap := make(map[string]int)

	for _, timeUnit := range timeUnits {
		seconds := durationToSeconds(timeUnit.Duration)
		durationMap[timeUnit.CalendarDate] += seconds
	}

	var totalTimeUnits []TimeUnit

	for date, totalSeconds := range durationMap {
		totalTimeUnits = append(totalTimeUnits, TimeUnit{
			CalendarDate: date,
			Duration:     secondsToDuration(totalSeconds),
		})
	}

	sort.Slice(totalTimeUnits, func(i, j int) bool {
		date1, _ := time.Parse("2006-01-02", totalTimeUnits[i].CalendarDate)
		date2, _ := time.Parse("2006-01-02", totalTimeUnits[j].CalendarDate)
		return date1.Before(date2)
	})

	return totalTimeUnits
}

// Helper function to parse duration from "hh:mm:ss" format
func parseDuration(durationStr string) (time.Duration, error) {
	return time.ParseDuration(fmt.Sprintf("%sh%sm%ss", durationStr[0:2], durationStr[3:5], durationStr[6:8]))
}

// Helper function to format a duration to "hh:mm:ss" format
func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func convertToSnakeCase(line string) string {
	re := regexp.MustCompile(`[^\w\s]`)
	cleanLine := re.ReplaceAllString(line, "")
	lowercaseLine := strings.ToLower(cleanLine)
	words := strings.Fields(lowercaseLine)
	return strings.Join(words, "-")
}

func getTutorialFileName(_pathAndFileName string) string {
	pathAndFileName := filepath.ToSlash(_pathAndFileName) // changes \ to /
	lastSlashIndex := strings.LastIndex(pathAndFileName, "/")
	if lastSlashIndex != -1 {
		return pathAndFileName[lastSlashIndex+1:]
	}
	return pathAndFileName
}

func getRestOfLine(text, marker string) string {
	index := strings.Index(text, marker)
	if index == -1 {
		return ""
	}
	return strings.TrimSpace(text[index+len(marker):])
}

func getFieldValueFromLines(lines []string, fieldIdCode string) string {
	for i, line := range lines {
		marker := fieldIdCode + ":"
		if strings.Contains(line, marker) {
			return getRestOfLine(line, marker)
		}
		if i > 20 {
			return ""
		}
	}
	return ""
}

func getPlatformFromUrl(url string) string {
	if strings.Contains(url, "youtube.") {
		return "youtube"
	}
	if strings.Contains(url, "linkedin.") {
		return "linkedInLearning"
	}
	return ""
}

func softIncludes(main string, includesText string) bool {
	mainLower := strings.ToLower(main)
	includesTextLower := strings.ToLower(includesText)
	return strings.Contains(mainLower, includesTextLower)
}

func getLinesFromMarkerToEnd(lines []string, marker string) []string {
	for i, line := range lines {
		if strings.Contains(line, marker) {
			return lines[i:]
		}
	}
	return []string{}
}
