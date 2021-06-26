package main

// import (
// 	"flag"
// 	"fmt"
// 	"time"
// )

// //
// //
// //---- baselines ----//
// //
// // files:
// //	 config.txt
// //   one_time_rmds.txt
// //   recurrent_rmds.txt
// //   todos.txt
// //   /
// //	 >> export and import data for when you mess up
// //   /
// //   kinds of crons: server side, home side, emb for embed? / pi
// //
// //   ^ all should be fairly transparent from outside tailbox, in native formats or plaintext
// //
// //
// // change default directory
// // -dir flag for pwd on bx's $home
// // -h flag for forgetting all cmds

// //  ||
// //  |||||  ||  ||
// //  || ||    ||
// //  |||||  ||  ||

// func oldmain() {

// 	flag.Parse()
// 	configs := Configure(ArgCleaner("homedir conf"))
// 	prompt := string(Parser(configs, "=", "default_prompt") + " ")

// 	if len(flag.Args()) == 0 { // bx will only check stuff if called passively
// 		//                        i.e., any args will bypass these checks
// 		now := time.Now()
// 		todaysFormatDate := now.Format("02 Mon")
// 		month := now.Month().String()[0:3]
// 		fmt.Println(prompt, month, todaysFormatDate)

// 		daysOut := string(Parser(configs, "=", "default_days_check"))
// 		CheckReminders(now, month, daysOut)

// 	} else {

// 		for i := range flag.Args() {
// 			flag.Args()[i] = ArgCleaner(flag.Args()[i])
// 		}

// 		sortedArgs := ArgumentPrioritizer(flag.Args())
// 		eval(sortedArgs, prompt)
// 	}
// }

// func eval(args []string, promptstr string) {

// 	if args[0] != "test" && args[0] != "help" {
// 		for i := 0; i < len(args); i++ { // prompts
// 			if args[i] == "" {
// 				args[i] = prompt(i)
// 			}
// 		}
// 	}

// 	switch args[0] { // final args are digested into eval
// 	case "add":
// 		Add(args[1], promptstr)
// 	case "edit":
// 		Edit(args[1], promptstr)
// 	case "list":
// 		List(args[1])
// 	case "remove":
// 		Remove(args[1])
// 	case "test":
// 		Test()
// 	}
// }

///   ||||            |   |           ||||  |  |  |
///   |               |   |           |  |     | |
///   ||    ||||   ||||   |     ||||  ||||  |  ||
///   |     ||  |  |  |   |     |  |     |  |  | |
///   ||||  ||  |  ||||   ||||  ||||  ||||  |  |  |
///
///
//
// All below are for Main()'s input cleansing.
// All specific logic is elsewhere, in utils or exported funcs

////////////
// func prompt(i int) string {

// 	tmp := ""
// 	msg := "options for "
// 	fmt.Println(">", strings.Repeat("-", 15))

// 	if i == 0 { // verb was missing
// 		promptstr := "> " + string(msg+"bx's verbs")
// 		tmp = verbsPrompt(promptstr)

// 	} else { // noun was missing
// 		promptstr := "> " + string(msg+"bx's nouns")
// 		tmp = nounsPrompt(promptstr)
// 	}

// 	return tmp
// }

// func verbsPrompt(promptstr string) string {
// 	opts := []string{
// 		"add",
// 		"edit",
// 		"list",
// 		"help",
// 		"test",
// 		"remove"}

// 	return HandlesOpts(TinyArray(promptstr), opts)
// }

// func nounsPrompt(promptstr string) string {
// 	opts := []string{
// 		"onetime",
// 		"alias",
// 		"event",
// 		"idea",
// 		"recurrent",
// 		"todo",
// 		"wishlist"}

// 	return HandlesOpts(TinyArray(promptstr), opts)
// }
