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
	TotalDays int `json:"totalDays"`
	AverageDurationPerDay string `json:"averageDurationPerDay"`
	TotalDuration string `json:"totalDuration"`
	TimeUnits []TimeUnit `json:"timeUnits"`
}

// kinde; en; 00:09:37; 2024; 4.9; https://www.youtube.com/watch?v=_EjOHdRihjA; quick video showing how to build Kinde into a React site
type Tutorial struct {

}