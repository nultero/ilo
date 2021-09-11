package main

import (
	"bx/bx"
	"bx/cmds"
	"bx/conf"
	"bx/fn"
	"os"
)

// tview instead of promptui

const PATH = "~/.tailbox/"

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	p := bx.CheckPath(PATH)
	icon, config := conf.Ok(p)

	// still have to double check config
	// but mostly cleaner than it was before

	// the ux for adding things is extremely unclear
	// color the parts being editing for focus etc.

	args := os.Args[1:]

	if len(args) == 0 {
		bx.RunReminders(p, icon, config)

	} else {
		bus := fn.DefaultBus(icon, p)
		cmds.ParseArgs(args, bus)
	}
}
