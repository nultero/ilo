package conf

import (
	"bx/bx"
	"bx/cmds/bxfiles"
	"bx/errs"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"
)

func Ok(basePath string) (string, error) {

	p := basePath + bxfiles.Config()
	config, err := os.ReadFile(p)

	if err != nil {
		return "", err
	}

	return string(config), nil
}

func Fix(basePath string) {

	if !fs.ValidPath(basePath) {
		prompt := "❯➤ " // default
		confNotFound(basePath, prompt)
		writeDefaults(prompt, basePath)
		configsOutro(prompt)
	}
}

func confNotFound(path string, prompt string) {

	fmt.Printf("%s No config file found. \n", prompt)
	fmt.Printf("%s Default config path is: %s \n", prompt, bx.Emph(path+bxfiles.Config()))
	fmt.Printf("%s Create new config? (y | n) \t\n(bx will just exit if not 'y') : ", prompt)

	if strings.ToLower(grabInput()) == "y" {
		os.Mkdir(path, 0755)

	} else { // this is if manually exited
		fmt.Println(prompt, "sure thing, pardner")
		os.Exit(0)
	}
}

func grabInput() string {
	var input string
	_, err := fmt.Scanln(&input)

	if len(input) == 0 {
		errs.ThrowQuiet("> no inputs passed; quitting")
	}

	if err != nil {
		errs.ThrowSys(err)
	}

	return input
}

func configsOutro(prompt string) {

	printsWhenDone := []string{
		"Folder created.",
		"Configured defaults.", //           color the command?
		"Feel free to change them by running {bx defaults}."}

	for _, s := range printsWhenDone {
		fmt.Println(prompt, s)
		time.Sleep(700 * time.Millisecond)
	}
	fmt.Println(" ")
}

func writeDefaults(prompt string, path string) {

	files := bxfiles.All()

	for i, file := range files {
		frmt := fmt.Sprint(i+1) + " > Creating " + file + " in " + path
		fmt.Println(prompt, frmt)
		if file != bxfiles.Config() {
			os.WriteFile(
				(path + file), []byte(""), 0644)
		} else {
			os.WriteFile(
				(path + file), bxfiles.DefaultConfig, 0644)

		}
	}
}
