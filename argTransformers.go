package main

import (
	"os"
	"strings"
)

// Transform cleans up inter-function and CLI args & paths,
// these are the canonical function params --
// "why use lot word few word trick"
// this is more for passing in lazy/tired CLI arguments than for source
func ArgCleaner(arg string) string {

	//might be cleaner with autocompletes

	switch arg {

	case "1", "o", "on", "ot", "one":
		arg = "onetime"

	case "a", "ad", "dad", "dd":
		arg = "add"

	case "al", "as", "als", "asl", "ali":
		arg = "alias"

	case "ed", "edi", "edt":
		arg = "edit"

	case "ev", "evt", "eve", "evn", "events":
		arg = "event"

	case "i", "di", "id", "ida", "ide":
		arg = "idea"

	case "h", "he", "hel", "halp", "man":
		arg = "help"

	case "l", "ll", "ls", "lst", "lsa", "read", "rd":
		arg = "list"

	case "rc", "rer", "rr", "rec", "recr":
		arg = "recurrent"

	case "rmn", "rem", "rm", "rmv", "rv":
		arg = "remove"

	case "sd", "sod":
		arg = "sad"

	case "td", "tod", "todos":
		arg = "todo"

	case "wl", "ws", "wsl", "wsh", "wish":
		arg = "wishlist"

	// special case for pulling config
	case "homedir",
		"homedir cf",
		"homedir conf",
		"homedir config",
		"cf",
		"conf",
		"config":
		arg = path(arg)

	}

	return arg
}

func path(arg string) string {

	translatedArg := ""

	spl := strings.Split(arg, " ")
	splitLength := len(spl)

	if splitLength >= 1 {

		if spl[0] == "homedir" {
			path, _ := os.UserHomeDir()
			path += "/.config/tailbox/"
			translatedArg += path
		}

		if splitLength > 1 {
			arg = spl[1]
		}

	}

	switch arg {
	case "cf", "conf", "config":
		translatedArg += "config.txt"

	case "r", "rc", "rec", "recurrent":
		translatedArg += "recurrent_rmds.txt"

	case "o", "on", "one", "one time", "1":
		translatedArg += "one_time_rmds.txt"

	case "td", "tod", "todo", "todos":
		translatedArg += "todos.txt"
	}

	return translatedArg
}
