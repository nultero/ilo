package main

import (
	"bx/bxd"
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

	p := bxd.CheckPath(PATH)
	config, err := conf.Ok(p)
	if err != nil {
		conf.Fix(p)
		return
	}

	icon, _ := bxd.ParseString(config, "icon")
	icon += " "

	args := os.Args[1:]

	if len(args) == 0 {
		bxd.RunReminders(p, icon, config)

	} else {
		bus := fn.DefaultBus(icon, p)
		cmds.ParseArgs(args, bus)
	}
}
