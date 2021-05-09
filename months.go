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

var Months = []month{
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

// Takes a month name string of 3 letters and returns n days. If Feb, will check if leap year or not.
// ( if I'm still using this on a leap year, I'll be ecstatic. Written 2021 )
func GetDays(mn string) int {
	for i := range Months {
		if Months[i].name == mn {
			if Months[i].name == "Feb" {
				var leap int
				yr, _ := strconv.Atoi(time.Now().Format("2006"))
				if yr%4 == 0 {
					leap = 29
				} else {
					leap = 28
				}
				return leap
			}

			return Months[i].days
		}
	}
	return 0
}

func GetNextMonthsDays(mn string) int {
	for i := range Months {
		if Months[i].name == mn {

			nxt_i := Months[i].index + 1
			if nxt_i == 13 {
				nxt_i = 1
			}

			if Months[nxt_i].name == "Feb" {
				var leap int
				yr, _ := strconv.Atoi(time.Now().Format("2006"))
				if yr%4 == 0 {
					leap = 29
				} else {
					leap = 28
				}
				return leap
			}

			return Months[nxt_i].days
		}
	}
	return 0
}

// Takes month name string of 3 letters, returns the index of that month, starting at 1 for Jan.
func GetIndex(mn string) int {
	for i := range Months {
		if Months[i].name == mn {
			return Months[i].index
		}
	}
	return 0
}
