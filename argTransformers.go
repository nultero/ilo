package main

import (
	"os"
	"strings"
)

// pass in a single string, break apart and
// clean if multiple words
// else just return the clean version
//
// makes it possible to use ultralazy
// arguments for when tired
func ArgCleaner(arg string) string {

	if len(strings.Split(arg, " ")) == 1 {
		arg = canonicals(arg)

	} else { // more than 1 arg, then this
		// rips and cleans words individually

		args := []string(strings.Split(arg, " "))
		for i := range args {
			args[i] = canonicals(args[i])
		}

		arg = ""
		for i := range args {
			arg += canonicals(args[i])
		}
	}

	return arg
}

func canonicals(arg string) string {

	switch arg {

	//dir levels
	case "homedir":
		path, _ := os.UserHomeDir()
		arg = path + "/.config/tailbox/"

	// verbs
	case "a", "ad", "dad", "dd":
		arg = "add"

	case "ed", "edi", "edt":
		arg = "edit"

	case "h", "he", "hel", "halp", "man": // special case?
		arg = "help"

	case "l", "ll", "ls", "lst", "lsa", "read", "rd":
		arg = "list"

	case "rmn", "rem", "rm", "rmv", "rv":
		arg = "remove"
	/////////////////

	// nouns
	case "1", "o", "on", "ot", "one":
		arg = "onetime"

	case "al", "as", "als", "asl", "ali":
		arg = "alias"

	case "ev", "evt", "eve", "evn", "events":
		arg = "event"

	case "i", "di", "id", "ida", "ide":
		arg = "idea"

	case "rc", "rer", "rr", "rec", "recr":
		arg = "recurrent"

	case "td", "tod", "todos":
		arg = "todo"

	case "wl", "ws", "wsl", "wsh", "wish":
		arg = "wishlist"

	// context
	case "cf", "conf", "config":
		arg = "config.txt"

	case "rectxt":
		arg = "recurrent_reminders.txt"

	case "onetxt":
		arg = "one_time_reminders.txt"

	case "todotxt":
		arg = "todos.txt"
		/////////////////

	} /////////////////

	return arg
}
