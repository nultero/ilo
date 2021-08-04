package bx

import (
	"errors"
	"fmt"
	"strings"
)

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
