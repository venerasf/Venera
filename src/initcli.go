package src

import (
	"strings"
	"os/exec"
	"os"
	"github.com/c-bata/go-prompt"
)


func HandleExit() {
	/*
	it is necessary to deactivate the prompt in an 
	appropriate way so as not to misconfigure the user's terminal.
	Reset tty executing stty
	disable raw mode
	*/
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
		prompt.OptionPrefix("[vnr]>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(3),
		//prompt.OptionAddKeyBind(prompt.KeyBind{})
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
	length := len(inputs)

	// Specific options \\ Commands written
	if (length == 2) {
	switch inputs[0] {
		case "use":
			return prompt.FilterHasPrefix(*ScriptSuggestions, inputs[1], true)

		case "search":
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "match", Description: "Match string"},
				{Text: "tag",   Description: "Search tags"},
			}, inputs[1], true)

		case "export":
			return prompt.FilterHasPrefix(*ScriptSuggestions, inputs[1], true)

		case "globals":
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "set", Description: "Set global variable kv"},
				{Text: "rm",   Description: "Remove global variable"},
			}, inputs[1], true)

		case "vpm":
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "search", 	Description: "Search for scripts with a pattern"},
				{Text: "install",	Description: "Install a script"},
				{Text: "sync",		Description: "Sincronize with remote repository"},
			}, inputs[1], true)
		}
	}

	// General options \\ No written commands
	promptSuggestions := []prompt.Suggest {
		{Text: "help",    	Description: "Show help menu"},
		{Text: "bash",    	Description: "Spawn a command shell"},
		{Text: "import",    Description: "Import a (edited) script"},
		{Text: "export",    Description: "Export a script (to edit)"},
		{Text: "globals",   Description: "Show global variables"},
		{Text: "vpm", 		Description: "Venera package manager"},
		{Text: "exit", 		Description: "Exit from the prompt"},
	}

 	if p.SSet { // Options only valid when there is a selected script.
		promptSuggestions = append(promptSuggestions,
			prompt.Suggest {Text: "set",     Description: "Set value for a var"},
			prompt.Suggest {Text: "run",     Description: "Run a script/module"},
			prompt.Suggest {Text: "back",    Description: "Exit module/script"},
			prompt.Suggest {Text: "options", Description: "Show variables of script/module"},
			prompt.Suggest {Text: "lua",     Description: "Run Lua code in running mod"},
			prompt.Suggest {Text: "info",    Description: "Info/metadata about script/module"},
			prompt.Suggest {Text: "reload",    Description: "Reloads the selected script"},
		)
	} else {	// Options only valid when there is no selected script.
		promptSuggestions = append(promptSuggestions,
			prompt.Suggest {Text: "search", Description: "Search script/module"},
			prompt.Suggest {Text: "use",    Description: "Load a script/module"},
		) 
	}

	return prompt.FilterHasPrefix(promptSuggestions, inputs[0], true)
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
