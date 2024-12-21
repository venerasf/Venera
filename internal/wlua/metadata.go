/*
	This file define functions o handle metadata of a in use script
*/
package wlua

import (
	"fmt"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

// Load metadata
func Meta(L *lua.LState) int {
	if err := gluamapper.Map(L.GetGlobal("METADATA").(*lua.LTable), &Metad); err != nil {
		panic(err)
	}
	return 1
}

func MetaShow() {
	println("## AUTHOR/S ##")
	MetaListAuthors()
	println("\n## TAGS ##")
	MetaListCats()
	println("\n## INFO ##")
	MetaShowInfo()
}

func MetaListAuthors() {
	for i := range(Metad.AUTHOR) {
		fmt.Printf("%d) %s\n",i+1,Metad.AUTHOR[i])
	}
}

func MetaListCats() {
	for i := range(Metad.TAGS) {
		fmt.Printf("%d) %s\n",i+1,Metad.TAGS[i])
	}
}

func MetaShowInfo() {
	println(Metad.INFO)
}

// return categories of a script
func ScriptGetTags(path string) []string {
	newMeta := METADATA{}
	//println(path)
	aux := lua.NewState()
	defer aux.Close()
	Sets(aux)
	err := aux.DoFile(path)
	if err != nil {
		return []string{"nil(f)"}
	}
	//println(path)
	x := aux.GetGlobal("METADATA").(*lua.LTable)
	if err = gluamapper.Map(x, &newMeta); err != nil {
		return []string{"nil(m)"}
	}
	return newMeta.TAGS
}

