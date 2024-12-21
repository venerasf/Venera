package core

import (
	"strings"

	"github.com/c-bata/go-prompt"

	"venera/internal/utils"
	"venera/internal/types"
)

var HelpSugg = []prompt.Suggest{}

// Will keep it global for now.
var Mapping = make(map[string]*types.Command)

func init() {
	loadFunctions()
	for k, v := range Mapping {
		HelpSugg = append(HelpSugg, prompt.Suggest{
			Text:        k,
			Description: v.Desc,
		})
	}
}

func (paux *ProfAux)Execute(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmds := strings.Split(cmd, " ")
	length := len(cmds)

	// Validates length
	if length == 0 {
		utils.PrintErr("Too few arguments. Try `help cmd`.")
		return
	}

	/*
		The mapping must return the pointer to a struct that describes the command x (where x = cmds[0]).
		We must certify if the command really exists by assuming nil if the command was't assigned.
		The command is called passing the array of arguments that are passed through the command line
		and the profile.
	*/
	cmdPtr := Mapping[cmds[0]]
	if cmdPtr != nil {
		functionP := *cmdPtr
		functionP.Call(cmds, paux.p) // paux.p = types.Profile
	} else {
		utils.PrintErr("Not a valid command or missing a selected script. Type `help`.")
	}
}



/*
Register the default commands

	type Command struct {
		Call 	func([]string, *Profile) int // Callback entrypoint
		Usage 	func([]string) // help function callback
		Desc 	string // hight level description.
		Promp 	[][]string // Prompt help and auto-complete for subcmds
	}
*/
func loadFunctions() {
	Mapping["help"] = &types.Command{
		Call:  runHelp,
		Usage: usageHelp,
		Desc:  "Show help menu.",
		Promp: nil,
	}


	Mapping["import"] = &types.Command{
		Call:  runImport,
		Usage: nil,
		Desc:  "Import a (edited) script.",
		Promp: nil,
	}

	Mapping["export"] = &types.Command{
		Call:  runExport,
		Usage: nil,
		Desc:  "Export a script.",
		Promp: nil,
	}

	Mapping["globals"] = &types.Command{
		Call:  runManageGlobals,
		Usage: usageGlobal,
		Desc:  "Manage global variables.",
		Promp: nil,
	}

	Mapping["run"] = &types.Command{
		Call:  runRunScript,
		Usage: nil,
		Desc:  "Execute the script.",
		Promp: nil,
	}

	Mapping["exit"] = &types.Command{
		Call:  runExit,
		Usage: nil,
		Desc:  "Properly exit the tool.",
		Promp: nil,
	}

	Mapping["reload"] = &types.Command{
		Call:  runReload,
		Usage: usageReload,
		Desc:  "Reload (root|script).",
		Promp: nil,
	}

	Mapping["search"] = &types.Command{
		Call:  runSearch,
		Usage: usageSearch,
		Desc:  "Search a script using patterns.",
		Promp: nil,
	}

	Mapping["info"] = &types.Command{
		Call:  runInfo,
		Usage: nil,
		Desc:  "Information regarding the loaded script.",
		Promp: nil,
	}

	Mapping["options"] = &types.Command{
		Call:  runOptions,
		Usage: nil,
		Desc:  "Show configurable variables for loaded script.",
		Promp: nil,
	}

	Mapping["lua"] = &types.Command{
		Call:  runLua,
		Usage: nil,
		Desc:  "Execute inline lua commands.",
		Promp: nil,
	}

	Mapping["back"] = &types.Command{
		Call:  runBack,
		Usage: nil,
		Desc:  "Free the script.",
		Promp: nil,
	}

	Mapping["use"] = &types.Command{
		Call:  runUse,
		Usage: usageUse,
		Desc:  "Use a script.",
		Promp: nil,
	}

	Mapping["banner"] = &types.Command{
		Call:  runBanner,
		Usage: nil,
		Desc:  "Show banner.",
		Promp: nil,
	}

	Mapping["vpm"] = &types.Command{
		Call:  runVPM,
		Usage: usageVPM,
		Desc:  "Use Venera package manager.",
		Promp: nil,
	}

	Mapping["set"] = &types.Command{
		Call:  runSet,
		Usage: usageSet,
		Desc:  "Set a variable",
		Promp: nil,
	}
}
