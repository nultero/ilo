package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// 	data = data + "\n" + s
// 	// still needs stuff like for DateTimes and whatnot
// 	// based on fileType
// 	//
// 	os.WriteFile(path, []byte(data), 0644)
// 	done()
// }

// //  |||| |||   ||||| |||||
// //  |    |  |    |     |
// //  ||   |   |   |     |
// //  |    |   |   |     |
// //  |||| ||||  |||||   |

// func edit(b bus) {

// 	path := pathGlob(b.FileType)
// 	datum, _ := os.ReadFile(path)

// 	icon := b.PromptIcon
// 	data := string(datum)

// 	opts := strings.Split(data, "\n")
// 	editedLine := handleOptions(icon, opts)
// 	index := indexMatch(editedLine, opts)

// 	fmt.Println("string selected for edit: " + editedLine + "\nreplace: ")
// 	s := handleStringInput(icon)

// 	checks := []string{}
// 	if needsInfo(b.FileType) {
// 		checks = addInfo(icon)
// 	}

// 	if !isEmpty(checks) {
// 		s = fmt.Sprintf("%s @ %s %s", s, checks[0], checks[1])
// 	}

// 	opts[index] = s
// 	writeOut(path, opts, true)
// }

// // Lists contents of argpath.
// func list(path string, fileType string) {
// 	ls, _ := os.ReadFile(pathGlob(fileType))
// 	contents := string(ls)

// 	if fileType != "events" && fileType != "recurrents" {
// 		fmt.Println(contents)
// 	} else {
// 		cleanPrint(contents)
// 	}
// } // \-----------------------

// //  |||   |     |  |     |
// //  |  |  ||| |||   |   |
// //  |||   |  |  |    | |
// //  |  |  |     |     |

// func remove(b bus) {

// 	path := pathGlob(b.FileType)
// 	datum, _ := os.ReadFile(path)

// 	icon := b.PromptIcon
// 	data := string(datum)

// 	opts := strings.Split(data, "\n")
// 	removedLine := handleOptions(icon, opts)
// 	index := indexMatch(removedLine, opts)

// 	removed := make([]string, len(opts)-1)

// 	//writing common elements to altOpts, skipping index to rmv
// 	removed = append(opts[:index], opts[index+1:]...)
// 	writeOut(path, removed, true)
// }

// // Common to Edit() and Remove(). Grabs index for writing to the spot in opts[] arg.
// func indexMatch(strMatch string, opts []string) int {
// 	var index int
// 	for i := 0; i < len(opts); i++ {
// 		if opts[i] == strMatch {
// 			index = i
// 		}
// 	}
// 	return index
// }

// // Final string export for both Edit() and Remove(). Iterates
// // over their string arrays and adds newlines to zip the split seam
// // from the original pass-in that makes Opts[].
// func writeOut(path string, opts []string, printDone bool) {
// 	finalStr := ""
// 	for i := 0; i < len(opts); i++ {
// 		finalStr += opts[i] + "\n"
// 	}
// 	finalStr = finalStr[:len(finalStr)-1] // cuts off last newline
// 	os.WriteFile(path, []byte(finalStr), 0644)

// 	if printDone {
// 		done()
// 	}
// }

// //Prints hallelujah amen.
// func done() {
// 	fmt.Println("> Done!")
// }

// func addInfo(p string) []string {
// 	d := handleStringInput(p + " day of this event? ")
// 	m := handleOptions(p, getMonths())

// 	return []string{d, m}
// }

// func needsInfo(s string) bool {
// 	e := "events"
// 	r := "recurrents"

// 	if s == e || s == r {
// 		return true
// 	}

// 	return false
// }
