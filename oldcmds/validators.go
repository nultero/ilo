package cmds

import (
	"strings"
)

var ValidArgs = []string{
	"add",
	"defaults",
	"edit",
	"list",
	"remove",
}

var ValidFiles = []string{
	"events",
	"ideas",
	"recurrents",
	"todos",
	"wishlist",
}

var CharMap = []byte("abcdef")

var ValidMapKeys = []string{
	"\033[1;31mtoday\033[0m",
	"\033[1;33m1 day\033[0m",
	"\033[1;32m2 days\033[0m",
	"\033[1;36m3 days\033[0m",
	"\033[1;36m4 days\033[0m",
	"\033[1;36m5 days\033[0m",
}

func _iterate(s string, searchSlice []string) bool {
	for i := range searchSlice {
		if s == searchSlice[i] {
			return true
		}
	}
	return false
}

func isValidFunc(arg string) bool {
	return _iterate(arg, ValidArgs)
}

func isValidFileType(arg string) bool {
	return _iterate(arg, ValidFiles)
}

func isFlag(arg string) bool {
	if strings.Contains(arg, "-") {
		return true
	}

	return false
}