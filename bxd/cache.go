package bxd

// import (
// 	"bx/cmds/bxfiles"
// 	"bx/errs"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func readCache(path string) []string {

// 	p := path + bxfiles.Cache()

// 	c, err := os.ReadFile(p)
// 	if err != nil {
// 		errs.ThrowSys(err)
// 	}

// 	cache := string(c)
// 	return strings.Split(cache, "\n")
// }

// func cacheIsOld(cacheTop, day, path string) bool {
// 	if cacheTop != day {
// 		return true
// 	}

// 	cacheFile, cErr := os.Stat(path + bxfiles.Cache())
// 	if cErr != nil {
// 		errs.ThrowSys(cErr)
// 	}

// 	cModTime := cacheFile.ModTime()
// 	files := bxfiles.CheckFiles()

// 	for _, f := range files {
// 		fmt.Println(path, f)
// 		file, err := os.Stat(path + f)
// 		if err != nil {
// 			errs.ThrowSys(err)
// 		}

// 		fMod := file.ModTime()

// 		if cModTime.Before(fMod) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func cacheMesh(day string, items []string) []string {
// 	mesh := []string{day}
// 	for _, i := range items {
// 		mesh = append(mesh, i)
// 	}
// 	return mesh
// }

// func cacheResults(data []string) {

// 	tmp := ""
// 	for i := range data {
// 		tmp += data[i] + "\n"
// 	}

// 	// slice off last newline
// 	tmp = tmp[:len(tmp)-1]

// 	// os.WriteFile(
// 	// 	pathGlob("check"),
// 	// 	[]byte(tmp),
// 	// 	0644,
// 	// )
// }
