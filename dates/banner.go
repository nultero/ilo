package dates

import (
	"strings"
	"time"
)

func BannerDate() (string, int) {
	_, mn, day := time.Now().Date()
	ms := strings.ToLower(mn.String())
	if len(ms) > 3 {
		ms = ms[:3]
	}

	bn := make([]rune, len(ms))
	for i, c := range ms {
		bn[i] = chars[c]
	}

	bstr := string(bn)
	return bstr, day
}
