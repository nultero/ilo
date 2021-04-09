package main

import (
	"flag"
	"fmt"
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
		fmt.Println("plop")
	}
}

func prompt(i int) string {

	tmp := ""
	if i == 0 { // verb was missing
		tmp = verbsPrompt()

	} else { // noun was missing
		tmp = nounsPrompt()
	}
	return tmp
}

func verbsPrompt() string {

	tmp := ""

	return tmp
}

func nounsPrompt() string {

	tmp := ""

	return tmp
}
