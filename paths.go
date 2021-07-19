package main

import (
	"os"
	"path/filepath"
)

func getHome() string {
	home, _ := os.UserHomeDir()
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

func pathGlob(fileType string) string {
	switch fileType {

	case "check", "cache":
		fileType = "tailbox_cache"

	case "recurrents":
		fileType = "recurrent_reminders"

	case "todo":
		fileType = "todos"

	}

	h := handleHomePath(PATH)
	p, _ := filepath.Abs(
		string(h + fileType + ".txt"),
	)
	return p
}

func homeSlashPath() string {
	p := handleHomePath(PATH)
	if p[len(p)-1] != '/' {
		p += "/"
	}
	return p
}

func configPath() string {
	return homeSlashPath() + "config.txt"
}
