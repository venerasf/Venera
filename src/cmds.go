package src

import (
	"strings"

	"venera/src/utils"

	"github.com/c-bata/go-prompt"
)

var HelpSugg = []prompt.Suggest{}

// Will keep it global for now.
var Mapping = make(map[string]*Command)

func init() {
	loadFunctions()
	for k,v := range Mapping {
		HelpSugg = append(HelpSugg, prompt.Suggest{
			Text: k,
			Description: v.Desc,
		})
	}
}

func (profile *Profile) Execute(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmds := strings.Split(cmd, " ")
	length := len(cmds)

	// Validates length
	if length == 0 {
		utils.PrintErr("Too few arguments. Try `help cmd`.")
		return
	}

	switch cmds[0] {
	case "help":
		if length >= 2 {
			/*
				If there are more arguments the callback usage function is called.
				Search for the structure that describes the command that `help` is being called.
				Certify that command exists comparing the pointer with nil.
				Run the `usage` function.
			*/
			cmdPtr := Mapping[cmds[1]]
			if cmdPtr != nil {
				functionP := *cmdPtr
				if functionP.Usage == nil {
					utils.PrintAlert("The command does not have a valid usage callback.")
					utils.PrintLn(functionP.Desc)
				} else {
					functionP.Usage(cmds)
				}
			} else {
				utils.PrintAlert("The command does not have help menu.")
			}
		} else {
			CmdHelp()
		}
	default:
		/*
			The mapping must return the pointer to a struct that describes the command x (where x = cmds[0]).
			We must certify if the command really exists by assuming nil if the command was't assigned.
			The command is called passing the array of arguments that are passed through the command line
			and the profile.
		*/
		cmdPtr := Mapping[cmds[0]]
		if cmdPtr != nil {
			functionP := *cmdPtr
			functionP.Call(cmds, profile)
		} else {
			utils.PrintErr("Not a valid command or missing a selected script. Type `help`.")
		}
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
	Mapping["help"] = &Command{
		Call:  runHelp,
		Usage: usageHelp,
		Desc:  "Show help menu.",
		Promp: nil,
	}

	Mapping["bash"] = &Command{
		Call:  runBash,
		Usage: nil,
		Desc:  "Spawns a shell.",
		Promp: nil,
	}

	Mapping["import"] = &Command{
		Call:  runImport,
		Usage: nil,
		Desc:  "Import a (edited) script.",
		Promp: nil,
	}

	Mapping["export"] = &Command{
		Call:  runExport,
		Usage: nil,
		Desc:  "Export a script.",
		Promp: nil,
	}

	Mapping["globals"] = &Command{
		Call:  runManageGlobals,
		Usage: usageGlobal,
		Desc:  "Manage global variables.",
		Promp: nil,
	}

	Mapping["run"] = &Command{
		Call:  runRunScript,
		Usage: nil,
		Desc:  "Execute the script.",
		Promp: nil,
	}

	Mapping["exit"] = &Command{
		Call:  runExit,
		Usage: nil,
		Desc:  "Properly exit the too.",
		Promp: nil,
	}

	Mapping["reload"] = &Command{
		Call:  runReload,
		Usage: usageReload,
		Desc:  "Reload (root|script).",
		Promp: nil,
	}

	Mapping["search"] = &Command{
		Call:  runSearch,
		Usage: usageSearch,
		Desc:  "Search a script using patterns.",
		Promp: nil,
	}

	Mapping["info"] = &Command{
		Call:  runInfo,
		Usage: nil,
		Desc:  "Information reguarding the loaded script.",
		Promp: nil,
	}

	Mapping["options"] = &Command{
		Call:  runOptions,
		Usage: nil,
		Desc:  "Show configurable variables for loaded script.",
		Promp: nil,
	}

	Mapping["lua"] = &Command{
		Call:  runLua,
		Usage: nil,
		Desc:  "Execute inline lua commands.",
		Promp: nil,
	}

	Mapping["back"] = &Command{
		Call:  runBack,
		Usage: nil,
		Desc:  "Free the script.",
		Promp: nil,
	}

	Mapping["use"] = &Command{
		Call:  runUse,
		Usage: usageUse,
		Desc:  "Use a script.",
		Promp: nil,
	}

	Mapping["banner"] = &Command{
		Call:  runBanner,
		Usage: nil,
		Desc:  "Show banner.",
		Promp: nil,
	}

	Mapping["vpm"] = &Command{
		Call:  runVPM,
		Usage: nil,
		Desc:  "Use Venera package manager.",
		Promp: nil,
	}

	Mapping["set"] = &Command{
		Call:  runSet,
		Usage: nil,
		Desc:  "Set a ariable",
		Promp: nil,
	}
}
