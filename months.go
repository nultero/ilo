package main

import (
	"strconv"
	"time"
)

type month struct {
	index int
	name  string
	days  int
}

var months = []month{
	{1, "Jan", 31},
	{2, "Feb", 28},
	{3, "Mar", 31},
	{4, "Apr", 30},
	{5, "May", 31},
	{6, "Jun", 30},
	{7, "Jul", 31},
	{8, "Aug", 31},
	{9, "Sep", 30},
	{10, "Oct", 31},
	{11, "Nov", 30},
	{12, "Dec", 31},
}

func rollForwardDay(d int, mn string) (int, string) {
	d++
	if d > getDays(mn) {
		d = 1
		mn = getNextMonth(mn)
	}
	return d, mn
}

func rollBackDay(d int, mn string) (int, string) {
	d--
	if d < 1 {
		mn = getPrevMonth(mn)
		d = getDays(mn)
	}
	return d, mn
}

// Takes a month name string of 3 letters and returns n days. If Feb, will check if leap year or not.
// ( if I'm still using this on a leap year, I'll be ecstatic. Written 2021 )
func getDays(mn string) int {
	for i := range months {
		if months[i].name == mn {
			if months[i].name == "Feb" {
				var leap int
				yr, _ := strconv.Atoi(time.Now().Format("2006"))
				if yr%4 == 0 {
					leap = 29
				} else {
					leap = 28
				}
				return leap
			}

			return months[i].days
		}
	}
	return 0
}

func getMonths() []string {
	var mn []string
	for i := range months {
		mn = append(mn, months[i].name)
	}
	return mn
}

func getPrevMonth(mn string) string {
	i := getIndex(mn)
	if i-1 == 0 {
		i = 12
	} else {
		i = i - 1
	}

	return months[i].name
}

func getNextMonth(mn string) string {
	i := getIndex(mn)
	if i+1 == 13 {
		i = 1
	} else {
		i++
	}

	return months[i].name
}

func getNextMonthsDays(mn string) int {
	for i := range months {
		if months[i].name == mn {

			nxt_i := months[i].index + 1
			if nxt_i == 13 {
				nxt_i = 1
			}

			if months[nxt_i].name == "Feb" {
				var leap int
				yr, _ := strconv.Atoi(time.Now().Format("2006"))
				if yr%4 == 0 {
					leap = 29
				} else {
					leap = 28
				}
				return leap
			}

			return months[nxt_i].days
		}
	}
	return 0
}

// Takes month name string of 3 letters, returns the index of that month, starting at 1 for Jan.
func getIndex(mn string) int {
	for i := range months {
		if months[i].name == mn {
			return months[i].index
		}
	}
	return 0
}

func isInRange() bool {
	return false
}
