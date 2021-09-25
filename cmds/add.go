package cmds

import (
	"bx/bxd"
	"bx/fn"

	"fmt"
)

//     |     |||   |||
//    | |    |  |  |  |
//   |||||   |  |  |  |
//  |     |  |||   |||

func add(b fn.Bus) {

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

	writeOut(b.GetPath(), output)
}
