package main

import (
	"fmt"
	"os"
)

// multiple kinds of args

// verb == priority 1, i.e., add, remove, edit, etc.
// can't be multiple verbs, toss error if true
// noun == priority 2, i.e., reminder, cron, etc
// modifier == p 3

func ArgumentPrioritizer(args []string) []string {

	verbNums, nounNums := 0, 0
	// all args being checked should already be cleaned via ArgCleaner()
	for i := range args {
		if verbs(args[i]) {
			verbNums++
		} else if nouns(args[i]) {
			nounNums++
		}
	}

	// limited to only 1 of each
	if verbNums > 1 || nounNums > 1 {
		// display what was wrong
		fmt.Printf("%v - verbs, %v - nouns \n", verbNums, nounNums)
		bootBackToTerminal(verbNums)
	}

	return args
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
	if faultReason > 1 {
		fault += "verb"
	} else {
		fault += "noun"
	}
	fault += " args have been overloaded"

	Charprint("red", "<!>", fault)
	Charprint("blue", "<!>", "remember: bx can only do one command at a time")
	os.Exit(0)
}
