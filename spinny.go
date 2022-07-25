package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var spinChars = []rune{
	'⡇', '⠏', '⠛', '⠹', '⢸', '⠼', '⠶', '⠧',
}

func clearLine() {
	fmt.Print("\r")
	fmt.Print(strings.Repeat(" ", 50))
	fmt.Print("\r")
}

func spinny(idx int) {
	clearLine()
	fmt.Print(string(spinChars[idx]), " Running ")
}

func main() {
	args := os.Args[1:]
	app := "bash"

	quit := make(chan struct{})

	go func() {
		cmd := exec.Command(app, args...)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		quit <- struct{}{}
	}()

	idx := 0 // spinChars idx

spinloop:
	for {
		select {
		case <-quit:
			clearLine()
			fmt.Print("\n")
			break spinloop

		case <-time.After(100 * time.Millisecond):
			spinny(idx)
			idx++
			if idx == len(spinChars) {
				idx = 0
			}
		}
	}
}
