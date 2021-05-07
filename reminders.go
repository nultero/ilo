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
//    -- one time reminders
//    -- recurrent stuff
func CheckReminders(now time.Time, month string, daysOut string, lastDayChecked string) {

	tdysFmt := now.Format("02 Mon")

	if todayHasNotBeenChecked(tdysFmt) {

		nDays, _ := strconv.Atoi(daysOut)
		checkRecurrents(now, nDays)
		// checkOneTimes(now, month, nDays)
		writeCheckedDay(tdysFmt)

	} else {
		fmt.Println("else on the check")
	}

}

var Months = map[string]int{
	"jan": 31,
	"feb": 28, // leaps too
	"mar": 31,
	"apr": 30,
	"may": 31,
	"jun": 30,
	"jul": 31,
	"aug": 31,
	"sep": 30,
	"oct": 31,
	"nov": 30,
	"dec": 31,
}

func TrimMonth(month string) string {
	return month[0:3]
}

// flack  "\033[1;30m%s\033[0m"
// zed    "\033[1;31m%s\033[0m"
// lreen  "\033[1;32m%s\033[0m"
// bellow "\033[1;33m%s\033[0m"
// yurple "\033[1;34m%s\033[0m"
// bagenta"\033[1;35m%s\033[0m"
// peal   "\033[1;36m%s\033[0m"
// khite  "\033[1;37m%s\033[0m"

func checkRecurrents(now time.Time, daysOut int) {

	reminderRead, _ := os.ReadFile(TxtPath("recurrent"))
	recurrents := string(reminderRead)
	recArray := strings.Split(recurrents, "\n")

	today, _ := strconv.Atoi(now.Format("02"))

	for i := range recArray {
		atSplt := strings.Split(recArray[i], "@")

		// not actually n^2, should only be 1 '@' in a line
		for ithSlice := range atSplt {

			piece := strings.TrimSpace(atSplt[ithSlice])
			checkDay, er := strconv.Atoi(piece)

			if er == nil {

				daysWithin := checkDay - today

				if daysWithin >= 0 && daysWithin <= daysOut {

					formatStr := " >> "
					switch daysWithin {
					case 0:
						formatStr += "\033[1;31m" + "today" + "\033[0m"

					case 1:
						formatStr += "1 day"

					default:
						formatStr += fmt.Sprintf("%v days", daysWithin)
					}

					fmt.Println(strings.TrimSpace(atSplt[0]) + formatStr)
				}

			}
		}
	}
}

func readInternals() []string {
	lastData, _ := os.ReadFile(
		ArgCleaner("homedir check"),
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

func writeCheckedDay(now string) {

	fmt.Println(now)

}
