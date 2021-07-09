package main

import (
	"os"
)

func getHome() string {
	home, _ := os.UserHomeDir()
	return home
}

func HandleHomePath(confPath string) string {
	rebuiltPath := ""
	if confPath[0] == '~' {
		rebuiltPath += getHome() + confPath[1:]
	} else {
		rebuiltPath = confPath
	}
	return rebuiltPath
}

func PathGlob(path string, fileType string) string {
	switch fileType {

	case "recurrent":
		fileType = "recurrent_reminders"

	case "todo":
		fileType = "todos"

	}

	return string(path + fileType + ".txt")
}
