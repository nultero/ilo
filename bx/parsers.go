package bx

import (
	"bx/errs"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Finds search string in a text space; ignores lines starting with '#'
// and throws an error if not found.
func ParseString(txt, search string) (string, error) {

	lines := strings.Split(txt, "\n")

	for i := range lines {
		ln := strings.TrimSpace(lines[i])

		if ln[0] != '#' { // ignore comment blocks

			s := strings.Split(ln, "=")
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

// Takes a str or slice of strs as inputs.
func IsEmpty(s interface{}) bool {

	t, ok := s.(string)
	if ok {
		if len(t) == 0 {
			return true
		}
	}

	sl, ok := s.([]string)
	if ok {
		if len(sl) == 0 {
			return true
		}
	}

	return false
}
