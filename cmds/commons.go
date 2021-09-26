package cmds

import (
	"bx/bxd"
	"bx/errs"
	"bx/fn"
	"fmt"
	"os"
	"strings"

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

// True == yes.
func getYesOrNo(s ...string) bool {
	for {

		fmt.Print("\n")
		if len(s) > 0 {
			fmt.Println(s[0])
		}

		fmt.Print("> ", bxd.Emph("[ Y / n ] "))
		confirm := strings.ToLower(fn.GetInput())

		if len(confirm) == 0 || strings.Contains(confirm, "y") {
			return true
		}

		if strings.Contains(confirm, "n") {
			return false
		}

		fmt.Println("\n" + bxd.Emph("invalid input"))
	}
}
