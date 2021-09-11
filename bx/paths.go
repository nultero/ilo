package bx

import (
	"bx/errs"
	"os"
)

// For helping validate several different kinds of paths.
func CheckPath(path string) string {

	if path[0] == '~' {
		home := getHomeDir()
		path = home + path[1:]
	} // put in else for other

	return path
}

func getHomeDir() string {
	home, err := os.UserHomeDir()

	if err != nil {
		errs.ThrowSys(err)
	}

	return home
}

func glob(fileType string) string {
	switch fileType {
	case "cache":
		fileType = "tailbox_cache"

	case "recurrents":
		fileType = "recurrent_reminders"

	case "todo":
		fileType = "todos"

	}

	p := fileType + ".txt"
	return p
}

// func homeSlashPath(path string) string {
// 	p := handleHomePath(path)
// 	if p[len(p)-1] != '/' {
// 		p += "/"
// 	}
// 	return p
// }

// func configPath() string {
// 	return homeSlashPath() + "config.txt"
// }
