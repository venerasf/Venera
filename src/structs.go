package src

import lua "github.com/yuin/gopher-lua"

// Main profile holds users configs
type Profile struct {
	Prompt 	string
	Script 	string
	SSet 	bool // Script setted
	BPath 	string // Base path

	State 	*lua.LState
}

// Live prefix for prompt configs
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}