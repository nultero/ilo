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

		m := merge(
			checkEvents(day, "events", advance),
			checkEvents(day, "recurrents", advance),
		)

		vals := []string{day}

		for k := range ValidMapKeys {
			if !isEmpty(m[ValidMapKeys[k]]) {
				s := ValidMapKeys[k] + " -> " + m[ValidMapKeys[k]]
				fmt.Println(s)
				vals = append(vals, s)
			}
		}

		cacheResults(vals)

	} else {
		c := cache[1:]
		if !isEmpty(c) {
			for i := range c {
				fmt.Println(c[i])
			}
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
	if cacheTop != day {
		return true
	}
	return false
} /////////////////////////

//  |||||  ||    ||
//  ||     ||    ||
//  ||||   ||    ||
//  ||      ||  ||
//  |||||     ||

func checkEvents(today, funcType string, advance int) map[string]string {

	ev := map[string]string{}

	data, _ := os.ReadFile(pathGlob(funcType))
	lines := strings.Split(string(data), "\n")

	events, dues := breakAt(lines)
	keyedEvents := breakDates(dues, events)

	d, mn := parseDay(today)

	// autocleaner logic for nonrecurrent events
	if funcType == "events" {

		adv := -1 * advance

		for i := range dues {

			r := strings.TrimSpace(dues[i])
			s := strings.Split(r, " ")
			day, _ := strconv.Atoi(s[0])

			_d, _mn := d, mn // reset daycheck vars per item in range

			for _i := 0; _i >= adv; _i-- {
				_d, _mn = rollBackDay(_d, _mn)

				if s[1] == _mn { // matches month, and then day-based
					if day == _d { // on rollback
						autoclean(i)
					}
				}
			}
		}
	} //// end autocleaner

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
	line = line[:len(line)-1]

	for i := range line {
		s := strings.TrimSpace(line[i])
		tmp += s + ", "
	}

	// shaves off the comma / space
	tmp = strings.TrimSpace(tmp)
	tmp = string(tmp[:len(tmp)-1])
	return tmp
}

func cacheResults(data []string) {

	tmp := ""
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

	s := colorate("** bx's autoclean has run", []int{90, 210, 110})
	fmt.Println(s)
}
