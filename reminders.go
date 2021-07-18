package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Reminders
//    -- check list :
//    -- events
//    -- recurrent stuff
func CheckReminders(now time.Time, month string, daysOut string) {

	tdysFmt := now.Format("02 Mon")
	mn := now.Month().String()[0:3]

	// the newline slice at the end seems to be malfunctioning
	// also a bug in the dates' arithmetic somewhere

	if todayHasNotBeenChecked(tdysFmt) {

		intDays, _ := strconv.Atoi(daysOut)
		intToday, _ := strconv.Atoi(now.Format("02"))

		recs := checkRecurrents(intToday, intDays, mn)
		printall(recs)

		// ots := checkOneTimes(now, month, nDays)

		writeCheckedDay(tdysFmt, recs)

	} else {
		// this is on the already-written file
		// so skip the date line and just read content
		printall(readInternals()[1:])
	}
}

func printall(data []string) {
	for i := range data {
		fmt.Println(data[i])
	}
}

//  |||      |||||  |||||
//  |||||    ||     ||
//  || ||    ||||   ||
//  ||  ||   ||     ||
//  ||   ||  |||||  |||||

func checkRecurrents(today int, daysOut int, mn string) []string {

	recs, _ := os.ReadFile("recurrent")
	recArray := strings.Split(
		string(recs), "\n",
	)

	var targetStrings []string

	for ln := range recArray {
		atSplt := strings.Split(recArray[ln], "@")

		for i := range atSplt {

			splLine := strings.TrimSpace(atSplt[i])
			day, er := strconv.Atoi(splLine)

			if er == nil {

				daysWithin := day - today

				if passesDayCheck(today, day, daysOut, mn) {

					formatStr := " >> "
					// still have to fix in case of negative
					formatStr += colorFmt(daysWithin)

					tmp := strings.TrimSpace(atSplt[0]) + formatStr
					targetStrings = append(targetStrings, tmp)

				}

			}
		}
	}

	return targetStrings
}

// Logic for the case where the target day being checked against is a low number, like the 1st
// of a month, and the current day is the 30th or so -- the original code didn't "wraparound".
// This handles the overall threshold of days for both recurrents and one-time checks.
func passesDayCheck(today int, targ int, daysOut int, mn string) bool {
	var daysWithin int
	if (targ - today) < 0 {
		nxtMonths := GetNextMonthsDays(mn)
		daysWithin = (targ + nxtMonths) - today
	} else {
		daysWithin = targ - today
	}

	if daysWithin <= daysOut {
		return true
	} else {
		return false
	}
}

func colorFmt(numDays int) string {
	tmp := ""
	switch numDays {
	case 0:
		tmp += "\033[1;31m" + "today" + "\033[0m"

	case 1:
		tmp += "\033[1;33m" + "1 day" + "\033[0m"

	case 2:
		tmp += fmt.Sprintf("\033[1;32m%v days\033[0m", numDays)

	default:
		tmp += fmt.Sprintf("\033[1;36m%v days\033[0m", numDays)
	}
	return tmp
}

func readInternals() []string {
	lastData, _ := os.ReadFile(
		pathGlob("check"),
	)
	return strings.Split(string(lastData), "\n")
}

func todayHasNotBeenChecked(today string) bool {
	lastDay := readInternals()[0]
	if lastDay == today {
		return false
	} else {
		return true
	}
}

func writeCheckedDay(now string, data []string) {

	tmp := now + "\n"

	for i := range data {
		tmp += data[i] + "\n"
	}

	// slice off last newline
	tmp = tmp[:len(tmp)-1]

	os.WriteFile(
		pathGlob("check"),
		[]byte(tmp),
		0644,
	)
}
