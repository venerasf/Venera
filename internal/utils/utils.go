package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// useless func that prob will f*** someone as a privesc vector :)
func GetBash() {
	cmd := exec.Command("bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// Types of pretty printing
func PrintSuccs(a ...any) {
	fmt.Printf("[\u001B[1;32mOK\u001B[0;0m]- %s\n", fmt.Sprint(a...))
}

func PrintErr(a ...any) {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s\n", fmt.Sprint(a...))
}

func PrintAlert(a ...any) {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s\n", fmt.Sprint(a...))
}

func PrintLn(a ...any) {
	fmt.Print(fmt.Sprint(a...), "\n")
}

/*
PrintPanic will print the message and exit with status code 1
*/
func PrintPanic(a ...any) {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s\n", fmt.Sprint(a...))
	os.Exit(1)
}

/*
logPath: path to the log file usualy `~venera/message.log`
tp: type of log

	0 - inf = information
	1 - err = error
	2 - wng = warning
	3 - pnc = panic
	4 - evt = event
	5 - sys = system
	default - nil

module: the module that is logging like `core` for venera
or the path if the a script is logging.
message: the message
*/
func LogMsg(logPath string, tp int, module string, message string) {
	var ltype string
	switch tp {
	case 0:
		ltype = "inf"
	case 1:
		ltype = "err"
	case 2:
		ltype = "wng"
	case 3:
		ltype = "pnc"
	case 4:
		ltype = "evt"
	case 5:
		ltype = "sys"
	default:
		ltype = "nil"
	}

	// since it is not used all the time, lets open for each use
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	logMessage := fmt.Sprintf("type=%s module=%s message=%s",
		ltype, module, message)
	nLog := log.New(f, "", log.LstdFlags)
	nLog.Println(logMessage)
	if err != nil {
		panic(err.Error())
	}
	f.Close()
}

/*
We handle script path like /home/farinap/.venera/scripts/cms/wp_user_enum.lua
It is big and kinda useless due the rootPath (base path) is always the same.
It must be process to be just cms/wp_user_enum.lua.
*/
func HideBasePath(rootePath, scrptName string) string {
	// return scrptName[len(rootePath):]
	return strings.TrimPrefix(scrptName, rootePath)
}

/*
Remove lua extension from path
from cms/wp_user_enum.lua
to cms/wp_user_enum
*/
func HideLuaExtension(scrptName string) string {
	return strings.TrimSuffix(scrptName, ".lua")
}
