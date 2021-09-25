package cmds

import (
	"bx/bxd"
	"bx/errs"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func selector(prompt string, s []string) int {

	q := promptui.Select{
		Label: prompt,
		Items: s,
	}

	_, v, err := q.Run()

	if err != nil {
		v = fmt.Sprintf("prompt failure: %v", err)
		errs.ThrowQuiet(v)
	}

	idx := 0

	for i, ln := range s {
		if v == ln {
			idx = i
		}
	}

	return idx

}

func writeOut(path, s string) {
	r := os.WriteFile(path, []byte(s), 0644)
	if r != nil {
		errs.ThrowSys(r)
	}

	fmt.Println(bxd.Blue("<✓>"))
}

func cleanLines(lines []string) string {
	s := ""
	for _, l := range lines {
		s += l + "\n"
	}

	return s[:len(s)-1]
}

func sliceOffTxt(s string) string {
	return s[:len(s)-4]
}
