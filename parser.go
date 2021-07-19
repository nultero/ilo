package main

import (
	"fmt"
	"strings"
)

// Parser
func parser(txt string, search string) string {

	txt = strings.TrimSpace(txt)
	lines := strings.Split(txt, "\n")
	find := ""

	for i := range lines {

		if lines[i][0] != '#' {

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

func cleanPrint(s string) {
	lines := strings.Split(s, "\n")

	firstHalf := []string{}
	dates := []string{}

	for i := range lines {
		ln := strings.Split(lines[i], "@")
		firstHalf = append(firstHalf, ln[0])
		dates = append(dates, ln[1])
	}

	mx := maximum(firstHalf)

	for i := range firstHalf {
		spaces := strings.Repeat(" ", mx-len(firstHalf[i]))
		fmt.Println(firstHalf[i], spaces, "@", dates[i])
	}
}

func maximum(s []string) int {
	x := 0
	for i := range s {
		if len(s[i]) > x {
			x = len(s[i])
		}
	}
	return x
}
