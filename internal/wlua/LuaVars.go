// Variables that a script needs setted to run
// Like remote host, port etc...
package wlua

import (
	"strings"
	"venera/internal/utils"

	//"strings"
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

// List variables
func VarsList() {
	t := tabby.New()
	t.AddHeader("VARIABLE","DEFAULT","REQUIRED","DESCRIPTION")
	for i,j := range(LoadVar) {
		t.AddLine(i,j.VALUE,j.NEEDED,j.DESCRIPT)
	}
	print("\n")
	t.Print()
	print("\n")
}


/* 
	Set variales in manual use
*/
func SetVarValue(L *lua.LState, key string, value string) {
	key = strings.ToUpper(key)

	ex := false
	// Iterate all avaliable vars
	for i, _ := range(LoadVar) {
		if i == key {
			ex = true // The var exists
		}
	}

	if ex {
		//L.DoString(fmt.Sprintf(`VARS.%s.VALUE="%s"`, key, value))

		lvalue := L.GetGlobal("VARS")
		lvalue1 := L.GetField(lvalue, key)
		newValue := lua.LString(value)
		L.SetField(lvalue1, "VALUE", newValue)

		LoadVars(L) // Update var struct
		utils.PrintSuccs(key," <- ",value)
		// println("[\u001B[1;32mOK\u001B[0;0m]",)
	} else {
		utils.PrintErr(key," <- ",value)
	}
}

// InstSet variables from globals 
func SetFromGlobals(L *lua.LState,p LuaProfile) {
	vars := new(map[string]VarDef)

	if err := gluamapper.Map(L.GetGlobal("VARS").(*lua.LTable), &vars); err != nil {
		panic(err)
	}

	for i := range(p.Globals) {
		//println("VARS."+i+".VALUE=\""+p.Globals[i]+"\"")
		lvalue := L.GetGlobal("VARS")
		lvalue1 := L.GetField(lvalue, i)
		newValue := lua.LString(p.Globals[i])
		L.SetField(lvalue1, "VALUE", newValue)

		//L.DoString(fmt.Sprintf(`VARS.%s.VALUE="%s"`,i,p.Globals[i]))
		//L.DoString("VARS."+i+".VALUE=\""+p.Globals[i]+"\"")
	}
}


// Set vars from globals when running `use script/luascript.lua`
func SetFromVarsScriptGlobals(L *lua.LState, p LuaProfile) {
	for i := range(LoadVar) {
		for j,y := range(p.Globals) {
			if strings.ToUpper(j) == i {
				SetVarValue(L,i,y)
				break
			}
		}
	}
}


// Get vars from scripts
func GetVarsToChainTAGS(p LuaProfile) {
	//fmt.Println(p.Scriptslist)
	for _,f := range(p.Scriptslist) {
		L := lua.NewState()
		Sets(L)
		L.DoFile(f)

		auxVar := make(map[string]VarDef)
		if err := gluamapper.Map(L.GetGlobal("VARS").(*lua.LTable), &auxVar); err != nil {
			panic(err)
		}

		for i,s := range(auxVar) {
			if LoadVar[i].VALUE == "" {
				//fmt.Println(i,LoadVar[i].VALUE)
				LoadVar[i] = s
			}
		}

		L.Close()
	}
}

// Populate map that keep all variables from allf scripts grouped
/*func PopulateLoadVarsFromGlobals(p LuaProfile) {
	//xVars := *LoadVar
	for i,v := range(p.Globals) {
		LoadVar[i].VALUE = v
	}
}*/