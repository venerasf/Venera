package src

import lua "github.com/yuin/gopher-lua"

// Main profile holds users configs
type Profile struct {
	Prompt 	string
	Script 	string // script path
	SSet 	bool   // Script setted
	BPath 	string // Base path

	Globals map[string]string // Script Global variables
	State 	*lua.LState
}

// Live prefix for prompt configs
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

type ScriptTAG struct {
	Path 	string
	Tag 	[]string
}