package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

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
		eval(sortedArgs, prompt)
	}
}

func eval(args []string, promptstr string) {

	if args[0] != "test" && args[0] != "help" {
		for i := 0; i < len(args); i++ { // prompts
			if args[i] == "" {
				args[i] = prompt(i, promptstr)
			}
		}
	}

	switch args[0] { // final args are digested into eval
	case "add":
		Add(args[1], promptstr)
	case "edit":
		fmt.Println("editing thing")
	case "list":
		List(args[1])
	case "remove":
		fmt.Println("removing thing")
	case "test":
		Test()
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
func prompt(i int, promptstr string) string {

	tmp := ""
	msg := "options for "
	delimitr := strings.Repeat("-", 15)
	if i == 0 { // verb was missing
		fmt.Println(">! ", string(msg+"bx's verbs"))
		fmt.Println(">", delimitr)
		tmp = verbsPrompt(promptstr)

	} else { // noun was missing
		fmt.Println(">! ", string(msg+"bx's nouns"))
		fmt.Println(">", delimitr)
		tmp = nounsPrompt(promptstr)
	}

	return tmp
}

func verbsPrompt(promptstr string) string {
	opts := []string{
		"add",
		"edit",
		"list",
		"help",
		"test",
		"remove"}
	tmp := ""
	printsOptions(opts)
	tmp = HandleOptionsInput(promptstr)
	return optionsChecks(tmp, opts)
}

func nounsPrompt(promptstr string) string {
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
	tmp = HandleOptionsInput(promptstr)
	return optionsChecks(tmp, opts)
}

func printsOptions(opts []string) {
	for i := range opts {
		fmt.Println(">", string(
			fmt.Sprint(i+1)+ //index
				". "+ // "1. remove" etc
				string(opts[i])))
	}
}

// Exports a string that is unchecked against any args.
// Pass to a separate function to map cleaner args in passthrough.
func HandleOptionsInput(promptstr string) string { // only takes/rds string

	fmt.Print(promptstr + " ")
	nex, e := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print(nex)
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
		fmt.Println("!!", "wrong int inputs")
	}

	return tmp
}
