package main

import (
	"strings"
)

// Parser ... what it says on the tin
func Parser(txt string, splitter string, search string) string {

	txt = strings.TrimSpace(txt)
	splitNewlines := strings.Split(txt, "\n")
	find := ""

	for i := range splitNewlines {

		sor := strings.Split(splitNewlines[i], splitter)
		sor[0] = strings.TrimSpace(sor[0])
		sor[1] = strings.TrimSpace(sor[1])

		if sor[0] == search {
			if sor[0][0] != byte('#') {
				find = sor[1]
			}
		} // if nothing's found,
		// empty string will be returned
	}
	return find
}
