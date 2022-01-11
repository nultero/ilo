package bxd

// import (
// 	"bx/cmds/bxfiles"
// 	"bx/errs"
// 	"bx/fn"
// 	"fmt"
// 	"os"
// 	"time"
// )

// // Checks bx's events cache -- if old, re-runs reminder calculations,
// // if not old, prints the date with whatever is in the cache.
// func RunReminders(path, icon, config string) {

// 	now := time.Now()
// 	today := now.Format("02 Mon")
// 	month := now.Month().String()[0:3]

// 	printStatus(icon, month, today)

// 	cache := readCache(path)

// 	if cacheIsOld(cache[0], today, path) {

// 		nDays := getNumDaysAhead(config)
// 		items := doChecks(nDays, today, month, path)
// 		events := checkEvents(nDays, today, month, path)

// 		cacheResults(cacheMesh(today, items))

// 	} else {
// 		cache = cache[1:]
// 		if !fn.IsEmpty(cache) {
// 			for _, line := range cache {
// 				fmt.Println(line)
// 			}
// 		}
// 	}
// }

// func checkEvents(nDays int, today, month, path string) []string {

// 	// Events are *not* recurrent, and ergo need an autocleaner
// 	// for ones that have already passed

// 	file := bxfiles.Events()

// }

// func doChecks(nDays int, today, month, path string) []string {

// 	chx := bxfiles.CheckFiles()

// 	for _, file := range chx {
// 		_, err := os.ReadFile(path + file)
// 		if err != nil {
// 			errs.ThrowX(err, fmt.Sprintf("error with the bx file at '%s'", path+file))
// 		}

// 		//

// 		//

// 		d := fn.SetRollDay(today, month)

// 		//

// 		if file == bxfiles.Events() { // autocleaner logic

// 			for i := 0; i < nDays; i++ {
// 				d = fn.RollBackDay(d)
// 			}

// 			rev := -1 * nDays // flips to behind current day

// 			for rev < nDays {
// 				d = fn.RollForwardDay(d)
// 				rev++
// 				fmt.Println("d:= ", d.Index, d.MonthName)
// 			}

// 		} else {
// 			for i := 0; i < nDays; i++ {
// 				d = fn.RollForwardDay(d)
// 				fmt.Println("d:= ", d.Index, d.MonthName)
// 			}
// 		}

// 		// still need to do the file parsing on the date checks --
// 		// and do it right, since the last time it was buggy
// 	}

// 	return chx
// }

// // func autoclean(index int) {
// // 	p := pathGlob("events")
// // 	f, _ := os.ReadFile(p)
// // 	lines := strings.Split(string(f), "\n")
// // 	cleaned := make([]string, len(lines)-1)
// // 	cleaned = append(lines[:index], lines[index+1:]...)
// // 	writeOut(p, cleaned, false)

// // 	s := colorate("** bx's autoclean has run", []int{90, 210, 110})
// // 	fmt.Println(s)
// // }

// func printStatus(icon, month, day string) {
// 	fmt.Println(icon+" ", month, day)
// }
