package wlua

import (
	//"venera/src"
	//"os"
	"github.com/yuin/gopher-lua"
)

// This global var receives metadata from the running script
var Metad Metadata

func sets(l *lua.LState) {
	l.SetGlobal("RandomString", l.NewFunction(RandomString))

	// PrettyPrints
	l.SetGlobal("PrintSuccs", l.NewFunction(PrintSuccs))
	l.SetGlobal("PrintErr", l.NewFunction(PrintErr))

	l.SetGlobal("Meta", l.NewFunction(Meta))
	l.SetGlobal("LoadVars",l.NewFunction(LoadVars))
}

// Start Lua sandbox run once at a time
// The func will instance the source code
// so it can be configured from main prompt
// return lua.state and if the script culd be
// runned (true) or not (false)
func LuaInitUniq(p LuaProfile) (*lua.LState, bool) {
	l := lua.NewState()
	//defer l.Close()
	sets(l) // Set main funcs
	err := l.DoFile(p.Script)
	if err != nil {
		return nil,false
	}
	l.DoString("Init()")
	return l,true
}

// Run a LState already instatiated
func LuaRunUniq(l *lua.LState) {
	l.DoString("Main()")
}


// Start lua chai for working with multiple scripts
func LuaInitChain(p LuaProfile) {
	//l := lua.NewState()
}