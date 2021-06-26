package main

import (
	"bufio"
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

//Tacks home directory onto paths to .txts.
func TxtPath(typeOf string) string {
	return ArgCleaner("homedir") + txtFileNamer(typeOf)
}

// Master function for .txt controls.
// End result is a full string read from a file in bx's $path.
func contents(typeOf string) string {
	conts, _ := os.ReadFile(TxtPath(typeOf))
	return string(conts)
}

//     |     |||   |||
//    | |    |  |  |  |
//   |||||   |  |  |  |
//  |     |  |||   |||

// takes & adds a str input onto a given filetype
func Add(typeOf string, promptstr string) {
	dat := contents(typeOf)
	fmt.Println("Thing to add:")
	addLine := HandleStringInput(promptstr)
	dat = dat + "\n" + addLine
	// still needs stuff like for DateTimes and whatnot
	// based on typeOf
	//
	os.WriteFile(TxtPath(typeOf), []byte(dat), 0644)
	Done()
}

//  |||| |||   ||||| |||||
//  |    |  |    |     |
//  ||   |   |   |     |
//  |    |   |   |     |
//  |||| ||||  |||||   |

// func Edit(typeOf string, promptstr string) {
// 	dat := contents(typeOf)
// 	opts := strings.Split(dat, "\n")
// 	// toEdit := HandlesOpts(TinyArray("entry to edit"), opts)
// 	index := indexMatch(toEdit, opts)

// 	fmt.Println("string selected for edit: " + toEdit + "\nreplace: ")
// 	replacer := HandleStringInput(promptstr)
// 	opts[index] = replacer
// 	writeOut(typeOf, opts)
// }

//  ||   ||| |||| |||
//  ||    |   |    |
//  ||    |    |   |
//  |||| ||| ||||  |

// \-------------------------

// Lists contents of argpath. What it says on the tin.
func List(typeOf string) {
	fmt.Println(contents(typeOf))
} // \----------------------- LIST() basically done lol

func Test() {
	fmt.Println(passesDayCheck(
		30, 3, 3, "Jan",
	))
}

func TermClear() {
	fmt.Print("\033[H\033[2J")
}

//  |||   |     |  |     |
//  |  |  ||| |||   |   |
//  |||   |  |  |    | |
//  |  |  |     |     |

// func Remove(typeOf string) {
// 	dat := contents(typeOf)
// 	opts := strings.Split(dat, "\n")
// 	toEdit := HandlesOpts(TinyArray("entry to remove"), opts)
// 	index := indexMatch(toEdit, opts)

// 	altOpts := make([]string, len(opts)-1)

// 	//writing common elements to altOpts, skipping index to rmv
// 	altOpts = append(opts[:index], opts[index+1:]...)
// 	writeOut(typeOf, altOpts)
// }

// Common to Edit() and Remove(). Grabs index for writing to the spot in opts[] arg.
func indexMatch(strMatch string, opts []string) int {
	var index int
	for i := 0; i < len(opts); i++ {
		if opts[i] == strMatch {
			index = i
		}
	}
	return index
}

// Final string export for both Edit() and Remove(). Iterates
// over their string arrays and adds newlines to zip the split seam
// from the original pass-in that makes Opts[].
func writeOut(typeOf string, opts []string) {
	finalStr := ""
	for i := 0; i < len(opts); i++ {
		finalStr += opts[i] + "\n"
	}
	finalStr = finalStr[:len(finalStr)-1] // cuts off last newline
	os.WriteFile(TxtPath(typeOf), []byte(finalStr), 0644)
	Done()
}

// Exports a string that is unchecked against any args.
// Pass to a separate function to map cleaner args in passthrough.
func HandleStringInput(promptstr string) string { // only takes/rds string
	fmt.Print(promptstr + " ")
	nex, _, e := bufio.NewReader(os.Stdin).ReadLine()
	for e != nil {
		fmt.Println(e)
	}
	return string(nex)
}

// What it says on the tin. Takes a string and puts it into an array of size(1).
// For args passed to the Handler exported in Main, which is itself called from
// funcs exported in /Utils, which are then called in Main. Roundabout.
func TinyArray(onlyMember string) []string {
	tiny := make([]string, 1)
	tiny[0] = onlyMember
	return tiny
}

//Prints hallelujah amen.
func Done() {
	fmt.Println("> Done!")
}
