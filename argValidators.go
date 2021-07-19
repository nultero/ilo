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
	"events",
	"ideas",
	"recurrents",
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

func isEmpty(s interface{}) bool {

	t, ok := s.(string)
	if ok {
		if len(t) == 0 {
			return true
		}
	}

	sl, ok := s.([]string)
	if ok {
		if len(sl) == 0 {
			return true
		}
	}

	return false
}
