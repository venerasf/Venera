// Ultils

package wlua

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	//"github.com/c-bata/go-prompt"
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
func Print(L *lua.LState) int {
	fmt.Printf("%s",L.ToString(1))
	return 1
}

// Pretty print with line ending
func PrintSuccsln(L *lua.LState) int {
	fmt.Printf("[\u001B[1;32mOK\u001B[0;0m]- %s\n",L.ToString(1))
	return 1
}
func PrintErrln(L *lua.LState) int {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s\n",L.ToString(1))
	return 1
}
func PrintInfoln(L *lua.LState) int {
	fmt.Printf("[\u001B[1;34mi\u001B[0;0m]- %s\n",L.ToString(1))
	return 1
}
func Println(L *lua.LState) int {
	fmt.Printf("%s\n",L.ToString(1))
	return 1
}

//##################################################
// Openfile and get content
func Open(L *lua.LState) int {
	p := L.ToString(1)
	cont, err := ioutil.ReadFile(p)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LString(string(cont)))
	return 1
}

// Input
func Input(L *lua.LState) int {
	p := L.ToString(1)
	if p == "" {
		p = ">> "
	}
	fmt.Print(p)
	reader := bufio.NewReader(os.Stdin)
	c,err := reader.ReadString('\n')
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	L.Push(lua.LString(c))
	return 1
}

/*func Input(L *lua.LState) int {
	p := L.ToString(1)
	if p == "" {
		p = ">> "
	}
	x := prompt.Input(p,nil)
	L.Push(lua.LString(x))
	return 1
}*/