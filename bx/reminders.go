package bx

import (
	"fmt"
	"time"
)

// Checks bx's events cache -- if old, re-runs reminder calculations,
// if not old, will just print the date with whatever icon is present in config.
func RunReminders(path, icon, config string) {

	now := time.Now()
	today := now.Format("02 Mon")
	month := now.Month().String()[0:3]

	fmt.Println(icon, month, today)

	///// run checks
	cache := readCache(path)

	if cacheIsOld(cache[0], today, path) {

		fmt.Println("donkey sauce")

	} else {
		c := cache[1:]
		if !IsEmpty(c) {
			for i := range c {
				fmt.Println(c[i])
			}
		}
	}
}
