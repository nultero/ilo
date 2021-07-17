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

type Bus struct {
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
	p, _ := filepath.Abs(HandleHomePath(PATH))
	config := Configure()

	b := Bus{
		PromptIcon: Parser(config, "promptIcon") + " ",
		FileType:   "",
		Funct:      "",
		Path:       p,
		Help:       false,
	}

	if len(args) == 0 { // zero args calls checks, prints reminders if within config's threshold, and current date
		//						  and if already called today, prints cached info

		now := time.Now()
		todaysFormatDate := now.Format("02 Mon")
		month := now.Month().String()[0:3]
		fmt.Println(b.PromptIcon, month, todaysFormatDate)

		//checks
		//checks
		//checks

	} else { // if bx is called with args, it will parse args instead of doing its checks

		for i := range args {
			if IsValidFunc(args[i]) {
				if IsEmpty(b.Funct) {
					b.Funct = args[i]
				} else {
					throwDuplArgError(args[i], b.Funct)
				}

			} else if IsValidFileType(args[i]) {
				if IsEmpty(b.FileType) {
					b.FileType = args[i]
				} else {
					throwDuplArgError(args[i], b.FileType)
				}

			} else if IsFlag(args[i]) { // parse flags here
				fmt.Println("parse flags blah blah")

			} else {
				s := fmt.Sprintf("'%s' is not a valid argument", args[i])
				fmt.Println(RedError(), s)
				os.Exit(1)
			}

		}

		if IsEmpty(b.Funct) {
			fmt.Println(RedError(), "no function passed to bx")
		}

		if IsEmpty(b.FileType) { // in case of only CRUD arg, prompt is also in bus
			promptStr := fmt.Sprintf("%s '%s' needs an argument", b.PromptIcon, b.FileType)
			b.FileType = HandleArguments(promptStr)
		}

		_eval(b)
	}
}

func _eval(b Bus) {
	switch b.Funct {
	case "add":
		Add(b)

	case "edit":
		Edit(b)

	case "list":
		List(b.Path, b.FileType)

	case "remove":
		Remove(b)

	case "test", "home", "config", "help":
		fmt.Println("poo")

	default:
		fmt.Println("dump")
	}
}

func throwDuplArgError(this, prevFound string) {
	fmt.Printf("! >> '%s' found, but already passed '%s' as argument \n", this, prevFound)
	fmt.Println(RedError(), "cannot have two of the type of argument")
	os.Exit(1)
}

func RedError() string {
	return "\033[31;1;4merror:\033[0m"
}
