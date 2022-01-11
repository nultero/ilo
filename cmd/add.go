package cmd

import (
	"bx/cmd/sched"
	"bx/cmd/vars"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",

	ValidArgs: vars.Args,
	Args:      cobra.MaximumNArgs(1),
	Run:       add,
}

func add(cmd *cobra.Command, args []string) {
	sched.Selector()
}

func init() {
	rootCmd.AddCommand(addCmd)
}
