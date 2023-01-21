// Variables that a script needs setted to run
// Like remote host, port etc...
package wlua

import (
	"github.com/cheynewallace/tabby"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// Load vars
func LoadVars(L *lua.LState) int {
	if err := gluamapper.Map(L.GetGlobal("VARS").(*lua.LTable), &LoadVar); err != nil {
		panic(err)
	}
	//print
	return 1
}

func VarsList() {
	t := tabby.New()
	t.AddHeader("VARIABLE","DEFAULT","NEEDED","DESCRIPTION")
	for i,j := range(*LoadVar) {
		t.AddLine(i,j.VALUE,j.NEEDED,j.DESCRIPT)
	}
	print("\n")
	t.Print()
	print("\n")
}


// Set variales in manual use
func SetVarValue(L *lua.LState, key string, value string) {
	ex := false
	for i,_ := range(*LoadVar) {
		if i == key {
			ex = true
		}
	}
	if ex {
		L.DoString("VARS."+key+".VALUE=\""+value+"\"")
		LoadVars(L)
		println("[\u001B[1;32mOK\u001B[0;0m]",key,"<-",value)
	} else {
		println("[\u001B[1;31m!\u001B[0;0m] No variable",key,"<-",value)
	}
}

// Set variables from globals 
func SetFromGlobals(L *lua.LState,p LuaProfile) {
	vars := new(map[string]VarDef)

	if err := gluamapper.Map(L.GetGlobal("VARS").(*lua.LTable), &vars); err != nil {
		panic(err)
	}

	for i := range(p.Globals) {
		//println("VARS."+i+".VALUE=\""+p.Globals[i]+"\"")
		L.DoString("VARS."+i+".VALUE=\""+p.Globals[i]+"\"")
	}
}


// Set vars from globals to `use script/luascript.lua`
func SetFromVarsScriptGlobals(L *lua.LState, p LuaProfile) {
	for i := range(*LoadVar) {
		for j,y := range(p.Globals) {
			if j==i {
				SetVarValue(L,i,y)
				break
			}
		}
	}
}