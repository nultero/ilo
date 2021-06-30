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
