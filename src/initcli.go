package src

import (
	"github.com/c-bata/go-prompt"
)

func (p *Profile)InitCLI() {
	prom := prompt.New(
		p.Execute,
		completer,
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
func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterContains(
		[]prompt.Suggest{
			{Text: "help", Description: "Show help menu"},
			{Text: "bash", Description: "Spawn shell"},

			{Text: "use", Description: "Load a script/module"},
			{Text: "options", Description: "Show variables of script/module"},
			{Text: "info", Description: "Info/metadata about script/module"},
			{Text: "run", Description: "Run a script/module"},
		}, d.GetWordBeforeCursor(),true)
}