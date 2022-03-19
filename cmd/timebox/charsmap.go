package timebox

var chars = map[rune]rune{
	'a': '🄰', 'b': '🄱', 'c': '🄲', 'd': '🄳',
	'e': '🄴', 'f': '🄵', 'g': '🄶', 'h': '🄷',
	'i': '🄸', 'j': '🄹', 'k': '🄺', 'l': '🄻',
	'm': '🄼', 'n': '🄽', 'o': '🄾', 'p': '🄿',
	'q': '🅀', 'r': '🅁', 's': '🅂', 't': '🅃',
	'u': '🅄', 'v': '🅅', 'w': '🅆', 'x': '🅇',
	'y': '🅈', 'z': '🅉',
}

func trnsltMnth(m string) string {
	chrs := []rune{}
	for _, lt := range m {
		if sqr, ok := chars[lt]; ok {
			chrs = append(chrs, sqr)
			chrs = append(chrs, ' ')
		}
	}
	return string(chrs)
}
