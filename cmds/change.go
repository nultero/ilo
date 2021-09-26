package cmds

import (
	"bx/bxd"
	"bx/errs"
	"bx/fn"
	"fmt"
	"os"
	"strings"
)

func getData(path string) string {
	rawdata, err := os.ReadFile(path)
	if err != nil {
		errs.ThrowSys(err)
	}

	return string(rawdata)
}

func isDone(remove bool) bool {
	prompt := "done "
	if remove {
		prompt += "tagging for removal?"
	} else {
		prompt += "editing?"
	}

	return getYesOrNo(prompt)
}

func change(b fn.Bus, remove bool) {
	data := getData(b.GetPath())
	lines := strings.Split(data, "\n")

	var done = false
	for !done {
		var s string = "which line to "
		if remove {
			s += "remove?"
		} else {
			s += "edit?"
		}

		i := selector(s, lines)
		if remove {
			filteredLines := []string{}
			for li, s := range lines {
				if li != i {
					filteredLines = append(filteredLines, s)
				}
			}
			lines = filteredLines

		} else { // !remove == edit
			fmt.Println("editing '" + bxd.Emph(lines[i]) + "'")
			lines[i] = fn.GetInput()
		}

		done = isDone(remove)
	}

	if confirmChanges() {
		writeOut(b.GetPath(), cleanLines(lines))
	}

}

func confirmChanges() bool {
	return getYesOrNo("> confirm changes?")
}
