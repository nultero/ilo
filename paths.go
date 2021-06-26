package main

import (
	"fmt"
	"os"
)

func getHome() string {
	home, _ := os.UserHomeDir()
	return home
}

func HandleConfPath(confPath string) string {
	rebuiltPath := ""
	if confPath[0] == '~' {
		rebuiltPath += getHome() + confPath[1:]
	}
	fmt.Printf("rebuilt path = %s \n", rebuiltPath)
	return rebuiltPath
}
