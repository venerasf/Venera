package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	INF = iota
	ERR
	WNG
	PNC
	EVT
	SYS
)

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
logPath: path to the log file usually `~/.venera/message.log`
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
	case INF:
		ltype = "inf"
	case ERR:
		ltype = "err"
	case WNG:
		ltype = "wng"
	case PNC:
		ltype = "pnc"
	case EVT:
		ltype = "evt"
	case SYS:
		ltype = "sys"
	default:
		ltype = "nil"
	}

	// since it is not used all the time, lets open for each use
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		return
	}
	defer f.Close()

	logMessage := fmt.Sprintf("type=%s module=%s message='%s'", ltype, module, strings.ReplaceAll(message, "'", `\'`))
	nLog := log.New(f, "", log.LstdFlags)
	nLog.Println(logMessage)
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

/*
ValidatePath checks if a path is safe and within allowed directory.
Prevents directory traversal attacks.
Returns the cleaned absolute path or error if path is unsafe.
*/
func ValidatePath(path string, allowedDir string) (string, error) {
	// Clean the paths to resolve . and ..
	cleanPath := filepath.Clean(path)
	cleanAllowedDir := filepath.Clean(allowedDir)
	
	// Convert to absolute paths
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return "", fmt.Errorf("invalid path: %w", err)
	}
	
	absAllowedDir, err := filepath.Abs(cleanAllowedDir)
	if err != nil {
		return "", fmt.Errorf("invalid allowed directory: %w", err)
	}
	
	// Check if the path starts with the allowed directory
	if !strings.HasPrefix(absPath, absAllowedDir) {
		return "", fmt.Errorf("path traversal detected: %s is outside allowed directory %s", path, allowedDir)
	}
	
	// Additional check for path separators
	if strings.Contains(path, "..") {
		return "", fmt.Errorf("path contains directory traversal sequence: %s", path)
	}
	
	return absPath, nil
}

/*
SafeJoinPath safely joins a base directory with a relative path.
Prevents directory traversal by validating the result is within base.
*/
func SafeJoinPath(base string, elem string) (string, error) {
	// Clean both paths
	cleanBase := filepath.Clean(base)
	cleanElem := filepath.Clean(elem)
	
	// Join them
	joined := filepath.Join(cleanBase, cleanElem)
	
	// Validate the result
	return ValidatePath(joined, cleanBase)
}
