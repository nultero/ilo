package main

import (
	"fmt"
	"os"
	"strings"
)

//     |     |||   |||
//    | |    |  |  |  |
//   |||||   |  |  |  |
//  |     |  |||   |||

// takes & adds a str input onto a given filetype
func add(b bus) {

	path := pathGlob(b.FileType)
	datum, _ := os.ReadFile(path)

	data := string(datum)

	fmt.Println("Adding to '" + b.FileType + "':")
	addLine := handleStringInput(b.PromptIcon)
	data = data + "\n" + addLine
	// still needs stuff like for DateTimes and whatnot
	// based on fileType
	//
	os.WriteFile(path, []byte(data), 0644)
	done()
}

//  |||| |||   ||||| |||||
//  |    |  |    |     |
//  ||   |   |   |     |
//  |    |   |   |     |
//  |||| ||||  |||||   |

func edit(b bus) {

	path := pathGlob(b.FileType)
	datum, _ := os.ReadFile(path)

	icon := b.PromptIcon
	data := string(datum)

	opts := strings.Split(data, "\n")
	editedLine := handleOptions(icon, opts)
	index := indexMatch(editedLine, opts)

	fmt.Println("string selected for edit: " + editedLine + "\nreplace: ")
	opts[index] = handleStringInput(icon)

	writeOut(path, opts)
}

//  ||   ||| |||| |||
//  ||    |   |    |
//  ||    |    |   |
//  |||| ||| ||||  |

// \-------------------------

// Lists contents of argpath.
func list(path string, fileType string) {
	ls, _ := os.ReadFile(pathGlob(fileType))
	contents := string(ls)
	fmt.Println(contents)
} // \-----------------------

//  |||   |     |  |     |
//  |  |  ||| |||   |   |
//  |||   |  |  |    | |
//  |  |  |     |     |

func remove(b bus) {

	path := pathGlob(b.FileType)
	datum, _ := os.ReadFile(path)

	icon := b.PromptIcon
	data := string(datum)

	opts := strings.Split(data, "\n")
	removedLine := handleOptions(icon, opts)
	index := indexMatch(removedLine, opts)

	removed := make([]string, len(opts)-1)

	//writing common elements to altOpts, skipping index to rmv
	removed = append(opts[:index], opts[index+1:]...)
	writeOut(path, removed)
}

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
func writeOut(path string, opts []string) {
	finalStr := ""
	for i := 0; i < len(opts); i++ {
		finalStr += opts[i] + "\n"
	}
	finalStr = finalStr[:len(finalStr)-1] // cuts off last newline
	os.WriteFile(path, []byte(finalStr), 0644)
	done()
}

//Prints hallelujah amen.
func done() {
	fmt.Println("> Done!")
}
