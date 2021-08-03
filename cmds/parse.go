package cmds

func ParseArgs(args []string) {

}

// for i := range args {
// 	if isValidFunc(args[i]) {
// 		if isEmpty(b.Funct) {
// 			b.Funct = args[i]
// 		} else {
// 			throwDuplArgError(args[i], b.Funct)
// 		}

// 	} else if isValidFileType(args[i]) {
// 		if isEmpty(b.FileType) {
// 			b.FileType = args[i]
// 		} else {
// 			throwDuplArgError(args[i], b.FileType)
// 		}

// 	} else if isFlag(args[i]) { // parse flags here
// 		fmt.Println("parse flags blah blah")

// 	} else {
// 		s := fmt.Sprintf("'%s' is not a valid argument", args[i])
// 		fmt.Println(redError(), s)
// 		os.Exit(1)
// 	}

// }

// if isEmpty(b.Funct) {
// 	fmt.Println(redError(), "no function passed to bx")
// }

// if isEmpty(b.FileType) { // in case of only CRUD arg, prompt is also in bus
// 	promptStr := fmt.Sprintf("%s '%s' needs an argument", b.PromptIcon, b.FileType)
// 	b.FileType = handleArguments(promptStr)
// }

// _eval(b)
