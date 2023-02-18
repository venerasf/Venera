package src

import (
	"strings"
	"os/exec"
	"os"
	"github.com/c-bata/go-prompt"
)

func HandleExit() {
	// Reset tty executing stty
	// disable raw mode
	rawoff := exec.Command("/bin/stty", "-raw", "echo")
	rawoff.Stdin = os.Stdin
	_ = rawoff.Run()
	rawoff.Wait()
}

func (p *Profile) InitCLI() {
	p.SCLoadScripts()
	Banner()
	p.SCGetPath()

	defer HandleExit()
	prom := prompt.New(
		p.Execute,
		p.completer,
		prompt.OptionPrefix("[*]>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(3),
	)
	prom.Run()
}

// Function to change prompt string dynamicaly
func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

// Load suggestions
func (p *Profile) completer(d prompt.Document) []prompt.Suggest {
	//inputs := strings.Split(d.CurrentLine(), " ")
	inputs := strings.Split(d.TextBeforeCursor(), " ")
	switch inputs[0] {
	case "use":
		// TODO filtering suggestions after use
		if len(inputs) == 2 {
			aux := prompt.FilterHasPrefix(*ScriptSuggestions, inputs[1], true)
			return aux
		}
		aux := *ScriptSuggestions
		return aux

	case "search":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "match", Description: "Match string"},
			{Text: "tag",   Description: "Search tags"},
		}, inputs[1], true)

	case "export":
		aux := prompt.FilterHasPrefix(*ScriptSuggestions, inputs[1], true)
		return aux

	case "set":
		return []prompt.Suggest{
			{Text: "global", Description: "Set a global variable"},
		}
	}

	// General options
	// If script setted then show script options.
	if p.SSet {
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "set",     Description: "Set value for a var"},
			{Text: "run",     Description: "Run a script/module"},
			{Text: "help",    Description: "Show help menu"},
			{Text: "search",  Description: "Search script/module"},
			// Inside script/module options
			{Text: "back",    Description: "Exit module/script"},
			{Text: "options", Description: "Show variables of script/module"},
			{Text: "info",    Description: "Info/metadata about script/module"},
			{Text: "lua",     Description: "Run Lua code in running mod"},
			{Text: "exit",    Description: "Exit from prompt"},
		}, inputs[0], true)
	} else {
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "help",   Description: "Show help menu"},
			{Text: "use",    Description: "Load a script/module"},
			{Text: "search", Description: "Search script/module"},
			{Text: "import", Description: "Import a script"},
			{Text: "export", Description: "Export a script"},
			{Text: "exit",   Description: "Exit from prompt"},
		}, inputs[0], true)
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
