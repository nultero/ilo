package cmd

import (
	"ilo/cal"

	"github.com/spf13/cobra"
)

var cfgFile string = "?"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ilo",
	Short: "a calendar; 'interactive ledger for ops' in fancy",

	Run: func(cmd *cobra.Command, args []string) {
		cal.Glance()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config/.ilo.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
