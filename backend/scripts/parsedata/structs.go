package main

type Config struct {
	TutorialFilesDirectory string `json:"tutorialFilesDirectory"`
	WebDataDirectory       string `json:"webDataDirectory"`
}

type Flashcard struct {
	Language    string `json:"language"`
	Front       string `json:"front"`
	Back        string `json:"back"`
	WhenCreated string `json:"whenCreated"`
	Extras      string `json:"extras"`
}

type TimeUnit struct {
	CalendarDate string `json:"calendarDate"`
	Duration     string `json:"duration"`
}

type TimesInfo struct {
	TotalDays             int        `json:"totalDays"`
	AverageDurationPerDay string     `json:"averageDurationPerDay"`
	TotalDuration         string     `json:"totalDuration"`
	TimeUnits             []TimeUnit `json:"timeUnits"`
}

// kinde; en; 00:09:37; 2024; 4.9; https://www.youtube.com/watch?v=_EjOHdRihjA; quick video showing how to build Kinde into a React site
type Tutorial struct {
	Topics      string `json:"topics"`
	Language    string `json:"language"`
	Duration    string `json:"duration"`
	Year        string `json:"year"`
	Rank        string `json:"rank"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Platform    string `json:"platform"`
	FileIdCode  string `json:"fileIdCode"`
	Status      string `json:"status"`
}
