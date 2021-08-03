package errs

import (
	"fmt"
	"os"
)

func ThrowX(r error) {
	s := fmt.Sprintf("%s %s", redError(), r)
	fmt.Println(s)
	os.Exit(1)
}

func throwDuplArgError(this, prevFound string) {
	fmt.Printf("! >> '%s' found, but already passed '%s' as argument \n", this, prevFound)
	fmt.Println(redError(), "cannot have two of the same type of argument")
	os.Exit(1)
}

func redError() string {
	return "\033[31;1;4merror:\033[0m"
}

func throwErr(r error) {
	fmt.Println(redError(), r)
	os.Exit(1)
}
