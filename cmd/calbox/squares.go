package calbox

func baseLft(cs charSet, n int) []string {
	b := []string{cs.ul}
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			b = append(b, cs.pr)
		} else {
			b = append(b, cs.vt)
		}
	}
	b = append(b, cs.bl)
	return b
}
