// This file is to work with modules/script
// like list modules, search for a module, etc...

// Function must initiate with SC
package src

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/c-bata/go-prompt"
	//"strings"
)

var ScriptSuggentions *[]prompt.Suggest

func SCGetPath() []string {
	root := "scripts/"
	filePath := []string{} // List of file paths

	err := filepath.Walk(root,func(path string, info os.FileInfo, err error) error {
		// Validate filw
		m, _ := regexp.MatchString(`.*\.(lua|vnr)$`,path)
		if m {
			filePath = append(filePath, path)
		}
		return nil
	})
	if err != nil {
		PrintErr(err.Error())
	}
	return filePath
}

func SCLoadScripts() {
	paths := SCGetPath()
	
	aux := []prompt.Suggest{}
	for _,file := range(paths) {
		aux = append(aux, prompt.Suggest{Text: file,Description: "null"})
	}
	ScriptSuggentions = &aux
}