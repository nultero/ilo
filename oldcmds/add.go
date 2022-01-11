package cmds

import (
	"bx/bxd"
	"bx/errs"
	"bx/fn"
	"os"

	"fmt"
)

//     |     |||   |||
//    | |    |  |  |  |
//   |||||   |  |  |  |
//  |     |  |||   |||

func add(b fn.Bus) {

	original, err := os.ReadFile(b.GetPath())
	if err != nil {
		errs.ThrowSys(err)
	}

	slice := sliceOffTxt(b.FileType)
	fmt.Println("what to add to", bxd.Emph(slice), "?")

	output := "\n"

	switch slice {
	case "events":
		fmt.Println("fmt for time-based checks:")
		fmt.Println(bxd.Blue("-> event"), "@ day month :")
		output += fn.GetInput() + " @ "
		fmt.Println("-> event @", bxd.Blue("day month"), ":")
		output += fn.GetInput()

	case "recurrents":
		fmt.Println("fmt for time-based checks:")
		fmt.Println(bxd.Blue("-> event"), "@ day :")
		output += fn.GetInput() + " @ "
		fmt.Println("-> event @", bxd.Blue("day"), ":")
		output += fn.GetInput()

	default:
		output += fn.GetInput()
	}

	output = string(original) + output
	writeOut(b.GetPath(), output)
}
