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
func CheckReminders(now time.Time, month string) {

	// fmt.Println(Months[TrimMonth(now.Month().String())])
	checkRecurrents(now, month)

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
func checkRecurrents(now time.Time, month string) {

	reminderRead, _ := os.ReadFile(TxtPath("recurrent"))
	recurrents := string(reminderRead)
	recArray := strings.Split(recurrents, "\n")

	var dates []int

	for i := range recArray {
		// the format should be, i.e.
		// 'phone bill @ 9'
		// meaning a recurrent reminder would be
		// something repeatedly occurring on the 9th
		atSplt := strings.Split(recArray[i], "@")

		for ithSlice := range atSplt { // not actually n^2
			// should only ever be a single '@' char on a line
			piece := strings.TrimSpace(atSplt[ithSlice])
			checkDay, er := strconv.Atoi(piece)
			if er == nil { // adds int days of, i.e., the 9
				dates = append(dates, checkDay)
			}
		}
	}

	fmt.Println(dates)

	// fmt.Println(recArray)
	// for i := range Months {
	// 	fmt.Println(Months[i])
	// }
	// fmt.Println(Months[2].month)
	// mth := strings.ToLower(month)
	// fmt.Println(Months[mth])
}
