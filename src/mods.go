// This file is to work with modules/script
// like list modules, search for a module, etc...

// Function must initiate with SC
package src

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/cheynewallace/tabby"
	//"strings"
)

var ScriptSuggentions *[]prompt.Suggest

func (p Profile)SCGetPath() []string {
	root := p.BPath
	filePath := []string{} // List of file paths

	err := filepath.Walk(root,func(path string, info os.FileInfo, err error) error {
		// Validate file
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

func (p Profile)SCLoadScripts() {
	re := regexp.MustCompile(`Metadata(\s)*=(\s)*\{((.|\n)*)INFO(\s)*=(\s)*\[\[((.|\n)*?)\]\]((.|\n)*)\}`)
	//rea := *re
	paths := p.SCGetPath()
	
	aux := []prompt.Suggest{}
	for _,file := range(paths) {
		aux = append(aux, prompt.Suggest{Text: file,Description: SCExtractINFO(file,re)})
	}
	ScriptSuggentions = &aux
}

// Use for seaarch functions
func (p Profile)SCListScripts(key []string) {
	pathList := p.SCGetPath()
	t := tabby.New()
	
	if len(key) == 1 {
		// TODO: Put a limit
		t.AddHeader("COUNT","PATH","DESCRIPTION")
		aux := *ScriptSuggentions
		for i := range(pathList) {
			t.AddLine(i+1,aux[i].Text,aux[i].Description)
		}
		print("\n")
		t.Print()
		print("\n")


	} else if (key[1] == "all" || key[1] == "a") && len(key) >= 2 {
		t.AddHeader("COUNT","PATH","DESCRIPTION")
		aux := *ScriptSuggentions
		for i := range(aux) {
			t.AddLine(i+1,aux[i].Text,aux[i].Description)
		}
		print("\n")
		t.Print()
		print("\n")


	// List match just path
	} else if (key[1] == "match:path" || key[1] == "m:path" || key[1] == "m:p" || key[1] == "match:p") && len(key) >= 3 {
		t.AddHeader("COUNT","PATH","DESCRIPTION")
		aux := *ScriptSuggentions
		for i := range(aux) {
			if strings.Contains(aux[i].Text,key[2]) {
				t.AddLine(i+1,aux[i].Text,aux[i].Description)
			}
		}
		print("\n")
		t.Print()
		print("\n")

	// List match description
	} else if (key[1] == "match:description" || key[1] == "m:description" || key[1] == "m:d" || key[1] == "match:d") && len(key) >= 3 {
		t.AddHeader("COUNT","PATH","DESCRIPTION")
		aux := *ScriptSuggentions
		for i := range(aux) {
			if strings.Contains(aux[i].Description,key[2]) {
				t.AddLine(i+1,aux[i].Text,aux[i].Description)
			}
		}
		print("\n")
		t.Print()
		print("\n")

	// Match anything, path and description
	} else if (key[1] == "match" || key[1] == "m") && len(key) >= 3 {
			t.AddHeader("COUNT","PATH","DESCRIPTION")
			aux := *ScriptSuggentions
			for i := range(aux) {
				if strings.Contains(aux[i].Text,key[2]) || strings.Contains(aux[i].Description,key[2]) {
					t.AddLine(i+1,aux[i].Text,aux[i].Description)
				}
			}
			print("\n")
			t.Print()
			print("\n")

			
	} else {
		// TODO: Put a limit
		t.AddHeader("COUNT","PATH","DESCRIPTION")
		aux := *ScriptSuggentions
		for i := range(aux) {
			t.AddLine(i+1,aux[i].Text,aux[i].Description)
		}
		print("\n")
		t.Print()
		print("\n")
	}
}


//// Extract INFO from script
func SCExtractINFO(path string, re *regexp.Regexp) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "Nil info..."
	}
	match := re.FindStringSubmatch(string(content))
	if len(match) < 7 {
		return "Nil info..."
	}
	if len(match[7]) >= 15 {
		return match[7][:20]+"..."
	} else {
		return match[7]
	}
}