package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// Exports a string that is unchecked against any args.
// Pass to a separate function to map cleaner args in passthrough.
func HandleStringInput(promptstr string) string { // only takes/rds string
	fmt.Print(promptstr + " ")
	nex, _, e := bufio.NewReader(os.Stdin).ReadLine()
	for e != nil {
		fmt.Println(e)
	}
	return string(nex)
}

func HandleArguments(promptIcon string) string {

	ptui := promptui.Select{
		Label: promptIcon,
		Items: ValidKinds,
	}

	_, tmp, err := ptui.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return tmp
}

func HandleOptions(promptIcon string, opts []string) string {

	ptui := promptui.Select{
		Label: promptIcon,
		Items: opts,
	}

	_, tmp, err := ptui.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return tmp
}
