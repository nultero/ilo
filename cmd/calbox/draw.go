package calbox

const (
	wks = 4
	dys = 7
)

func DrawBasic() string {
	n := wks * 2
	cs := thinSet
	cal := baseLft(cs, n)

	for wk := 0; wk < n; wk++ {

		if wk%2 == 0 {
			for i := 0; i < dys*4; i++ {
				cal[wk] += cs.fl
			}

		} else {
			for i := 0; i < dys*4; i++ {
				cal[wk] += "0"
			}

		}
	}

	s := ""
	for _, rw := range cal {
		s += rw + "\n"
	}

	return s
}
