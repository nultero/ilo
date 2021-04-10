package main

import (
	"fmt"
	"os"
	"strings"
)

// outline --
//   .txt grabber
//   .txt reader
//      add()
//      edit()
//      list()
//      remove()
//
//
//    ||||| |  | |||||
//      |    ||    |
//      |    ||    |
// ||   |   |  |   |

// Canonicals in 'argTransformers.go' don't append '.txt'
// onto files, so this private func makes the end filenames
// slightly more legible / transparent for finding them in
// file exploration. (and also makes file grab easier)
func txtFileNamer(location string) string {
	switch location {

	case "onetime":
		location = "one_time_reminders"

	case "alias":
		location = "aliases"

	case "event":
	case "idea":
	case "todo":
		location += "s"

	case "recurrent":
		location = "recurrent_reminders"

	}
	location += ".txt"
	return location
}

//     |     |||   |||
//    | |    |  |  |  |
//   |||||   |  |  |  |
//  |     |  |||   |||

// takes & adds a str input onto a given filetype
func Add(typeOf string, promptstr string) {
	typeOf = txtFileNamer(typeOf)
	readPath := ArgCleaner("homedir") + typeOf
	conts, _ := os.ReadFile(readPath)
	typeOf = string(conts)
	fmt.Println("Thing to add:")
	dat := HandleOptionsInput(promptstr)
	typeOf = typeOf + "\n" + dat
	// still needs stuff like for DateTimes and whatnot
	//
	//
	os.WriteFile(readPath, []byte(typeOf), 0644)
}

//  ||   ||| |||| |||
//  ||    |   |    |
//  ||    |    |   |
//  |||| ||| ||||  |

// Lists contents of argpath. What it says on the tin.
func List(typeOf string) {
	typeOf = txtFileNamer(typeOf)
	fmt.Println(typeOf + " > test")
	readPath := ArgCleaner("homedir") + typeOf
	fmt.Println(readPath + " > test")
	conts, _ := os.ReadFile(readPath)
	fmt.Println(string(conts))
}

func Test() {
	TermClear()
	fmt.Print("Input? : ")
	// i := HandleOptionsInput("> ")
	fmt.Println("red", "", strings.Repeat("Jackson", 90))
}

func TermClear() {
	fmt.Print("\033[H\033[2J")
}
