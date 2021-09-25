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

	for {
		fmt.Println(prompt)
		fmt.Print(bxd.Emph("[ Y / n ] "))
		confirm := strings.ToLower(fn.GetInput())

		if len(confirm) == 0 || strings.Contains(confirm, "y") {
			return true
		}

		if strings.Contains(confirm, "n") {
			return false
		}

		fmt.Println(bxd.Emph("invalid input"))
	}
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
		}

		done = isDone(remove)
	}

	writeOut(b.GetPath(), cleanLines(lines))
}

func edit(b fn.Bus) {
	data := getData(b.GetPath())
	lines := strings.Split(data, "\n")

	s := "which line to edit?"
	i := selector(s, lines)

	fmt.Println(i, lines[i])

}

func remove(b fn.Bus) {
	// data := getData(b.GetPath())
	// lines := strings.Split(data, "\n")

	// s := "which line to remove?"
	// i := selector(s, lines)

	// writeOut(b.GetPath(), cleanLines(filteredLines))
	fmt.Println("done!")
}
