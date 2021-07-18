package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// bx's PATH should be, for instance:
const PATH = "~/.nultero/tailbox/"

// I should put the PATH in some kind of init script

type bus struct {
	PromptIcon string
	FileType   string
	Funct      string
	Path       string
	Help       bool
}

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	args := os.Args[1:]
	p, _ := filepath.Abs(handleHomePath(PATH))
	config := configure()
	icon := Parser(config, "promptIcon") + " "

	if len(args) == 0 { // zero args calls checks, prints reminders if within config's threshold, and current date
		//						  and if already called today, prints cached info

		now := time.Now()
		todaysDate := now.Format("02 Mon")
		month := now.Month().String()[0:3]
		fmt.Println(icon, month, todaysDate)

		runChecks(month, todaysDate)
		// check cache if already done today
		// run new checks if not already run

	} else { // if bx is called with args, it will parse args instead of doing its checks

		b := bus{
			PromptIcon: icon,
			FileType:   "",
			Funct:      "",
			Path:       p,
			Help:       false,
		}

		for i := range args {
			if isValidFunc(args[i]) {
				if isEmpty(b.Funct) {
					b.Funct = args[i]
				} else {
					throwDuplArgError(args[i], b.Funct)
				}

			} else if isValidFileType(args[i]) {
				if isEmpty(b.FileType) {
					b.FileType = args[i]
				} else {
					throwDuplArgError(args[i], b.FileType)
				}

			} else if isFlag(args[i]) { // parse flags here
				fmt.Println("parse flags blah blah")

			} else {
				s := fmt.Sprintf("'%s' is not a valid argument", args[i])
				fmt.Println(redError(), s)
				os.Exit(1)
			}

		}

		if isEmpty(b.Funct) {
			fmt.Println(redError(), "no function passed to bx")
		}

		if isEmpty(b.FileType) { // in case of only CRUD arg, prompt is also in bus
			promptStr := fmt.Sprintf("%s '%s' needs an argument", b.PromptIcon, b.FileType)
			b.FileType = handleArguments(promptStr)
		}

		_eval(b)
	}
}

func _eval(b bus) {
	switch b.Funct {
	case "add":
		add(b)

	case "edit":
		edit(b)

	case "list":
		list(b.Path, b.FileType)

	case "remove":
		remove(b)

	case "test", "home", "config", "help":
		fmt.Println("poo")

	default:
		fmt.Println("dump")
	}
}

func throwDuplArgError(this, prevFound string) {
	fmt.Printf("! >> '%s' found, but already passed '%s' as argument \n", this, prevFound)
	fmt.Println(redError(), "cannot have two of the type of argument")
	os.Exit(1)
}

func redError() string {
	return "\033[31;1;4merror:\033[0m"
}

func throwErr(r error) {
	fmt.Println(redError(), r)
	os.Exit(1)
}
