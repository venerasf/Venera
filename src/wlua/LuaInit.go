package wlua

import (
	// "venera/src"
	// "os"
	libs "github.com/vadv/gopher-lua-libs"
	"github.com/yuin/gopher-lua"
)


// TODO: Create a sctruct and mas to methods or as var, i don't like globals
// This global var receives metadata from the running script
var Metad METADATA
// Global variable vars
var LoadVar = make(map[string]VarDef)
var LuaProf LuaProfile

// Execute arbitrary strings
func LuaExecString(l *lua.LState,s string) {
	l.DoString(s)
}

func loadLibs(l *lua.LState) {
	libs.Preload(l)
	//l.PreloadModule("lio",lio.Loader)
}

func Sets(l *lua.LState) {
	l.SetGlobal("RandomString", l.NewFunction(RandomString))

	// PrettyPrints
	l.SetGlobal("PrintSuccs", l.NewFunction(PrintSuccs))
	l.SetGlobal("PrintErr", l.NewFunction(PrintErr))
	l.SetGlobal("PrintInfo", l.NewFunction(PrintInfo))
	l.SetGlobal("PrintSuccsln", l.NewFunction(PrintSuccsln))
	l.SetGlobal("PrintErrln", l.NewFunction(PrintErrln))
	l.SetGlobal("PrintInfoln", l.NewFunction(PrintInfoln))
	l.SetGlobal("Print", l.NewFunction(Print))
	l.SetGlobal("Println", l.NewFunction(Println))

	//Input/prompt
	l.SetGlobal("Input",l.NewFunction(Input))
	//Open file
	l.SetGlobal("Open",l.NewFunction(Open))

	l.SetGlobal("Meta", l.NewFunction(Meta))
	l.SetGlobal("LoadVars",l.NewFunction(LoadVars))
	l.SetGlobal("Call",l.NewFunction(LuaProf.Calls))
	loadLibs(l)
}

// Start Lua sandbox run once at a time
// The func will instance the source code
// so it can be configured from main prompt
// return lua.state and if the script culd be
// runned (true) or not (false)
func LuaInitUniq(p LuaProfile) (*lua.LState, bool) {
	// Activate script global variables
	// it can pass to another script running in chain
	LuaProf = p
	l := lua.NewState()
	Sets(l) // Set main funcs
	err := l.DoFile(p.Script)
	if err != nil {
		println(err.Error())
		return nil,false
	}
	l.DoString("Init()")
	SetFromVarsScriptGlobals(l,p)
	return l,true
}

// Run a LState already instatiated
func LuaRunUniq(l *lua.LState) {
	l.DoString("Main()")
}



func LuaRunChaining(p LuaProfile) {
	for _,i := range(p.Scriptslist) {
		p.Script = i
		LuaInitChain(p)
	}
}


// Start lua chai for working with multiple scripts
// when we use tags to index.
func LuaInitChain(p LuaProfile) {
	l := lua.NewState()
	defer l.Close() // Applying close() here
	Sets(l)
	err := l.DoFile(p.Script)

	if p.Globals["VERBOSE"] == "true" {
		println("-> "+p.Script)
	}
	if err != nil {
		println(err.Error())
		return
	}
	l.DoString("Init()")
	SetFromGlobals(l,p)
	l.DoString("Main()")
}



// LuaFreeScript deletes everything of a script from the memory
func LuaFreeScript() {
	LoadVar = make(map[string]VarDef)
	Metad = METADATA{}
	LuaProf = LuaProfile{}
}