package vars

var Args = []string{
	"event",
	"recurrent",
	"task",
	"idea",
}

const DaysAhead = "days ahead"
const defPrompt = "❯➤"
const ConfFile = "confFile"
const DataDir = "dataDir"
const PromptIcon = "prompt icon"

var ConfMap = map[string]string{
	ConfFile: "$USER/.config/cursebox.yaml",
	DataDir:  "$USER/.cursebox",
}

const s = ": "

var DefaultSettings = []string{
	DataDir + s + ConfMap[DataDir],
	PromptIcon + s + defPrompt,
	DaysAhead + s + "3", // only hardcoded as this default
}
