package main

import (
	"fmt"
	"os"
	"time"
)

// Config is for prompts, controlling fmt.Println colors, and what/when types print
// function itself just inits / grabs those
func ConfigureStuff(confPath string) string {

	config, excptn := os.ReadFile(confPath)

	for excptn != nil { // files don't exist / deleted

		prompt := "❯➤ " // default

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

	fmt.Println("<bx>", "No config file found.")
	fmt.Println("<bx>", "Default config path is: "+confPath)
	fmt.Printf(" %s Create new config? (y | n)"+"\t\n(bx will just exit if not 'y') : ", prompt)

	if grabInput() == "y" {
		//bx's homedir in config or wherever
		os.Mkdir(ArgCleaner("homedir"), 0755)

	} else { // this is if manually exited
		fmt.Println(prompt, "sure thing, pardner")
		os.Exit(0)
	}
}

func grabInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return input
}

func configsOutro(prompt string) {

	printsWhenDone := []string{
		"Folder created.",
		"Configured defaults.", //           color the command?
		"Feel free to change them by running {bx defaults}."}

	for confirmString := range printsWhenDone {
		fmt.Println(prompt, printsWhenDone[confirmString])
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println(" ") //buffer
}

func writeDefaults(prompt string, confPath string) {

	outbound := []byte(
		"default_todo_symbol = ○" +
			"\n" + "default_prompt = " + prompt)

	os.WriteFile(confPath, outbound, 0644)

	basePath := ArgCleaner("homedir")

	files := []string{
		"one_time_reminders.txt",
		"recurrent_reminders.txt",
		"todos.txt",
		"aliases.txt"}

	for index, file := range files {
		frmt := fmt.Sprint(index+1) + " > Creating " + file + " in " + basePath
		fmt.Println("<bx>", frmt)
		fmt.Println(" ")
		os.WriteFile(
			(basePath + file), []byte(""), 0644)
	}
}
