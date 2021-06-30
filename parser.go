package main

import (
	"strings"
)

// Parser
func Parser(txt string, search string) string {

	txt = strings.TrimSpace(txt)
	lines := strings.Split(txt, "\n")
	find := ""

	for i := range lines {

		if lines[i][0] != byte('#') {

			sor := strings.Split(lines[i], "=")
			sor[0] = strings.TrimSpace(sor[0])
			sor[1] = strings.TrimSpace(sor[1])

			if sor[0] == search {
				find = sor[1]
			}
		} // if nothing's found,
		// empty string will be returned
	}
	return find
}
