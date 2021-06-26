package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// bx's PATH should be, for instance:
const PATH = "~/.nultero/tailbox/"

const (
	promptIcon = "promptIcon"
	fileType   = "fileType"
	funct      = "func"
)

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	flag.Parse()
	logicBus := map[string]string{}
	config := Configure(PATH)
	logicBus[promptIcon] = string(Parser(config, "=", "default_prompt") + " ")

	if len(flag.Args()) == 0 { // zero args calls checks, prints reminders if close, and the date
		//						  and if already called today, prints cached info

		now := time.Now()
		todaysFormatDate := now.Format("02 Mon")
		month := now.Month().String()[0:3]
		fmt.Println(logicBus[promptIcon], month, todaysFormatDate)

	} else { // if bx is called with args, it will parse args instead of doing its checks

		for i := 0; i < len(flag.Args()); i++ {
			vals := validateArg(i, flag.Args()[i])
			logicBus[vals[0]] = vals[1]
		}

		if len(logicBus) == 2 { // in case of only CRUD arg, prompt is also in bus
			logicBus[fileType] = HandlesPrompts(logicBus[promptIcon], "args")
		}

		_eval(logicBus[funct], logicBus)
	}
}

func _eval(function string, bus map[string]string) {
	switch function {
	case "add":

	case "edit":

	case "list":

	case "remove":
		for i := range bus {
			fmt.Println(i, ": ", bus[i])
		}

	case "test", "home", "config", "help":
		Test()
	}
}

func _invalidateArg(rg string, i string) {
	fmt.Printf("'%v' is an invalid %v argument to bx", rg, i)
	os.Exit(0)
}

func validateArg(index int, rg string) []string {

	var vals []string

	// digest to see if help or home or configs or whatever
	// AND bx_defaults argument for changing stuff

	if index == 0 {
		fn := "func"
		switch rg {
		case "add", "edit", "list", "remove":
			vals = append(vals, fn, rg)
		default:
			_invalidateArg(rg, "1st")
		}

	} else if index == 1 {
		it := "item"
		switch rg {
		case "cron", "event", "idea", "recurrent", "todo", "wishlist":
			vals = append(vals, it, rg)
		default:
			_invalidateArg(rg, "2nd")
		}

	} else if index >= 2 {
		_invalidateArg(rg, "3rd")
	}
	return vals
}
