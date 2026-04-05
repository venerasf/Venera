// Variables that a script needs setted to run
// Like remote host, port etc...
package wlua

import (
	"strings"
	"venera/internal/types"
	"venera/internal/utils"

	//"strings"
	"github.com/cheynewallace/tabby"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func getVarsTable(L *lua.LState) (*lua.LTable, bool) {
	vars := L.GetGlobal("VARS")
	varsTable, ok := vars.(*lua.LTable)
	if !ok {
		utils.PrintErr("Failed to load variables: script VARS table is missing")
		return nil, false
	}

	return varsTable, true
}

func getScriptVarTable(varsTable *lua.LTable, key string) (*lua.LTable, bool) {
	varKey := key
	value := varsTable.RawGetString(varKey)
	if _, ok := value.(*lua.LTable); !ok {
		varsTable.ForEach(func(k lua.LValue, _ lua.LValue) {
			if strings.EqualFold(k.String(), key) {
				varKey = k.String()
			}
		})
		value = varsTable.RawGetString(varKey)
	}

	varTable, ok := value.(*lua.LTable)
	if !ok {
		return nil, false
	}

	return varTable, true
}

// Load vars
func LoadVars(L *lua.LState) int {
	varsTable, ok := getVarsTable(L)
	if !ok {
		return 0
	}

	if err := gluamapper.Map(varsTable, &LoadVar); err != nil {
		utils.PrintErr("Failed to load script variables: " + err.Error())
		return 0
	}
	//print
	return 1
}

// List variables
func VarsList() {
	t := tabby.New()
	t.AddHeader("VARIABLE", "DEFAULT", "REQUIRED", "DESCRIPTION")
	for i, j := range LoadVar {
		t.AddLine(i, j.VALUE, j.REQUIRED, j.DESCRIPT)
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
	for i, _ := range LoadVar {
		if i == key {
			ex = true // The var exists
		}
	}

	if ex {
		//L.DoString(fmt.Sprintf(`VARS.%s.VALUE="%s"`, key, value))
		varsTable, ok := getVarsTable(L)
		if !ok {
			return
		}

		lvalue1, ok := getScriptVarTable(varsTable, key)
		if !ok {
			utils.PrintErr("Variable not found in script: " + key)
			return
		}

		newValue := lua.LString(value)
		L.SetField(lvalue1, "VALUE", newValue)

		LoadVars(L) // Update var struct
		utils.PrintSuccs(key, " <- ", value)
		// println("[\u001B[1;32mOK\u001B[0;0m]",)
	} else {
		utils.PrintErr(key, " <- ", value)
	}
}

// InstSet variables from globals
func SetFromGlobals(L *lua.LState, p *types.Profile) {
	varsTable, ok := getVarsTable(L)
	if !ok {
		return
	}

	varsByUpper := make(map[string]string)
	varsTable.ForEach(func(k lua.LValue, _ lua.LValue) {
		varsByUpper[strings.ToUpper(k.String())] = k.String()
	})

	for i := range p.Globals {
		varKey, exists := varsByUpper[strings.ToUpper(i)]
		if !exists {
			continue
		}

		lvalue1 := varsTable.RawGetString(varKey)
		if _, ok := lvalue1.(*lua.LTable); !ok {
			continue
		}

		newValue := lua.LString(p.Globals[i])
		L.SetField(lvalue1, "VALUE", newValue)

		//L.DoString(fmt.Sprintf(`VARS.%s.VALUE="%s"`,i,p.Globals[i]))
		//L.DoString("VARS."+i+".VALUE=\""+p.Globals[i]+"\"")
	}
}

// Set vars from globals when running `use script/luascript.lua`
func SetFromVarsScriptGlobals(L *lua.LState, p *types.Profile) {
	for i := range LoadVar {
		for j, y := range p.Globals {
			if strings.ToUpper(j) == i {
				SetVarValue(L, i, y)
				break
			}
		}
	}
}

// Get vars from scripts
func GetVarsToChainTAGS(p *types.Profile) {
	//fmt.Println(p.Scriptslist)
	for _, f := range p.Scriptslist {
		L := lua.NewState()
		Sets(L)
		err := L.DoFile(f)
		if err != nil {
			utils.PrintErr("Failed to load script " + f + ": " + err.Error())
			L.Close()
			continue
		}

		varsTable, ok := getVarsTable(L)
		if !ok {
			L.Close()
			continue
		}

		auxVar := make(map[string]VarDef)
		if err := gluamapper.Map(varsTable, &auxVar); err != nil {
			utils.PrintErr("Failed to load variables from script " + f + ": " + err.Error())
			L.Close()
			continue
		}

		for i, s := range auxVar {
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
