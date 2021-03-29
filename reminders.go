package main

import (
	"time"
)

// Reminders does things
// check list :: -->
// one time reminders
// recurrent stuff
func Reminders(now time.Time) {

	// fmt.Println(Months[TrimMonth(now.Month().String())])

}

var Months = map[string]int{
	"Jan": 31,
	"Feb": 28, // leaps too
	"Mar": 31,
	"Apr": 30,
	"May": 31,
	"Jun": 30,
	"Jul": 31,
	"Aug": 31,
	"Sep": 30,
	"Oct": 31,
	"Nov": 30,
	"Dec": 31,
}

func TrimMonth(month string) string {
	return month[0:3]
}
func oneTimes() {

}
