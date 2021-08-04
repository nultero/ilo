package main

import (
	"bx/bx"
	"bx/conf"
	"os"
)

const PATH = "~/.nultero/tailbox/"

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	p := bx.BxPath(PATH) // ^ verifies the above const, expands, etc.
	icon, config := conf.Ok(p)

	// still have to double check config
	// but mostly cleaner than it was before

	args := os.Args[1:]

	if len(args) == 0 {
		bx.RunReminders()

	} else {

		// b := bx.DefaultBus(icon, p)

		// fmt.Println(b)

	}
}
