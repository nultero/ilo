package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func _args() []string {
	return []string{
		"cron",
		"event",
		"idea",
		"recurrent",
		"wishlist",
	}
}

func HandlesPrompts(promptIcon string, promptType string) string {

	var opts []string

	switch promptType {
	case "args":
		opts = _args()
	}

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
