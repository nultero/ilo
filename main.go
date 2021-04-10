package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

//
//
//---- baselines ----//
//
// files:
//	 config.txt
//   one_time_rmds.txt
//   recurrent_rmds.txt
//   todos.txt
//   /
//   aliases.txt <-- later
//   kinds of aliases: nav, helpers, formats, shorteners, tooling, funcs
//   kinds of crons: server side, home side, emb for embed? / pi
//   kinds of backups: track bys, gitlike, dotfiles
//
//   ^ all should be fairly transparent from outside tailbox, in native formats or plaintext
//
// * AUTOCOMPLETES IN CONFIG
// * use last confirm to tidy up dotfile migration
//
// change default directory
// -dir flag for pwd on bx's $home
// -h flag for forgetting all cmds

func main() {

	flag.Parse()
	configs := ConfigureStuff(ArgCleaner("homedir conf"))
	prompt := string(Parser(configs, "=", "default_prompt") + " ")

	if len(flag.Args()) == 0 { // bx will only check stuff if called passively
		//                        i.e., any args will bypass these checks
		now := time.Now()
		formattedTime := now.Format("02 Mon")
		month := TrimMonth(now.Month().String())
		fmt.Println(prompt, month, formattedTime)

		CheckReminders(now)

	} else {

		for arg := range flag.Args() {
			// i.e., 'rm' returns 'remove' from Cleaner()
			flag.Args()[arg] = ArgCleaner(flag.Args()[arg])
		}

		sortedArgs := ArgumentPrioritizer(flag.Args())
		eval(sortedArgs)
	}
}

func eval(args []string) {

	for i := 0; i < len(args); i++ { // prompts
		if args[i] == "" {
			args[i] = prompt(i)
		}
	}

	switch args[0] { // final args are digested into eval
	case "add":
		Add(args[1])
	case "edit":
		fmt.Println("editing thing")
	case "list":
		fmt.Println("listing thing")
	case "remove":
		fmt.Println("removing thing")
	}
}

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
func prompt(i int) string {

	tmp := ""
	msg := "options for "
	if i == 0 { // verb was missing
		Charprint("nil", ">! ", string(msg+"bx's verbs"))
		tmp = verbsPrompt()

	} else { // noun was missing
		Charprint("nil", ">! ", string(msg+"bx's nouns"))
		tmp = nounsPrompt()
	}

	return tmp
}

func verbsPrompt() string {
	opts := []string{
		"add",
		"edit",
		"list",
		"remove"}
	tmp := ""
	printsOptions(opts)
	tmp = HandleOptionsInput()
	return optionsChecks(tmp, opts)
}

func nounsPrompt() string {
	opts := []string{
		"onetime",
		"alias",
		"event",
		"idea",
		"recurrent",
		"todo",
		"wishlist"}
	tmp := ""
	printsOptions(opts)
	tmp = HandleOptionsInput()
	return optionsChecks(tmp, opts)
}

func printsOptions(opts []string) {
	for i := range opts {
		Charprint("nil", "> ", string(
			fmt.Sprint(i+1)+ //index
				". "+ // "1. remove" etc
				string(opts[i])))
	}
}

// Exports a string that is unchecked against any args.
// Pass to a separate function to map cleaner args in passthrough.
func HandleOptionsInput() string { // only takes/rds string
	nex, _, e := bufio.NewReader(os.Stdin).ReadLine()
	for e != nil {
		fmt.Println(e)
	}
	return string(nex)
}

func optionsChecks(tmp string, opts []string) string {

	// try to convert str to int opt 1st
	opt, e := strconv.Atoi(tmp)

	// string input
	if e != nil {
		tmp = ArgCleaner(tmp) // clean again

		///	// int input
	} else if 0 < opt && opt <= len(opts) {
		tmp = opts[opt-1] // i - 1, index

	} else {
		Charprint("red", "!!", "wrong int inputs")
	}

	return tmp
}
