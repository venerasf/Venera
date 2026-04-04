package core

import (
	"os"
	"os/exec"
	"strings"
	"venera/internal/types"

	"github.com/c-bata/go-prompt"
)

// Live prefix for prompt configs
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

/*
Only used for the types.Profile to be passed as closure for
prompt.New()
*/
type ProfAux struct {
	p *types.Profile
}

func exitFunct(f *prompt.Buffer) {
	HandleExit()
	os.Exit(0)
}

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

// Cleanup performs resource cleanup before shutdown
func Cleanup(p *types.Profile) {
	if p != nil && p.Database != nil {
		p.Database.Close()
	}
}

/*
Initialize command prompt.
p.Execute() from cmds.go has all commands.
*/
func InitCLI(p *types.Profile) {
	defer Cleanup(p)

	SCLoadScripts(*p) // Load scripts from `script` root dir to memory.
	Banner()
	SCGetPath(*p)

	paux := new(ProfAux)
	paux.p = p

	defer HandleExit()
	prom := prompt.New(
		paux.Execute,
		paux.completer,
		prompt.OptionPrefix("[vnr]>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionCompletionOnDown(),
		prompt.OptionMaxSuggestion(3),
		prompt.OptionAddKeyBind(prompt.KeyBind{prompt.ControlQ, exitFunct}),
	)
	prom.Run()
}

// Function to change prompt string dynamicaly
func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

// Load suggestions
func (paux *ProfAux) completer(d prompt.Document) []prompt.Suggest {
	inputs := strings.Split(d.TextBeforeCursor(), " ")
	length := len(inputs)

	// Specific options - commands with 2 tokens
	if length == 2 {
		return paux.completeSecondArg(inputs)
	}

	// VPM subcommands - commands with 3 tokens
	if length == 3 {
		return paux.completeThirdArg(inputs)
	}

	// General options - first argument completion
	return paux.completeFirstArg(inputs)
}

// completeFirstArg returns suggestions for the first argument
func (paux *ProfAux) completeFirstArg(inputs []string) []prompt.Suggest {
	promptSuggestions := []prompt.Suggest{
		{Text: "help", Description: "Show help menu"},
		{Text: "import", Description: "Import a (edited) script"},
		{Text: "export", Description: "Export a script (to edit)"},
		{Text: "globals", Description: "Show global variables"},
		{Text: "reload", Description: "Reload a memory resource"},
		{Text: "vpm", Description: "Venera package manager"},
		{Text: "exit", Description: "Exit from the prompt"},
	}

	if paux.p.SSet { // Options only valid when there is a selected script
		promptSuggestions = append(promptSuggestions,
			prompt.Suggest{Text: "set", Description: "Set value for a var"},
			prompt.Suggest{Text: "run", Description: "Run a script/module"},
			prompt.Suggest{Text: "back", Description: "Exit module/script"},
			prompt.Suggest{Text: "options", Description: "Show variables of script/module"},
			prompt.Suggest{Text: "lua", Description: "Run Lua code in running mod"},
			prompt.Suggest{Text: "info", Description: "Info/metadata about script/module"},
		)
	} else { // Options only valid when there is no selected script
		promptSuggestions = append(promptSuggestions,
			prompt.Suggest{Text: "search", Description: "Search script/module"},
			prompt.Suggest{Text: "use", Description: "Load a script/module"},
		)
	}

	return prompt.FilterHasPrefix(promptSuggestions, inputs[0], true)
}

// completeSecondArg returns suggestions for second arguments based on the command
func (paux *ProfAux) completeSecondArg(inputs []string) []prompt.Suggest {
	switch inputs[0] {
	case "use", "export":
		if ScriptSuggestions != nil {
			return prompt.FilterHasPrefix(*ScriptSuggestions, inputs[1], true)
		}
		return []prompt.Suggest{}

	case "search":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "match", Description: "Match string"},
			{Text: "tag", Description: "Search tags"},
		}, inputs[1], true)

	case "globals":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "set", Description: "Set global variable kv"},
			{Text: "rm", Description: "Remove global variable"},
		}, inputs[1], true)

	case "vpm":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "search", Description: "Search for scripts with a pattern"},
			{Text: "install", Description: "Install a script"},
			{Text: "sync", Description: "Synchronize with remote repository"},
			{Text: "verify", Description: "Verify package signature"},
			{Text: "key", Description: "Manage keys"},
		}, inputs[1], true)

	case "reload":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "root", Description: "Reload root directory"},
			{Text: "script", Description: "Reload script in memory"},
		}, inputs[1], true)

	case "help":
		return prompt.FilterHasPrefix(HelpSugg, inputs[1], true)
	}

	return []prompt.Suggest{}
}

// completeThirdArg returns suggestions for third arguments (VPM subcommands)
func (paux *ProfAux) completeThirdArg(inputs []string) []prompt.Suggest {
	switch inputs[1] {
	case "key":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "show", Description: "Show keys"},
			{Text: "del", Description: "Delete key"},
			{Text: "add", Description: "Add key from file"},
		}, inputs[2], true)

	case "search":
		return prompt.FilterHasPrefix([]prompt.Suggest{
			{Text: "all", Description: "Return all packages"},
		}, inputs[2], true)
	}

	return []prompt.Suggest{}
}

/*

{"cmds":[{
	"name":"network",
	"description":"Do network scan",
	"subcmd":[
		{"name": "ping", "description":"Do ping", "subcmd":null},
		{"name": "scan", "description":"Do network scan", "subcmd":[{"name": "tcp", "description":"choose tcp scan", "subcmd":null}]}
	]}
}]}
*/
