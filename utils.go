package main

import (
	"fmt"
)

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

func Test() {
	fmt.Println(passesDayCheck(
		30, 3, 3, "Jan",
	))
}

func TermClear() {
	fmt.Print("\033[H\033[2J")
}
