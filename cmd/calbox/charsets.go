package calbox

type charSet struct {
	fl, bl, br, pl, pr, ul, ur, vt string
}

var thinSet = charSet{
	fl: "─",
	bl: "└",
	br: "┘",
	pl: "┤",
	pr: "├",
	ul: "┌",
	ur: "┐",
	vt: "│",
}

var boldSet = charSet{
	fl: "━",
	bl: "┗",
	br: "┛",
	pl: "┫",
	pr: "┣",
	ul: "┏",
	ur: "┓",
	vt: "┃",
}
