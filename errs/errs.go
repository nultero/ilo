package errs

import (
	"fmt"
	"os"
)

func quit(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func ThrowX(err error, desc string) {
	s := fmt.Sprintf("%s %s \n >> %s", redError(), err, desc)
	quit(s)
}

func ThrowSys(err error) {
	s := fmt.Sprintf("%s %s", redError(), err)
	quit(s)
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
