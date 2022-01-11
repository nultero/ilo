package cmds

import (
	"bx/errs"
	"bx/fn"
	"fmt"
	"os"
)

// //  ||   ||| |||| |||
// //  ||    |   |    |
// //  ||    |    |   |
// //  |||| ||| ||||  |

// // \-------------------------

func ls(b fn.Bus) {

	p := b.Path + b.FileType
	f, err := os.ReadFile(p)
	if err != nil {
		errs.ThrowSys(err)
	}

	fmt.Println(string(f))
}
