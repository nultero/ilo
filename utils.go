package main

import (
	"fmt"
	"os"
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
func txtCaser(location string) string {
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
func Add(typeOf string) {
	typeOf = txtCaser(typeOf)
	readPath := ArgCleaner("homedir") + typeOf
	conts, _ := os.ReadFile(readPath)
	typeOf = string(conts)
	fmt.Println(typeOf) // read to byte, restring and all still left
}
