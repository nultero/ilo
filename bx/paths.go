package bx

import (
	"bx/errs"
	"os"
	"path/filepath"
)

func CheckPath(path string) string {
	return path
}

func getHome() string {
	home, err := os.UserHomeDir()

	if err != nil {
		errs.ThrowX(err)
	}

	return home
}

func handleHomePath(p string) string {

	rebuiltPath := ""
	if p[0] == '~' {
		rebuiltPath += getHome() + p[1:]
	} else {
		return p
	}
	return rebuiltPath
}

func pathGlob(fileType, path string) string {
	switch fileType {

	case "check", "cache":
		fileType = "tailbox_cache"

	case "recurrents":
		fileType = "recurrent_reminders"

	case "todo":
		fileType = "todos"

	}

	h := handleHomePath(path)
	p, _ := filepath.Abs(
		string(h + fileType + ".txt"),
	)
	return p
}

func homeSlashPath(path string) string {
	p := handleHomePath(path)
	if p[len(p)-1] != '/' {
		p += "/"
	}
	return p
}

// func configPath() string {
// 	return homeSlashPath() + "config.txt"
// }
