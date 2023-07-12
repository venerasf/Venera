package wlua

import (
	"venera/src/db"

	lua "github.com/yuin/gopher-lua"
)

type LuaProfile struct {
	Prompt 		string
	Script 		string   // script path
	Scriptslist []string // list of scripts for chaining
	BPath 		string   // Base path

	Globals map[string]string // Script Global variables
	State 	*lua.LState
	SSet 	bool // Script setted to validate if there is a script loaded
	Chain 	bool // Store the info, if it is running in tags mode

	Database *db.DBDef // Database for persistence data
}

// Script metadata
type METADATA struct {
	AUTHOR 	[]string
	VERSION string
	TAGS 	[]string
	INFO	string
}

/* Variavles in script
	TODO: use "REQUIRED" instead of "NEEDED".
 	This change will affect all scripts done.
 */
type VarDef struct {
	VALUE 		string
	NEEDED 		string
	DESCRIPT	string
}


type ScriptGlobals struct {
	Globals map[string]string
}