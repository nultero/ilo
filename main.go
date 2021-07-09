package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// bx's PATH should be, for instance:
const PATH = "~/.nultero/tailbox/"

// I should put the PATH in some kind of init script

const (
	promptIcon = "promptIcon"
	fileType   = "fileType"
	funct      = "func"
	bxPath     = "bxPath"
)

//  ||
//  |||||  ||  ||
//  || ||    ||
//  |||||  ||  ||

func main() {

	flag.Parse()
	logicBus := map[string]string{}
	path := HandleHomePath(PATH)
	config := Configure(path)
	logicBus[promptIcon] = Parser(config, promptIcon) + " "

	if len(flag.Args()) == 0 { // zero args calls checks, prints reminders if within config's threshold, and current date
		//						  and if already called today, prints cached info

		now := time.Now()
		todaysFormatDate := now.Format("02 Mon")
		month := now.Month().String()[0:3]
		fmt.Println(logicBus[promptIcon], month, todaysFormatDate)

		//checks
		//checks
		//checks

	} else { // if bx is called with args, it will parse args instead of doing its checks

		logicBus[bxPath] = path

		for i := 0; i < len(flag.Args()); i++ {
			vals := validateArg(i, flag.Args()[i])
			logicBus[vals[0]] = vals[1]
		}

		if len(logicBus) == 3 { // in case of only CRUD arg, prompt is also in bus
			promptStr := fmt.Sprintf("%s '%s' needs an argument", logicBus[promptIcon], logicBus[funct])
			logicBus[fileType] = HandleArguments(promptStr)
		}

		_eval(logicBus[funct], logicBus)
	}
}

func _eval(function string, bus map[string]string) {
	switch function {
	case "add":
		Add(bus)

	case "edit":
		Edit(bus)

	case "list":
		List(bus[bxPath], bus[fileType])

	case "remove":
		Remove(bus)

	case "test", "home", "config", "help":
		Test()
	}
}

func _invalidateArg(rg string, i string) {
	fmt.Printf("'%v' is an invalid %v argument to bx \n", rg, i)
	os.Exit(0)
}

func validateArg(index int, rg string) []string {

	var vals []string

	// digest to see if help or home or configs or whatever
	// AND bx_defaults argument for changing stuff

	if index == 0 {
		fn := funct
		switch rg {
		case "add", "edit", "list", "remove":
			vals = append(vals, fn, rg)
		default:
			_invalidateArg(rg, "1st")
		}

	} else if index == 1 {
		ft := "fileType"
		switch rg {
		case "todo":
			vals = append(vals, ft, "todos")

		case "cron", "event", "idea", "recurrent", "wishlist":
			vals = append(vals, ft, rg)

		default:
			_invalidateArg(rg, "2nd")
		}

	} else if index >= 2 {
		_invalidateArg(rg, "3rd")
	}
	return vals
}
