package wlua

import lua "github.com/yuin/gopher-lua"

type LuaProfile struct {
	Prompt 	string
	Script 	string // script path
	SSet 	bool // Script setted

	State 	*lua.LState
}

// Script metadata
type Metadata struct {
	AUTHOR 	[]string
	VERSION string
	CATS 	[]string
	INFO	string
}

// Variavles in script
type VarDef struct {
	VALUE 		string
	NEEDED 		string
	DESCRIPT	string
}