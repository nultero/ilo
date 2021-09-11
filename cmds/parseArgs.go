package cmds

import (
	"bx/errs"
	"bx/fn"
	"fmt"
)

func ParseArgs(args []string, b fn.Bus) {

	for i := range args {
		if isValidFunc(args[i]) {
			if fn.IsEmpty(b.Funct) {
				b.Funct = args[i]
			} else {
				errs.ThrowDuplArgError(args[i], b.Funct)
			}

		} else if isValidFileType(args[i]) {
			if fn.IsEmpty(b.FileType) {
				b.FileType = args[i]
			} else {
				errs.ThrowDuplArgError(args[i], b.FileType)
			}

		} else if isFlag(args[i]) { // parse flags here
			fmt.Println("parse flags blah blah")

		} else {
			s := fmt.Sprintf("'%s' is not a valid argument", args[i])
			errs.Throw(s)
		}

	}

	if fn.IsEmpty(b.Funct) {
		errs.Throw("no function passed to bx")
	}

	if fn.IsEmpty(b.FileType) { // in case of only CRUD arg, prompt is also in bus
		promptStr := fmt.Sprintf("%s '%s' needs an argument", b.PromptIcon, b.FileType)
		fmt.Println(promptStr)
		// b.FileType = handleArguments(promptStr)
		// ignore until we get a tui
	}

	//

	switch b.Funct {
	case "add":
	case "edit":
	case "list":
	case "remove":
		fmt.Println(b.Funct)
	}
}
