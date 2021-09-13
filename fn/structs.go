package fn

type Day struct {
	Index     int
	MonthName string
}

type month struct {
	index int
	name  string
	days  int
}

type Bus struct {
	PromptIcon string
	FileType   string
	Funct      string
	Path       string
	Help       bool
}

func DefaultBus(icon, path string) Bus {
	return Bus{
		PromptIcon: icon,
		FileType:   "",
		Funct:      "",
		Path:       path,
		Help:       false,
	}
}
