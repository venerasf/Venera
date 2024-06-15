// This file for working with modules/script
// like list modules, search for a module, etc...

// Functions must have SC prefix
package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"venera/src/wlua"
	"venera/src/utils"

	"github.com/c-bata/go-prompt"
	"github.com/cheynewallace/tabby"
)

var ScriptSuggestions *[]prompt.Suggest // script list with descriptions
var SCTAG []ScriptTAGInfo               // script list with tags and infos,
										// it will be in memory for later use.

// Load all paths, get metadata INFO and tags
// TODO: The regex can be better
func (p Profile) SCLoadScripts() {
	re := regexp.MustCompile(`METADATA(\s)*=(\s)*\{((.|\n)*)INFO(\s)*=(\s)*\[\[((.|\n)*?)\]\]((.|\n)*)\}`)
	//rea := *re
	paths := p.SCGetPath()

	aux := []prompt.Suggest{}
	for _, file := range paths {
		info := SCExtractINFO(file, re)
		tags := wlua.ScriptGetTags(file)
		SCTAG = append(SCTAG, ScriptTAGInfo{file, tags, info})
		aux = append(aux, prompt.Suggest{Text: file, Description: info})
	}
	ScriptSuggestions = &aux
}

func (p Profile) SCGetPath() []string {
	//root := p.BPath
	root := p.Globals["root"]
	filePath := []string{} // List of file paths

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Validate file
		m, _ := regexp.MatchString(`.*\.(lua|vnr)$`, path)
		if m {
			filePath = append(filePath, path)
		}
		return nil
	})
	if err != nil {
		utils.PrintErr(err.Error())
	}
	return filePath
}

// Use for seaarch functions
// TODO: Use `strings.ToLower()` to match strings without case sensitive
func (p Profile) SCListScripts(key []string) {
	pathList := p.SCGetPath()
	t := tabby.New()

	if len(key) == 1 {
		// TODO: Put a limit
		t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
		aux := SCTAG
		for i := range pathList {
			t.AddLine(i+1, aux[i].Path, aux[i].Info, JoinTgs(aux[i].Tag))
		}
		print("\n")
		t.Print()
		print("\n")

	} else if (key[1] == "all" || key[1] == "a") && len(key) >= 2 {
		t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
		aux := SCTAG
		for i := range aux {
			t.AddLine(i+1, aux[i].Path, aux[i].Info, JoinTgs(aux[i].Tag))
		}
		print("\n")
		t.Print()
		print("\n")

		// List match just path
	} else if (key[1] == "match:path" || key[1] == "m:path" || key[1] == "m:p" || key[1] == "match:p") && len(key) >= 3 {
		t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
		aux := SCTAG
		for i := range aux {
			if strings.Contains(strings.ToLower(aux[i].Path), strings.ToLower(key[2])) {
				t.AddLine(i+1, aux[i].Path, aux[i].Info, JoinTgs(aux[i].Tag))
			}
		}
		print("\n")
		t.Print()
		print("\n")

		// List match description
	} else if (key[1] == "match:description" || key[1] == "m:description" || key[1] == "m:d" || key[1] == "match:d") && len(key) >= 3 {
		t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
		aux := SCTAG
		for i := range aux {
			if strings.Contains(aux[i].Info, key[2]) {
				t.AddLine(i+1, aux[i].Path, aux[i].Info, JoinTgs(aux[i].Tag))
			}
		}
		print("\n")
		t.Print()
		print("\n")

		// Match anything, path and description
	} else if (key[1] == "match" || key[1] == "m") && len(key) >= 3 {
		t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
		aux := SCTAG
		for i := range aux {
			if strings.Contains(aux[i].Path, key[2]) || strings.Contains(aux[i].Info, key[2]) {
				t.AddLine(i+1, aux[i].Path, aux[i].Info, JoinTgs(aux[i].Tag))
			}
		}
		print("\n")
		t.Print()
		print("\n")

	} else if (key[1] == "tag" || key[1] == "t") && len(key) == 2 {
		print("\n")
		fmt.Println("AVAILABLE TAGS:\n", TagsJoinALL())
		print("\n")
	} else if (key[1] == "tag" || key[1] == "t") && len(key) >= 3 {
		t.AddHeader("COUNT", "PATH", "INFO", "TAG")
		aux := SCTAG
		for x, tag := range aux {
			for i := range tag.Tag {
				for _, j := range key[2:] {
					if strings.Contains(strings.ToLower(tag.Tag[i]), strings.ToLower(j)) {
						tags := strings.Join(tag.Tag, ", ")
						if len(tags) < 20 {
							t.AddLine(x+1, tag.Path, tag.Info, tags)
						} else {
							t.AddLine(x+1, tag.Path, tag.Info, tags[:20]+"...")
						}
						break
					}
				}
			}
		}
		print("\n")
		t.Print()
		print("\n")

	} else {
		// TODO: Put a limit
		t.AddHeader("COUNT", "PATH", "DESCRIPTION")
		aux := *ScriptSuggestions
		for i := range aux {
			t.AddLine(i+1, aux[i].Text, aux[i].Description)
		}
		print("\n")
		t.Print()
		print("\n")
	}
}

// // Extract INFO from script based on the regex passed (in SCLoadScripts())
func SCExtractINFO(path string, re *regexp.Regexp) string {
	const l = 25
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "Nil info..."
	}
	match := re.FindStringSubmatch(string(content))
	if len(match) < 7 {
		return "Nil info..."
	}
	if len(match[7]) >= l-1 {
		return match[7][:l] + "..."
	} else {
		return match[7]
	}
}

// Create a string with ths in fixed length
func JoinTgs(t []string) string {
	const l = 25
	aux := strings.Join(t, ",")
	if len(aux) >= l-1 {
		return aux[:l]
	} else {
		return aux
	}
}

// Create a string with all tags from all scripts
func TagsJoinALL() string {
	t := []string{}
	for _, j := range SCTAG {
		t = append(t, j.Tag...)
	}
	sort.Strings(t)

	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range t {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return strings.Join(list, ",")
}


func (p Profile)SCInfoForChaining() {
	m := make(map[string]bool)
	utils.PrintSuccs("Listing loaded scripts.")
	for i := range(p.Scriptslist) {
		if !m[p.Scriptslist[i]] {
			m[p.Scriptslist[i]] = true
			fmt.Println("- "+p.Scriptslist[i])
		}
	}
}