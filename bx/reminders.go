package bx

import (
	"fmt"
)

// Checks bx's events cache -- if old, re-runs reminder calculations,
// if not old, prints the date with whatever is in the cache.
func RunReminders(path, icon, config string) {

	return

	// now := time.Now()
	// today := now.Format("02 Mon")
	// month := now.Month().String()[0:3]

	// fmt.Println(icon, month, today)

	// cache := readCache(path)

	// if cacheIsOld(cache[0], today, path) {

	// 	nDays := getNumDaysAhead(config)
	// 	items := doChecks(nDays, today, month, path)
	// 	fmt.Println(len(items))
	// 	// cacheResults(today, items)

	// } else {
	// 	c := cache[1:]
	// 	if !fn.IsEmpty(c) {
	// 		for i := range c {
	// 			fmt.Println(c[i])
	// 		}
	// 	}
	// }
}

func doChecks(nDays int, today, month, path string) []string {
	// chx := []string{"events", "recurrents"}

	fmt.Println("stopping here")
	return []string{}
	// for i := range chx {

	// 	p := path + glob(chx[i])
	// 	f, r := os.ReadFile(p)
	// 	if r != nil {
	// 		errs.ThrowX(r, fmt.Sprintf("error with the bx file at '%s'", p))
	// 	}

	// 	if chx[i] == "events" { // autocleaner logic

	// 		rev := -1 * nDays // flips to behind current day

	// 	}

	// 	fmt.Println(string(f))
	// }

	// return chx
}
