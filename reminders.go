package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Sees if current day has already had checks run, and if not,
// runs the date against events and recurrent content set up in
// the bx dir.
func runChecks(month, today, config string) {

	cache := readCache()
	day := fmt.Sprintf("%s %s", month, today)

	if cacheIsOld(cache[0], day) {

		advance := getDaysInAdvance(config)
		// ev := checkEvents(day, "events", advance)
		// rc := checkEvents(day, "recurrents", advance)

		m := merge(
			checkEvents(day, "events", advance),
			checkEvents(day, "recurrents", advance),
		)

		for k := range ValidMapKeys {
			if !isEmpty(m[ValidMapKeys[k]]) {
				fmt.Println(ValidMapKeys[k], "->", m[ValidMapKeys[k]])
			}
		}
		// ev = append(ev, checkEvents("30 Jul", 3)...)
		// cacheResults(day, ev)
		// put in autocleaner for events past

	} else {
		if !isEmpty(cache[1:]) {
			fmt.Println(cache[1:])
		}
	}
}

////////////////////////////
func readCache() []string {
	lastData, r := os.ReadFile(pathGlob("check"))
	if r != nil {
		throwErr(r)
	}
	return strings.Split(string(lastData), "\n")
}

func getDaysInAdvance(conf string) int {
	advance, r := strconv.Atoi(parser(conf, "days"))
	if r != nil {
		throwErr(r)
	}

	if advance < len(ValidMapKeys) {
		return advance
	} else {
		e := errors.New("invalid number of days in advance")
		throwErr(e)
	}
	return advance
}

func cacheIsOld(cacheTop, day string) bool {

	return true // just for testing
	// if cacheTop != day {
	// 	return true
	// }
	// return false
} /////////////////////////

//  |||||  ||    ||
//  ||     ||    ||
//  ||||   ||    ||
//  ||      ||  ||
//  |||||     ||

func checkEvents(today, kind string, advance int) map[string]string {

	ev := map[string]string{}

	data, _ := os.ReadFile(pathGlob(kind))
	lines := strings.Split(string(data), "\n")

	events, dues := breakAt(lines)
	keyedEvents := breakDates(dues, events)

	d, mn := parseDay(today)

	// autocleaner
	if kind == "events" {
		for i := range dues {
			r := strings.TrimSpace(dues[i])
			s := strings.Split(r, " ")
			day, _ := strconv.Atoi(s[0])

			if s[1] == mn { //same month
				if day < d { //if previous day of same month
					autoclean(i) // needs a rollback function actually
				} // work off of both day-based rollback, and cache rolledback stuff
			} //    so none of it is ever "lost"
		}
	}

	for i := 0; i <= advance; i++ {
		if !isEmpty(keyedEvents[d]) {
			s := preFmt(keyedEvents[d], i)
			ev[s[0]] += s[1]
		}
		d, mn = rollForwardDay(d, mn)
	}

	return ev
}

func merge(ev, rc map[string]string) map[string]string {

	m := map[string]string{}

	for k, v := range ev {
		m[k] = v
	}

	for k, v := range rc {
		if !isEmpty(m[k]) {
			// fmt.Println("testing out how to colorate from k", k)
			m[k] += colorate(v, []int{100, 170, 200})
		}
	}

	for k, v := range m {
		m[k] = prettify(v)
	}

	return m
}

func preFmt(data string, numDays int) []string {
	return []string{colorFmt(numDays), data}
}

func colorFmt(numDays int) string {
	var tmp string
	switch numDays {
	case 0:
		tmp = "\033[1;31m" + "today" + "\033[0m"

	case 1:
		tmp = "\033[1;33m" + "1 day" + "\033[0m"

	case 2:
		tmp = fmt.Sprintf("\033[1;32m%v days\033[0m", numDays)

	default:
		tmp = fmt.Sprintf("\033[1;36m%v days\033[0m", numDays)
	}
	return tmp
}

func colorate(s string, colorsRGB []int) string {

	c := []string{}
	for i := range colorsRGB {
		c = append(c,
			strconv.Itoa(colorsRGB[i]),
		)
	}

	return fmt.Sprintf("\033[38;2;%s;%s;%sm%s\033[0m", c[0], c[1], c[2], s)
}

func prettify(s string) string {
	tmp := ""
	line := strings.Split(s, "<:>")

	for i := range line {
		s := strings.TrimSpace(line[i])
		tmp += s + ", "
	}

	// shaves off the comma / space
	tmp = strings.TrimSpace(tmp)
	tmp = string(tmp[:len(tmp)-3])
	return tmp
}

func cacheResults(today string, data []string) {

	tmp := today + "\n"

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

func autoclean(index int) {
	p := pathGlob("events")
	f, _ := os.ReadFile(p)
	lines := strings.Split(string(f), "\n")
	cleaned := make([]string, len(lines)-1)
	cleaned = append(lines[:index], lines[index+1:]...)
	writeOut(p, cleaned, false)
}
