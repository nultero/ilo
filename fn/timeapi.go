package fn

import (
	"bx/errs"
	"strconv"
	"strings"
	"time"
)

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

func SetRollDay(unfmt, month string) Day {

	s := strings.TrimSpace(unfmt)
	ss := strings.Split(s, " ")

	di, err := strconv.Atoi(ss[0])
	if err != nil {
		errs.ThrowSys(err)
	}

	return Day{di, month}
}

func RollForwardDay(d Day) Day {
	d.Index++
	if d.Index > getDays(d.MonthName) {
		d.Index = 1
		d.MonthName = getNextMonth(d.MonthName)
	}
	return d
}

func RollBackDay(d Day) Day {
	d.Index--
	if d.Index < 1 {
		d.MonthName = getPrevMonth(d.MonthName)
		d.Index = getDays(d.MonthName)
	}
	return d
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
