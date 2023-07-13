package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func GetBash() {
	cmd := exec.Command("bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// Types of pretty printing
func PrintSuccs(s string) {
	fmt.Printf("[\u001B[1;32mOK\u001B[0;0m]- %s\n",s)
}
func PrintErr(s string) {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s\n",s)
}
func PrintAlert(s string) {
	fmt.Printf("[\u001B[1;31m!\u001B[0;0m]- %s\n",s)
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
	f, err := os.OpenFile(logPath,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return		
	}

	logMessage := fmt.Sprintf("type=%s module=%s message=%s",
		ltype, module,message)
	nLog := log.New(f,"", log.LstdFlags)
	nLog.Println(logMessage)
	if err != nil {
		panic(err.Error())
	}
	f.Close()
}