// https://github.com/c-bata/go-prompt/issues/18#issuecomment-1003841774

package utils

import (
	"syscall"
	_ "unsafe" // Required for go:linkname
	"fmt"
	"github.com/c-bata/go-prompt"
)

// Map the local variable "consoleWriter" to the one of go-prompt
//go:linkname consoleWriter github.com/c-bata/go-prompt.consoleWriter
var consoleWriter prompt.ConsoleWriter

func AsyncWriteStdout(a...any) {
	consoleWriter.EraseLine() // Erase current line
	consoleWriter.EraseDown() // Required to remove the completions menu
	consoleWriter.WriteRawStr("\r" + fmt.Sprint(a...)) // 'r' to go back to the start of line
	syscall.Kill(syscall.Getpid(), syscall.SIGWINCH) // Required to force the re-render of the prompt
}