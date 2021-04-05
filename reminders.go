package main

import (
	"fmt"
	"os"
	"time"
)

// Reminders does things
// check list :: -->
// one time reminders
// recurrent stuff
func CheckReminders(now time.Time) {

	// fmt.Println(Months[TrimMonth(now.Month().String())])
	checkRecurrents()

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
func checkRecurrents() {

	reminderRead, _ := os.ReadFile(ArgCleaner("homedir recurrent"))
	recurrents := string(reminderRead)

	// nope, going to pass in whole arg list when I get back to fixing this
	fmt.Println(ArgCleaner("homedir recurrent"))
	fmt.Println(recurrents)

}
