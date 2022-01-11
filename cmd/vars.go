package cmd

var defPrompt = "❯➤"
var confFile = "confFile"
var dataDir = "dataDir"
var promptIcon = "prompt icon"

var confMap = map[string]string{
	confFile: "$USER/.config/cursebox.yaml",
	dataDir:  "$USER/.cursebox",
}

var defaultSettings = []string{
	dataDir + ": " + confMap[dataDir],
	promptIcon + ": " + defPrompt,
}
