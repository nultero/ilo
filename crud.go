package main

import (
	"fmt"
	"os"
)

func crudBus() {

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

//  ||   ||| |||| |||
//  ||    |   |    |
//  ||    |    |   |
//  |||| ||| ||||  |

// \-------------------------

// Lists contents of argpath. What it says on the tin.
func List(typeOf string) {
	fmt.Println(contents(typeOf))
} // \----------------------- LIST() basically done lol
