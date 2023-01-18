package wlua

import (
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)


// This functions allows to run script from within other script.
func (s LuaProfile)Calls(l *lua.LState) int {
	// Chain global var disable chaining execution.
	if s.Globals["chain"] == "off" || s.Globals["chain"] == "false" {
		return 1
	}

	newFileVars := new(map[string]VarDef)
	newFile := lua.NewState()
	Sets(newFile)
	err := newFile.DoFile(l.ToString(1))	
	if err != nil {
		println(err.Error())
	}

	if err := gluamapper.Map(l.GetGlobal("Vars").(*lua.LTable), &newFileVars); err != nil {
		panic(err)
	}


	for key,value := range(*newFileVars) {
		newFile.DoString("Vars."+key+".VALUE=\""+value.VALUE+"\"")
	}

	newFile.DoString("Main()")
	newFile.Close()
	return 1
}