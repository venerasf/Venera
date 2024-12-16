package wlua

import (
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)


// This functions allows to run script from within other script.
func Calls(l *lua.LState) int {
	// Chain global var disable chaining execution, disallowing calls.
	if LuaProf.Globals["chain"] == "off" || LuaProf.Globals["chain"] == "false" {
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


	// TODO: test it better
	for key,value := range(*newFileVars) {
		lvalue := newFile.GetGlobal("VARS")
		lvalue1 := newFile.GetField(lvalue, key)
		newValue := lua.LString(value.VALUE)
		newFile.SetField(lvalue1, "VALUE", newValue)

		//newFile.DoString("Vars."+key+".VALUE=\""+value.VALUE+"\"")
	}

	newFile.DoString("Main()")
	newFile.Close()
	return 1
}