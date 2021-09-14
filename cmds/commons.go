package cmds

import (
	"bx/bxd"
	"bx/errs"
	"fmt"
	"os"
)

func WriteOut(path, s string) {

	b, err := os.ReadFile(path)
	if err != nil {
		errs.ThrowSys(err)
	}

	finalStr := []byte(string(b) + s)

	r := os.WriteFile(path, finalStr, 0644)
	if r != nil {
		errs.ThrowSys(r)
	}

	fmt.Println(bxd.Blue("<✓>"))
}
