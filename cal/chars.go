package cal

var empt = []rune{}

type charSet struct {
	fl, bl, br, pl, pr, ul, ur, vt rune
}

var thinSet = charSet{
	fl: '─',
	bl: '└',
	br: '┘',
	pl: '┤',
	pr: '├',
	ul: '┌',
	ur: '┐',
	vt: '│',
}

var boldSet = charSet{
	fl: '━',
	bl: '┗',
	br: '┛',
	pl: '┫',
	pr: '┣',
	ul: '┏',
	ur: '┓',
	vt: '┃',
}
