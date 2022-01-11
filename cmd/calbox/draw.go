package calbox

const (
	wks = 4
	dys = 7
)

func DrawBasic() string {
	n := wks * 2
	cal := baseLft(n)
	for wk := 0; wk < n; wk++ {
		if wk%2 == 0 {
			for i := 0; i < dys*4; i++ {
				cal[wk] += flat
			}
		}
	}

	s := ""
	for _, rw := range cal {
		s += rw + "\n"
	}

	return s
}
