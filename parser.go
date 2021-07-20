package main

import (
	"fmt"
	"strconv"
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

	firstHalf, dates := breakAt(lines)

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

func breakAt(lines []string) ([]string, []string) {

	lefts, rights := []string{}, []string{}

	for i := range lines {
		ln := strings.Split(lines[i], "@")
		lefts = append(lefts, ln[0])
		rights = append(rights, ln[1])
	}

	return lefts, rights
}

func breakDates(dates []string) ([]int, []string) {

	days, months := []string{}, []string{}

	for i := range dates {
		ln := strings.TrimSpace(dates[i])
		s := strings.Split(ln, " ")
		days = append(days, s[0])
		months = append(months, s[1])
	}

	conv := []int{}

	for i := range days {
		n, _ := strconv.Atoi(days[i])
		conv = append(conv, n)
	}

	return conv, months
}

func parseDay(s string) (int, string) {
	d := strings.TrimSpace(s)
	arr := strings.Split(d, " ")
	n, _ := strconv.Atoi(arr[1])

	return n, arr[0]
}
