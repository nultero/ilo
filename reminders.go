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
func CheckReminders(now time.Time, month string, daysOut string) {

	// fmt.Println(Months[TrimMonth(now.Month().String())])

	nDays, _ := strconv.Atoi(daysOut)
	checkRecurrents(now, month, nDays)

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
func checkRecurrents(now time.Time, month string, daysOut int) {

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
						formatStr += "today"
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
