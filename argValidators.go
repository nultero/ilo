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

func iterate(s string, sl []string) bool {
	for i := range sl {
		if s == sl[i] {
			return true
		}
	}
	return false
}

func IsValidFunc(arg string) bool {
	return iterate(arg, ValidArgs)
}

func IsValidFileType(arg string) bool {
	return iterate(arg, ValidKinds)
}

func IsFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	}

	return false
}

func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}

	return false
}
