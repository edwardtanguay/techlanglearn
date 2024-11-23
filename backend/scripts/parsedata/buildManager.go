package main

import (
	"fmt"
	"strings"
)

func build(directory string) {
	mdPathAndFileNames, _ := getFilesFromDirectory(directory+"/general", "tts.txt")
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
		markCreatedLinesAsFinished(mdPathAndFileName)
	}
	return nil
}

// kinde; en; 00:09:37; 2024; 4.9; https://www.youtube.com/watch?v=_EjOHdRihjA; Add Auth & Protect Routes in React in 3 Minutes (Kinde); quick video showing how to build Kinde into a React site
func buildTutorial(line string) {
	tutorial := parseTutorialLine(line)
	createTutorialFile(tutorial)
}

func createTutorialFile(tutorial Tutorial) {
	tutorialPathAndFileName := "../../../static/data/" + tutorial.FileIdCode + ".md"
	if tutorial.Platform == "youtube" {
		devlog(fmt.Sprintf("Creating Youtube file for \"%s\"", tutorial.Title))
		createFileWithTemplateAndData(tutorialPathAndFileName, "../../../static/data/templates/template-youtube.tuttmpl.txt", tutorial)
	} else {
		devlog(fmt.Sprintf("Sorry, %s is not yet a supported platform.", tutorial.Platform))
	}
}

func createFileWithTemplateAndData(targetPathAndFileName string, templatePathAndFileName string, tutorial Tutorial) {
	content := getContentFromFile(templatePathAndFileName)
	content = strings.ReplaceAll(content, "@@title", tutorial.Title)
	content = strings.ReplaceAll(content, "@@url", tutorial.Url)
	content = strings.ReplaceAll(content, "@@duration", tutorial.Duration)
	content = strings.ReplaceAll(content, "@@language", tutorial.Language)
	content = strings.ReplaceAll(content, "@@topics", tutorial.Topics)
	content = strings.ReplaceAll(content, "@@rank", tutorial.Rank)
	content = strings.ReplaceAll(content, "@@description", tutorial.Description)
	writeTextFile(targetPathAndFileName, content)
}

func markCreatedLinesAsFinished(coursesPathAndFileName string) {
	lines := getLinesFromFile(coursesPathAndFileName)
	for i, line := range lines {
		lines[i] = line + "-ADDED"
	}
	content := strings.Join(lines, "\n")
	writeTextFile(coursesPathAndFileName, content)
}

func parseTutorialLine(line string) Tutorial {
	parts := strings.Split(line, ";")

	if len(parts) < 8 {
		panic("Invalid input: not enough fields in the line")
	}

	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	tutorial := Tutorial{
		Topics:      strings.TrimSpace(parts[0]),
		Language:    strings.TrimSpace(parts[1]),
		Duration:    strings.TrimSpace(parts[2]),
		Year:        strings.TrimSpace(parts[3]),
		Rank:        strings.TrimSpace(parts[4]),
		Url:         strings.TrimSpace(parts[5]),
		Title:       strings.TrimSpace(parts[6]),
		Description: strings.TrimSpace(parts[7]),
		Platform:    "",
		FileIdCode:  "",
	}

	// add platform
	if strings.Contains(tutorial.Url, "youtube.") {
		tutorial.Platform = "youtube"
	}
	if strings.Contains(tutorial.Url, "linkedin.") {
		tutorial.Platform = "linkedInLearning"
	}

	// add fileIdCode
	tutorial.FileIdCode = buildFileIdCode(tutorial)

	return tutorial
}

func buildFileIdCode(tutorial Tutorial) string {
	platformAbbreviation := ""
	switch tutorial.Platform {
	case "youtube":
		platformAbbreviation = "yt"
	case "linkedInLearning":
		platformAbbreviation = "ll"
	}

	snakeCase := convertToSnakeCase(tutorial.Title)

	return fmt.Sprintf("%s-%s-%s", tutorial.Language, platformAbbreviation, snakeCase)
}
