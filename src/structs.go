package src

import lua "github.com/yuin/gopher-lua"

// Main profile holds users configs
type Profile struct {
	Prompt 		string
	Script 		string   // script path
	Scriptslist []string // list of scripts for chaining
	SSet 		bool     // Script setted
	BPath 		string   // Base path

	Globals map[string]string // Script Global variables
	State 	*lua.LState
	Chain 	bool // if it is running in tags mode
}

// Live prefix for prompt configs
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

type ScriptTAGInfo struct {
	Path 	string
	Tag 	[]string
	Info 	string
}