package bx

import (
	"bx/errs"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseString(txt, search string) (string, error) {

	lines := strings.Split(txt, "\n")

	for i := range lines {
		line := strings.TrimSpace(lines[i])
		if line[0] != '#' { // ignore comment blocks

			s := strings.Split(line, "=")
			s[0] = strings.TrimSpace(s[0])
			s[1] = strings.TrimSpace(s[1])

			if s[0] == search {
				return s[1], nil
			}
		}
	}

	s := fmt.Sprintf(
		"'%s' search string not found in search space",
		search,
	)

	return "", errors.New(s)
}

func getNumDaysAhead(conf string) int {

	days, err := ParseString(conf, "days")
	if err != nil {
		errs.ThrowX(err, "'days' setting not found / not properly configured in bx files")
	}

	numDays, nErr := strconv.Atoi(days)
	if nErr != nil {
		errs.ThrowSys(nErr)
	}

	return numDays
}
