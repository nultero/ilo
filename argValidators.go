package main

import (
	"strings"
)

var ValidArgs = []string{
	"add",
	"edit",
	"list",
	"remove",
}

var ValidKinds = []string{
	"event",
	"idea",
	"recurrent",
	"todos",
	"wishlist",
}

func _iterate(s string, sl []string) bool {
	for i := range sl {
		if s == sl[i] {
			return true
		}
	}
	return false
}

func isValidFunc(arg string) bool {
	return _iterate(arg, ValidArgs)
}

func isValidFileType(arg string) bool {
	return _iterate(arg, ValidKinds)
}

func isFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	}

	return false
}

func isEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}

	return false
}
