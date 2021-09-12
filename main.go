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
	config, err := conf.Ok(p)
	if err != nil {
		conf.Fix(p)
		return
	}

	icon, _ := bx.ParseString(config, "icon")

	args := os.Args[1:]

	if len(args) == 0 {
		bx.RunReminders(p, icon, config)

	} else {
		bus := fn.DefaultBus(icon, p)
		cmds.ParseArgs(args, bus)
	}
}
