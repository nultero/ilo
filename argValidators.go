package main

import (
	"fmt"
	"os"
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

func isValidFunc(arg string) bool {
	return iterate(arg, ValidArgs)
}

func isValidFileType(arg string) bool {
	return iterate(arg, ValidKinds)
}

func IsFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	}

	return false
}

func InvalidateArg(arg, i string) {
	fmt.Printf("'%v' is an invalid %v argument to bx \n", arg, i)
	os.Exit(1)
}

func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}

	return false
}
