package wlua

import (
	//"venera/src"
	//"os"
	libs "github.com/vadv/gopher-lua-libs"
	"github.com/yuin/gopher-lua"
)

// This global var receives metadata from the running script
var Metad Metadata
// Global variable vars
var LoadVar = new(map[string]VarDef)




// Execute arbitrary strings
func LuaExecString(l *lua.LState,s string) {
	l.DoString(s)
}

func loadLibs(l *lua.LState) {
	libs.Preload(l)
	//l.PreloadModule("lio",lio.Loader)
}

func sets(l *lua.LState) {
	l.SetGlobal("RandomString", l.NewFunction(RandomString))

	// PrettyPrints
	l.SetGlobal("PrintSuccs", l.NewFunction(PrintSuccs))
	l.SetGlobal("PrintErr", l.NewFunction(PrintErr))
	l.SetGlobal("PrintInfo", l.NewFunction(PrintInfo))

	l.SetGlobal("Meta", l.NewFunction(Meta))
	l.SetGlobal("LoadVars",l.NewFunction(LoadVars))
	loadLibs(l)
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
		println(err.Error())
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



// LuaFreeScript deletes everything of a script from the memory
func LuaFreeScript() {
	LoadVar = new(map[string]VarDef)
	Metad = Metadata{}
}