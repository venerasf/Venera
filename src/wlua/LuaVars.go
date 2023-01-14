// Variables that a script needs setted to run
// Like remote host, port etc...
package wlua

import (
	"github.com/cheynewallace/tabby"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// Global variable vars
var LoadVar = new(map[string]VarDef)

// Load vars
func LoadVars(L *lua.LState) int {
	if err := gluamapper.Map(L.GetGlobal("Vars").(*lua.LTable), &LoadVar); err != nil {
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
	t.Print()
}

func SetVarValue(L *lua.LState, key string, value string) {
	ex := false
	for i,_ := range(*LoadVar) {
		if i == key {
			ex = true
		}
	}
	if ex {
		L.DoString("Vars."+key+".VALUE=\""+value+"\"")
		LoadVars(L)
		println("[\u001B[1;32mOK\u001B[0;0m]",key,"<-",value)
	} else {
		println("[\u001B[1;31m!\u001B[0;0m] No variable",key,"<-",value)
	}
}