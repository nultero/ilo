package bx

import (
	"bx/errs"
	"os"
	"strings"
)

func readCache(path string) []string {

	p := path + glob("cache")

	c, err := os.ReadFile(p)
	if err != nil {
		errs.ThrowSys(err)
	}

	cache := string(c)
	return strings.Split(cache, "\n")
}

func cacheIsOld(cacheTop, day, path string) bool {
	if cacheTop != day {
		return true
	}

	cacheFile, cErr := os.Stat(path + glob("cache"))
	if cErr != nil {
		errs.ThrowSys(cErr)
	}

	cModTime := cacheFile.ModTime()

	fileChecks := []string{"events", "recurrents"}

	for i := range fileChecks {
		file, err := os.Stat(path + glob(fileChecks[i]))
		if err != nil {
			errs.ThrowSys(err)
		}

		fMod := file.ModTime()

		if cModTime.Before(fMod) {
			return true
		}
	}

	return false
}
