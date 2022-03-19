package sched

import (
	"bx/cmd/calbox"
	"bx/cmd/fsys"
	"bx/cmd/timebox"
	"bx/cmd/vars"
	"fmt"
	"os"

	"github.com/nultero/tics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Checks the current day against the schedule --
// Sample will see if there is a cache first and if not
// then it will run everything it needs to recompute.
//
// Final result is it prints the day, the fmt'd month,
// and whatever is in the schedule.
func Sample(cmd *cobra.Command, args []string) {
	tb := timebox.Today()
	pi := viper.GetString(vars.PromptIcon)
	fmt.Printf("%v  %v\n", pi, tb.String())

	// cacheBytes, err := fsys.ReadCache()
	_, err := fsys.ReadCache()
	if err != nil {
		if !os.IsNotExist(err) {
			tics.ThrowSys(Sample, err)
		}

		// perform the checks anyway, because cache does not exist
	}

	s := calbox.DrawBasic()
	fmt.Println(s)
	//check against cache, see if same as current day

	// if so, just print
	// if not, recompute

	// js := tics.ToJson([]byte(`{"yes": 5}`))
	// fmt.Println(js)

	// daysAhead := viper.GetString(vars.DaysAhead)
	// fmt.Println(daysAhead)
}
