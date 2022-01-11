package calbox

const (
	flat = "━"
	up   = "┻"
	down = "┳"
	blft = "┗"
	brt  = "┗"
	ulft = "┏"
	urt  = "┓"
	vt   = "┃"
)

func baseLft(n int) []string {
	b := []string{ulft}
	for i := 0; i < n; i++ {
		b = append(b, vt)
	}
	b = append(b, blft)
	return b
}
