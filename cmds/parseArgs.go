package cmds

import (
	"bx/errs"
	"bx/fn"
	"fmt"

	"github.com/manifoldco/promptui"
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
				b.FileType = args[i] + ".txt"
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
		promptStr := fmt.Sprintf("%s '%s' needs a filetype argument", b.PromptIcon, b.Funct)
		b.FileType = handleArgs(promptStr)
	}

	//

	switch b.Funct {
	case "add":
		add(b)

	case "edit":
	case "list":
		ls(b)

	case "remove":
		fmt.Println(b.Funct)
	}
}

func handleArgs(explanation string) string {

	prompt := promptui.Select{
		Label: explanation,
		Items: ValidFiles,
	}

	_, s, err := prompt.Run()

	if err != nil {
		s = fmt.Sprintf("prompt failure: %v", err)
		errs.ThrowQuiet(s)
	}

	if s == ValidFiles[2] {
		s = "recurrent_reminders"
	}

	return s + ".txt"
}
