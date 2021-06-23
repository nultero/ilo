package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// bx's PATH should be, for instance:
const PATH = "~/.nultero/tailbox/"

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	flag.Parse()

	if len(flag.Args()) == 0 { // zero args just prints reminders and date
		//						  and if already called today, prints cached info

		// configs := map[string]string{}

		now := time.Now()
		todaysFormatDate := now.Format("02 Mon")
		month := now.Month().String()[0:3]
		fmt.Println(month, todaysFormatDate)

	} else { // if bx is called with args, will parse args

		// argsMap := map[string]string{}

		for i := 0; i < len(flag.Args()); i++ {
			vals := validateArg(i, flag.Args()[i])
			fmt.Println(vals)
		}

	}
}

func invalidArg(rg string, i string) {
	fmt.Printf("'%v' is an invalid %v argument to bx", rg, i)
	os.Exit(0)
}

func validateArg(index int, rg string) []string {

	var vals []string

	if index == 0 {
		fn := "func"
		switch rg {
		case "add", "edit", "list", "remove":
			vals = append(vals, fn, rg)
		default:
			invalidArg(rg, "1st")
		}
	} else if index == 1 {
		switch rg {
		case "add", "edit", "list", "remove":
			vals = append(vals, rg)
		}

	} else if index >= 2 {

	}
	return vals
}
