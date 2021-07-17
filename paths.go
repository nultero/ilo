package main

import (
	"os"
	"path/filepath"
)

func getHome() string {
	home, _ := os.UserHomeDir()
	return home
}

func HandleHomePath(p string) string {
	rebuiltPath := ""
	if p[0] == '~' {
		rebuiltPath += getHome() + p[1:]
	} else {
		return p
	}
	return rebuiltPath
}

func PathGlob(fileType string) string {
	switch fileType {

	case "check", "cache":
		fileType = "tailbox_cache"

	case "recurrent":
		fileType = "recurrent_reminders"

	case "todo":
		fileType = "todos"

	}

	h := HandleHomePath(PATH)
	p, _ := filepath.Abs(
		string(h + fileType + ".txt"),
	)
	return p
}

func HomeSlashPath() string {
	p := HandleHomePath(PATH)
	if p[len(p)-1] != '/' {
		p += "/"
	}
	return p
}

func ConfigPath() string {
	return HomeSlashPath() + "config.txt"
}
