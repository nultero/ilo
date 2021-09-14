package bxd

import (
	"bx/errs"
	"os"
)

// For helping validate several different kinds of main paths.
func CheckPath(path string) string {

	if path[0] == '~' {
		home := getHomeDir()
		path = home + path[1:]
	} // put in else for other homepath

	return path
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		errs.ThrowSys(err)
	}
	return home
}
