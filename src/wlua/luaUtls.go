// Ultils

package wlua

import (
	"fmt"
	"math/rand"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// Generate random string
func RandomString(L *lua.LState) int {
	var letterBytes = ""
	
	leng := L.ToInt(1) // get firt arg as int
	chars := L.ToString(2) // second arg as str
	b := make([]byte, leng)

	if strings.Contains(chars,"a-z") {
		letterBytes = letterBytes+"abcdefghijklmnopqrstuvwxyz"
	}
	if strings.Contains(chars,"A-Z") {
		letterBytes = letterBytes+"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if strings.Contains(chars,"0-9") {
		letterBytes = letterBytes+"1234567890"
	}

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	L.Push(lua.LString(b))
	return 1
}


// Pretty print funcs
func PrintSuccs(L *lua.LState) int {
	fmt.Printf("[\u001B[1;32mOK\u001B[0;0m]- %s",L.ToString(1))
	return 1
}
func PrintErr(L *lua.LState) int {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s",L.ToString(1))
	return 1
}
func PrintInfo(L *lua.LState) int {
	fmt.Printf("[\u001B[1;34mi\u001B[0;0m]- %s",L.ToString(1))
	return 1
}


// Pretty print with line ending
func PrintSuccsln(L *lua.LState) int {
	fmt.Printf("[\u001B[1;32mOK\u001B[0;0m]- %s",L.ToString(1))
	return 1
}
func PrintErrln(L *lua.LState) int {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s",L.ToString(1))
	return 1
}
func PrintInfoln(L *lua.LState) int {
	fmt.Printf("[\u001B[1;34mi\u001B[0;0m]- %s",L.ToString(1))
	return 1
}