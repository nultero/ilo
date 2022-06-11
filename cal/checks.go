package cal

import (
	"github.com/gdamore/tcell/v2"
)

const (
	minWidth  = 30
	minHeight = 15
)

func hasMinSize(scr tcell.Screen) bool {
	w, h := scr.Size()
	if w < minWidth || h < minHeight {
		return false
	}
	return true
}
