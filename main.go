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
// * RETEST CONFIG
//
// change default directory
// -dir flag for pwd on bx's $home
// -h flag for forgetting

func main() {

	flag.Parse()
	configs := Config(Transformer("homedir conf"))
	prompt := string(Parser(configs, "=", "default_prompt") + " ")

	if len(flag.Args()) == 0 {

		currentDate := time.Now().Format("02 Mon")
		month := time.Now().Month().String()[0:3]
		fmt.Println(prompt, month, currentDate)

		// Reminders()

	} else {

		for i := range flag.Args() {
			flag.Args()[i] = Transformer(flag.Args()[i])
			// pass clean args to priority function
			Charprint("blue", prompt, flag.Args()[i])
		}

	}

}
