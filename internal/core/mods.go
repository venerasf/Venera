// This file for working with modules/script
// like list modules, search for a module, etc...

// Functions must have SC prefix
package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"venera/internal/constants"
	"venera/internal/types"
	"venera/internal/utils"
	"venera/internal/wlua"

	"github.com/c-bata/go-prompt"
	"github.com/cheynewallace/tabby"
)

var ScriptSuggestions *[]prompt.Suggest // script list with descriptions
var SCTAG []types.ScriptTAGInfo         // script list with tags and infos,
// it will be in memory for later use.

// Load all paths, get metadata INFO and tags
// TODO: The regex can be better
func SCLoadScripts(p types.Profile) {
	re := regexp.MustCompile(`METADATA(\s)*=(\s)*\{((.|\n)*)INFO(\s)*=(\s)*\[\[((.|\n)*?)\]\]((.|\n)*)\}`)
	//rea := *re
	paths := SCGetPath(p)

	aux := []prompt.Suggest{}
	for _, file := range paths {
		info := SCExtractMetadataINFO(file, re)
		tags := wlua.ScriptGetTags(file)
		SCTAG = append(SCTAG, types.ScriptTAGInfo{file, tags, info})
		aux = append(aux, prompt.Suggest{Text: file, Description: info})
	}
	ScriptSuggestions = &aux
}

func SCGetPath(p types.Profile) []string {
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

// Use for search functions
// TODO: Use `strings.ToLower()` to match strings without case sensitive
func SCListScripts(p types.Profile, key []string) {
	if len(key) == 1 {
		printAllScripts()
	} else if (key[1] == "all" || key[1] == "a") && len(key) >= 2 {
		printAllScripts()
	} else if (key[1] == "match:path" || key[1] == "m:path" || key[1] == "m:p" || key[1] == "match:p") && len(key) >= 3 {
		printMatchingScripts(func(sc types.ScriptTAGInfo) bool {
			return strings.Contains(strings.ToLower(sc.Path), strings.ToLower(key[2]))
		})
	} else if (key[1] == "match:description" || key[1] == "m:description" || key[1] == "m:d" || key[1] == "match:d") && len(key) >= 3 {
		printMatchingScripts(func(sc types.ScriptTAGInfo) bool {
			return strings.Contains(sc.Info, key[2])
		})
	} else if (key[1] == "match" || key[1] == "m") && len(key) >= 3 {
		printMatchingScripts(func(sc types.ScriptTAGInfo) bool {
			return strings.Contains(sc.Path, key[2]) || strings.Contains(sc.Info, key[2])
		})
	} else if (key[1] == "tag" || key[1] == "t") && len(key) == 2 {
		print("\n")
		fmt.Println("AVAILABLE TAGS:\n", TagsJoinALL())
		print("\n")
	} else if (key[1] == "tag" || key[1] == "t") && len(key) >= 3 {
		printMatchingTags(key[2:])
	} else {
		printScriptSuggestions()
	}
}

// printAllScripts displays all scripts in a table
func printAllScripts() {
	t := tabby.New()
	t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
	for i, sc := range SCTAG {
		t.AddLine(i+1, sc.Path, sc.Info, JoinTgs(sc.Tag))
	}
	print("\n")
	t.Print()
	print("\n")
}

// printMatchingScripts displays scripts matching a filter function
func printMatchingScripts(matchFunc func(types.ScriptTAGInfo) bool) {
	t := tabby.New()
	t.AddHeader("COUNT", "PATH", "DESCRIPTION", "TAGS")
	for i, sc := range SCTAG {
		if matchFunc(sc) {
			t.AddLine(i+1, sc.Path, sc.Info, JoinTgs(sc.Tag))
		}
	}
	print("\n")
	t.Print()
	print("\n")
}

// printMatchingTags displays scripts matching specific tags
func printMatchingTags(searchTags []string) {
	t := tabby.New()
	t.AddHeader("COUNT", "PATH", "INFO", "TAG")
	for x, tag := range SCTAG {
		for i := range tag.Tag {
			for _, j := range searchTags {
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
}

// printScriptSuggestions displays script suggestions in a table
func printScriptSuggestions() {
	t := tabby.New()
	t.AddHeader("COUNT", "PATH", "DESCRIPTION")
	if ScriptSuggestions != nil {
		for i, suggestion := range *ScriptSuggestions {
			t.AddLine(i+1, suggestion.Text, suggestion.Description)
		}
	}
	print("\n")
	t.Print()
	print("\n")
}

// // Extract INFO from script based on the regex passed (in SCLoadScripts())
func SCExtractMetadataINFO(path string, re *regexp.Regexp) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "Nil info..."
	}
	match := re.FindStringSubmatch(string(content))
	if len(match) < 7 {
		return "Nil info..."
	}
	if len(match[7]) >= constants.MaxInfoDisplayLength-1 {
		return match[7][:constants.MaxInfoDisplayLength] + "..."
	} else {
		return match[7]
	}
}

// Create a string with ths in fixed length
func JoinTgs(t []string) string {
	aux := strings.Join(t, ",")
	if len(aux) >= constants.MaxInfoDisplayLength-1 {
		return aux[:constants.MaxInfoDisplayLength]
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

func SCInfoForChaining(p types.Profile) {
	m := make(map[string]bool)
	utils.PrintSuccs("Listing loaded scripts.")
	for i := range p.Scriptslist {
		if !m[p.Scriptslist[i]] {
			m[p.Scriptslist[i]] = true
			fmt.Println("- " + p.Scriptslist[i])
		}
	}
}
