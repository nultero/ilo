package main

import (
	"fmt"
	"os"
	"time"
)

// Config is for prompts, controlling Charprint colors, and what/when types print
// function itself just inits / grabs those
func Config(confPath string) string {

	config, excptn := os.ReadFile(confPath)

	for excptn != nil { // files don't exist / deleted

		var prompt = "❯➤ " // default

		// user confirms dir  / not making sys changes w/out input
		confNotFound(confPath, prompt)
		writeDefaults(prompt, confPath)
		configsOutro(prompt)
		excptn = nil
		config, _ = os.ReadFile(confPath)

	}

	return string(config)
}

func confNotFound(confPath string, prompt string) {

	Charprint("red", "<bx>", "No config file found.")
	dfPath := "Default config path is: " + confPath
	Charprint("", "<bx>", dfPath)
	fmt.Println(" ")
	Charprint("", prompt, "Create new config? (y | n)")
	Charprint("", prompt, "(bx will just exit if not 'y')")

	if string(inpt()) == "y" {
		bxPath := Transformer("homedir")
		os.Mkdir(bxPath, 0755)
	} else {
		Charprint("", prompt, "sure thing, pardner")
		os.Exit(0)
	}
}

func inpt() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return input
}

func configsOutro(prompt string) {
	fmt.Println(prompt + " <bx> Folder created.")
	time.Sleep(700 * time.Millisecond)
	Charprint("", "", "<bx> Configured defaults.")
	time.Sleep(700 * time.Millisecond)
	Charprint("", "", "<bx> Feel free to change them by running {bx defaults}.")
	time.Sleep(700 * time.Millisecond)
	fmt.Println(" ")
}

func writeDefaults(prompt string, confPath string) {

	outbound := []byte("default_prompt = " + prompt +
		"\n" + "default_todo_symbol = ○")
	os.WriteFile(confPath, outbound, 0644)

	basePath := Transformer("homedir")

	files := []string{
		"one_time_reminders.txt",
		"recurrent_reminders.txt",
		"todos.txt",
		"aliases.txt"}

	for index, file := range files {
		frmt := fmt.Sprint(index+1) + " > Creating " + file + " in " + basePath
		Charprint("instant", "<bx>", frmt)
		fmt.Println(" ")
		os.WriteFile(
			(basePath + file), []byte(""), 0644)
	}
}
