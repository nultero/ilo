package cmd

import (
	"bx/cmd/sched"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

// When run with no args, cursebox will print the current month & day,
// and briefly check some schedules.
var rootCmd = &cobra.Command{
	Use:   "bx",
	Short: "Dead simple thingamajig",
	Run:   sched.Sample,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// TODO calendar flag for vis?
	// rootCmd.Flags().BoolVarP(&RecurseFlag, "recursive", "r", false, "traverses subdirectories wherever novem was called")
}

func initConfig() {
	confMap = tics.CobraRootInitBoilerPlate(confMap, true)
	confPath := confMap[confFile]
	viper.SetConfigFile(confPath)

	// If a config file is found, read it in, else make one with prompt.
	err := viper.ReadInConfig()
	if err != nil {
		tics.RunConfPrompts("cursebox", confMap, defaultSettings)
		tics.ThrowQuiet("")
	}
}
