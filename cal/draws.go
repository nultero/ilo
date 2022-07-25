package cal

import "github.com/gdamore/tcell/v2"

func drawBox(scr tcell.Screen) {
	drawTopLine(scr)
	for y := 1; y < minHeight; y++ {
		if y%3 == 0 {
			drawMidLine(y, scr)
		}
		drawOutLine(y, scr)
	}
	drawBtmLine(scr)
}

func drawOutLine(y int, scr tcell.Screen) {
	scr.SetContent(0, y, thinSet.vt, empt, defStyle)
	scr.SetContent(minWidth, y, thinSet.vt, empt, defStyle)
}

func drawMidLine(y int, scr tcell.Screen) {
	for x := 1; x < minWidth; x++ {
		scr.SetContent(x, y, thinSet.fl, empt, defStyle.Foreground(tcell.ColorRed))
	}
}

func drawTopLine(scr tcell.Screen) {
	scr.SetContent(0, 0, thinSet.ul, empt, defStyle)
	for x := 1; x < minWidth; x++ {
		scr.SetContent(x, 0, thinSet.fl, empt, defStyle)
	}
	scr.SetContent(minWidth, 0, thinSet.ur, empt, defStyle)
}
func drawBtmLine(scr tcell.Screen) {
	scr.SetContent(0, minHeight, thinSet.bl, empt, defStyle)
	for x := 1; x < minWidth; x++ {
		scr.SetContent(x, minHeight, thinSet.fl, empt, defStyle)
	}
	scr.SetContent(minWidth, minHeight, thinSet.br, empt, defStyle)
}
