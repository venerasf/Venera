// Ultils

package wlua

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"venera/internal/utils"

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

func LogMsg(L *lua.LState) int {
	utils.LogMsg(LuaProf.Globals["logfile"],L.ToInt(1),LuaProf.Script,L.ToString(2))
	return 1
}

//##################################################
// Open file and get content
// Restricted to only read files within the Venera scripts directory
func Open(L *lua.LState) int {
	requestedPath := L.ToString(1)
	
	// Get the allowed scripts directory from globals
	scriptsDir := LuaProf.Globals["root"]
	if scriptsDir == "" {
		L.Push(lua.LString("Error: scripts directory not configured"))
		return 1
	}
	
	// Validate that the requested path is within the scripts directory
	safePath, err := utils.ValidatePath(requestedPath, scriptsDir)
	if err != nil {
		// If validation fails, try joining with scripts directory
		safePath, err = utils.SafeJoinPath(scriptsDir, requestedPath)
		if err != nil {
			L.Push(lua.LString("Error: access denied - " + err.Error()))
			return 1
		}
	}
	
	// Additional check: ensure it's a regular file, not a directory or special file
	fileInfo, err := os.Stat(safePath)
	if err != nil {
		L.Push(lua.LString("Error: " + err.Error()))
		return 1
	}
	
	if !fileInfo.Mode().IsRegular() {
		L.Push(lua.LString("Error: not a regular file"))
		return 1
	}
	
	// Check file size to prevent reading extremely large files
	const maxFileSize = 10 * 1024 * 1024 // 10 MB limit
	if fileInfo.Size() > maxFileSize {
		L.Push(lua.LString(fmt.Sprintf("Error: file too large (max %d bytes)", maxFileSize)))
		return 1
	}
	
	cont, err := ioutil.ReadFile(safePath)
	if err != nil {
		L.Push(lua.LString("Error: " + err.Error()))
		return 1
	}
	
	L.Push(lua.LString(string(cont)))
	return 1
}

// Input
func Input(L *lua.LState) int {
	// fix tty disable raw mode
	rawoff := exec.Command("/bin/stty", "-raw", "echo")
	rawoff.Stdin = os.Stdin
	_ = rawoff.Run()
	rawoff.Wait()
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