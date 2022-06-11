package cal

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func newScr() tcell.Screen {
	scr, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := scr.Init(); err != nil {
		panic(err)
	}

	if !hasMinSize(scr) {
		fmt.Println("termulator screen size too small; this message not implemented at runtime yet")
		fmt.Println("crashing in response")
		os.Exit(1)
	}

	return scr
}

func render(scr tcell.Screen) {
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	scr.SetStyle(defStyle)

	sigs := make(chan byte)
	quit := make(chan struct{})

	go func() {
		for {
			switch ev := scr.PollEvent().(type) {
			case *tcell.EventResize:
				scr.Sync()

			case *tcell.EventKey:
				switch ev.Key() {

				case tcell.KeyEscape, tcell.KeyCtrlC, tcell.KeyCtrlQ:
					quit <- struct{}{}
					scr.Fini()
					return

				case tcell.KeyLeft, tcell.KeyEnter:
					sigs <- byte(2)
				}
			}
		}
	}()

	x, _ := scr.Size()

runloop:
	for {
		select {
		case <-quit:
			break runloop

		// TODOOOO replace dumb interactivity w/ calendar moves
		case b := <-sigs:
			if b == 2 {
				for i := 0; i < x; i++ {
					scr.SetContent(i, 1, 'x', []rune{}, defStyle)
				}
			}

		}

		scr.Show()
	}

	scr.Fini()
}
