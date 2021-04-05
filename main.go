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

	//
	// pass the whole arg list to the cleaner function instead of this iterating
	//

	if len(flag.Args()) == 0 { // bx will only check stuff if called passively

		now := time.Now()
		formattedTime := now.Format("02 Mon")
		month := TrimMonth(now.Month().String())
		fmt.Println(prompt, month, formattedTime)

		CheckReminders(now)

	} else {

		for arg := range flag.Args() {
			flag.Args()[arg] = ArgCleaner(flag.Args()[arg])
			// fmt.Println(arg)
		}
		ArgumentPrioritizer(flag.Args())
		// fmt.Println(flag.Args())

	}

}
