package sched

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// Checks the current day against the schedule --
// Sample will see if there is a cache first and if not
// then it will run everything it needs to recompute.
func Sample(cmd *cobra.Command, args []string) {
	t := time.Now().Format(time.RubyDate)
	fmt.Println(t)
}
