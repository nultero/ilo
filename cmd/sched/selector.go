package sched

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

func makebox(s tcell.Screen) {
	w, h := s.Size()

	if w == 0 || h == 0 {
		return
	}

	glyphs := []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}

	lx := rand.Int() % w
	ly := rand.Int() % h
	lw := rand.Int() % (w - lx)
	lh := rand.Int() % (h - ly)
	st := tcell.StyleDefault
	gl := ' '
	if s.Colors() > 256 {
		rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
		st = st.Background(rgb)
	} else if s.Colors() > 1 {
		st = st.Background(tcell.Color(rand.Int()%s.Colors()) | tcell.ColorValid)
	} else {
		st = st.Reverse(rand.Int()%2 == 0)
		gl = glyphs[rand.Int()%len(glyphs)]
	}

	for row := 0; row < lh; row++ {
		for col := 0; col < lw; col++ {
			s.SetCell(lx+col, ly+row, st, gl)
		}
	}
	s.Show()
}

func Selector() error {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err = s.Init(); err != nil {
		return err
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack))
	s.Clear()

	sig := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {

			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyRune:
					switch ev.Rune() {
					case 'j':
						close(sig)
						return
					}
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC, tcell.KeyRight:
					close(sig)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}

			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	rand.Seed(time.Hour.Nanoseconds())

	sct := struct{}{}
	cnt := 0
	dur := time.Duration(0)
	for cnt < 100000 {
		select {
		case <-sig:
			cnt = 100000
		case <-time.After(time.Millisecond * 50):
		}
		start := time.Now()
		makebox(s)
		cnt++
		dur += time.Since(start)

		// _, stillgoing := <-stop
		// if !stillgoing {
		// 	break
		// }
		// makebox(s)
	}
	s.Fini()
	fmt.Println(sct)

	return nil
}
