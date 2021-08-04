package conf

import (
	"bx/bx"
	"bx/errs"
	"os"
)

func Ok(path string) (string, string) {

	p := path + "config.txt"
	config, err := os.ReadFile(p)

	if err != nil {
		errs.ThrowSys(err)
	}

	p = string(config)

	icon, parseErr := bx.ParseString(p, "promptIcon")

	if parseErr != nil {
		errs.ThrowSys(parseErr)
	}

	return icon, p
}

// // Config is for prompts, controlling fmt.Println colors, and what/when types print
// // function itself just inits / grabs those
// func config(path string) (string, string) {

// 	path := homeSlashPath()
// 	cf := configPath()
// 	config, err := os.ReadFile(cf)

// 	for err != nil { // files don't exist / deleted

// 		prompt := "❯➤ " // default

// 		// user confirms dir  / not making sys changes w/out input
// 		confNotFound(path, prompt)
// 		writeDefaults(prompt, path, cf)
// 		configsOutro(prompt)
// 		err = nil
// 		config, _ = os.ReadFile(path)

// 	}

// 	return string(config)
// }

// func confNotFound(path string, prompt string) {

// 	fmt.Println("<bx>", "No config file found.")
// 	fmt.Println("<bx>", "Default config path is: "+path)
// 	fmt.Printf(" %s Create new config? (y | n) \t\n(bx will just exit if not 'y') : ", prompt)

// 	if grabInput() == "y" {
// 		//bx's homedir in config or wherever
// 		os.Mkdir(path, 0755)

// 	} else { // this is if manually exited
// 		fmt.Println(prompt, "sure thing, pardner")
// 		os.Exit(0)
// 	}
// }

// func grabInput() string {
// 	var input string
// 	_, err := fmt.Scanln(&input)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	return input
// }

// func configsOutro(prompt string) {

// 	printsWhenDone := []string{
// 		"Folder created.",
// 		"Configured defaults.", //           color the command?
// 		"Feel free to change them by running {bx defaults}."}

// 	for iStr := range printsWhenDone {
// 		fmt.Println(prompt, printsWhenDone[iStr])
// 		time.Sleep(700 * time.Millisecond)
// 	}
// 	fmt.Println(" ") //buffer
// }

// func writeDefaults(prompt string, path string, confFile string) {

// 	outbound := []byte(
// 		"todoSymbol = ○" +
// 			"\n" + "promptIcon = " + prompt +
// 			"\n" + "# days ahead to check for" +
// 			"\n" + "days = 3")

// 	os.WriteFile(confFile, outbound, 0644)

// 	files := []string{
// 		"tailbox_cache.txt",
// 		"events.txt",
// 		"recurrent_reminders.txt",
// 		"todos.txt",
// 		"crons.txt"}

// 	for index, file := range files {
// 		frmt := fmt.Sprint(index+1) + " > Creating " + file + " in " + path
// 		fmt.Println("<bx>", frmt)
// 		os.WriteFile(
// 			(path + file), []byte(""), 0644)
// 	}
// }
