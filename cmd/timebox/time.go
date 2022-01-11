package timebox

import (
	"fmt"
	"strings"
	"time"
)

type tb struct {
	DayN, DayWk, Month string
}

func Today() tb {
	t := time.Now()
	dayFmt := t.Format("02 Mon")
	spl := strings.Split(dayFmt, " ")
	day := spl[0]
	wk := spl[1]
	m := strings.ToLower(t.Month().String()[0:3])
	return tb{
		DayN:  day,
		DayWk: wk,
		Month: m,
	}
}

func (t *tb) String() string {
	return fmt.Sprintf(
		"%v %v %v",
		t.DayN, t.DayWk, trnsltMnth(t.Month),
	)
}
