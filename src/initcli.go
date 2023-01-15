package src

import (
	"strings"
	"github.com/c-bata/go-prompt"
)

func (p *Profile)InitCLI() {
	SCLoadScripts()
	Banner()
	SCGetPath()
	prom := prompt.New(
		p.Execute,
		p.completer,
		prompt.OptionPrefix("[*]>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(2),
	)
	prom.Run()
}


// Function to change prompt string dynamicaly
func changeLivePrefix() (string,bool) {
	return LivePrefixState.LivePrefix,LivePrefixState.IsEnable
}

// Suggentions
func (p *Profile)completer(d prompt.Document) []prompt.Suggest {
	inputs := strings.Split(d.TextBeforeCursor(), " ")
	switch inputs[0] {
	case "use":
			aux := *ScriptSuggentions
			return aux
	
	case "s":
			return []prompt.Suggest{
				{Text: "match", 	Description: "Match string"},
			}
	
	
	case "serach":
			return []prompt.Suggest{
				{Text: "match", 	Description: "Match string"},
			}
}

	// General options
	// If script setted, show script options
	if p.SSet {
		return []prompt.Suggest{
			{Text: "help", 	Description: "Show help menu"},
			{Text: "search", 	Description: "Search script/module"},
			// Inside script/module options
			{Text: "back", 		Description:"Exit module/script"},
			{Text: "set",		Description:"Set value for a ver"},
			{Text: "options", 	Description: "Show variables of script/module"},
			{Text: "info", 		Description: "Info/metadata about script/module"},
			{Text: "run", 		Description: "Run a script/module"},
			{Text: "lua", 		Description: "Run Lua code in running mod"},
		}
	} else {
		return []prompt.Suggest{
			{Text: "help", 	Description: "Show help menu"},
			{Text: "use", 	Description: "Load a script/module"},
			{Text: "search", 	Description: "Search script/module"},
		}
	}

/*	
	return prompt.FilterContains(
		[]prompt.Suggest{
			// General options
			{Text: "help", 	Description: "Show help menu"},
			{Text: "use", 	Description: "Load a script/module"},
			{Text: "bash", 	Description: "Spawn shell"},

			// Inside script/module options
			{Text:"back", 		Description:"Exit module/script"},
			{Text:"set",		Description:"Set value for a ver"},
			{Text: "options", 	Description: "Show variables of script/module"},
			{Text: "info", 		Description: "Info/metadata about script/module"},
			{Text: "run", 		Description: "Run a script/module"},
			{Text: "lua", 		Description: "Run Lua code in running mod"},
		}, d.GetWordBeforeCursor(),true)
*/
	}