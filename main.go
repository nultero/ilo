package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
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

		// CheckReminders(now, month)

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
				args[i] = prompt(i)
			}
		}
	}

	switch args[0] { // final args are digested into eval
	case "add":
		Add(args[1], promptstr)
	case "edit":
		Edit(args[1], promptstr)
	case "list":
		List(args[1])
	case "remove":
		Remove(args[1])
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
func prompt(i int) string {

	tmp := ""
	msg := "options for "
	fmt.Println(">", strings.Repeat("-", 15))

	if i == 0 { // verb was missing
		promptstr := "> " + string(msg+"bx's verbs")
		tmp = verbsPrompt(promptstr)

	} else { // noun was missing
		promptstr := "> " + string(msg+"bx's nouns")
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

	return HandlesOpts(TinyArray(promptstr), opts)
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

	return HandlesOpts(TinyArray(promptstr), opts)
}

func HandlesOpts(promptstr, opts []string) string {

	prompt := promptui.Select{
		Label: promptstr[0],
		Items: opts,
	}

	_, tmp, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return tmp
}
