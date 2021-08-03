package main

import (
	"bx/bx"
	"fmt"
	"os"
)

// bx's PATH should be, for instance:
const PATH = "~/.nultero/tailbox/"

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	p := bx.CheckPath(PATH) // ^ verifies the above const, expands, etc.
	icon, conf := config(p)

	args := os.Args[1:]
	// p, _ := filepath.Abs(handleHomePath(PATH))
	// config := configure()
	// icon := parser(config, "promptIcon") + " "

	if len(args) == 0 {
		bx.RunReminders()

	} else {

		b := bx.DefaultBus(icon, p)

		fmt.Println(b)

	}
}
