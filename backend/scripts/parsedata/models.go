package main

type Config struct {
	TutorialFilesDirectory string `json:"tutorialFilesDirectory"`
	WebDataDirectory       string `json:"webDataDirectory"`
}

type Flashcard struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

type TimeUnit struct {
	CalendarDate string `json:"calendarDate"`
	Duration  string `json:"duration"`
}

type TimesInfo struct {
	TotalDays int
	AverageDurationPerDay string
	TotalDuration string
	TimeUnits []TimeUnit
}