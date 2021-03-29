package main

import "fmt"

// multiple kinds of args

// verb == priority 1, i.e., add, remove, edit, etc.
// can't be multiple verbs, toss error if true
// noun == priority 2, i.e., reminder, cron, etc
// modifier == p 3

func PriorityCrunch(args []string) []string {

	verbNums, nounNums := 0, 0
	// all args being checked should already be cleaned via Transformer()
	for i := range args {
		if verbs(args[i]) {
			verbNums++
		} else if nouns(args[i]) {
			nounNums++
		}
	}

	fmt.Printf("%v - verbs, %v -nouns", verbNums, nounNums)

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
		"one time",
		"recurrent",
		"todo",
		"wishlist":
		return true
	}
	return false
}

func tooManyVerbs(args []string) {

}
