package main

import (
	"fmt"
	"math"
	"time"
)

// Charprint does the fancy typewriter effect w/
// modifier str, which makes it easy to set up as a string read
// from config, or to pass in and specify output speeds or colors
//
// probably still needs term width formatter
//
func Charprint(modifier string, prompt string, output string) {

	subString := ""
	wrapString := "%v %v"
	speed := 35 //sweet spot for printspeed imo, but is a configable var

	switch modifier { // maybe use a split and map vars from there?
	case "red":
		wrapString = "\033[1;31m%v %v\033[0m"
	case "blue":
		wrapString = "\033[1;36m%v %v\033[0m"

	// 	Magenta "\033[1;35m%s\033[0m")
	// 	Teal    "\033[1;36m%s\033[0m")
	//
	//
	case "kinda slow":
		speed = 45

	case "instant":
		speed = 0

	case "nil":
	}
	wrapString = "\r" + wrapString

	// print speeds vary depending on line length
	// the gimmick gets annoying if overused
	// might only use for special stuff to make Reminders() pop more
	// definitely helps to stop the eyeglaze from seeing the usual
	// might even want some fx like lolcat or smthn

	//logscale print reduction?
	normalizedSpeed := math.Log(float64(len(output)))

	speed -= int((normalizedSpeed - 3) * 5)
	fmt.Printf("\n speed is: %v \n", speed)
	if speed <= 0 {
		fmt.Printf(wrapString, prompt, output)

	} else {

		for i := range output {
			subString += string(output[i])
			fmt.Printf(wrapString, prompt, subString)
			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
		fmt.Println(" ") //buf
	}

}
