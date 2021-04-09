package main

import (
	"os"
)

// multiple kinds of args

// verb == high priority, i.e., add, remove, edit, etc.
// can't be multiple verbs, toss error if true
// noun == low priority, i.e., reminder, cron, etc
// modifier == lowest priority

func ArgumentPrioritizer(args []string) []string {

	verbNums, nounNums := 0, 0

	// only ever gonna be size 2, a verb & noun
	var sortedArgs = make([]string, 2)

	for i := range args {
		if verbs(args[i]) {
			sortedArgs[0] = args[i]
			verbNums++
		} else if nouns(args[i]) {
			sortedArgs[1] = args[i]
			nounNums++
		}
	}

	// limited to only 1 of each
	if verbNums > 1 || nounNums > 1 {
		// display what was wrong, sort of crappy err logging
		bootBackToTerminal(verbNums)
	}

	return sortedArgs
}

func verbs(match string) bool {
	switch match {
	case "add",
		"edit",
		"list",
		"remove":
		return true
	}
	return false
}

func nouns(match string) bool {
	switch match {
	case "alias",
		"event",
		"idea",
		"onetime",
		"recurrent",
		"todo",
		"wishlist":
		return true
	}
	return false
}

func bootBackToTerminal(faultReason int) {

	fault := "bx's "
	infoStr := "fault: bx can only evaluate one "
	if faultReason > 1 {
		fault += "verb"
		infoStr += "verb"

	} else {
		fault += "noun"
		infoStr += "noun"
	}
	fault += " args have been overloaded"
	infoStr += " at a time"

	Charprint("red", "<!>", fault)
	Charprint("blue", "<!>", infoStr)
	os.Exit(0)
}
