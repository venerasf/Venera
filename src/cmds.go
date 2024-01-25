package src

import (
	"strings"
	"venera/src/utils"
)

func init() {
	loadFunctions()
}

// Will keep it global for now.
var Mapping = make(map[string]*Command)

func (profile *Profile) Execute(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmds := strings.Split(cmd, " ")
	length := len(cmds)

	// Validates length
	if (length == 0) {
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
					functionP.Usage()
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

	/*// header
	h := cmds[0]

	// Generic commands
    if h == "help" {
		runHelp(cmds, profile)
	} else if h == "bash"{
		runBash(cmds, profile)
	} else if h == "import" {
		runImport(cmds, profile)
	} else if h == "export" {
		runExport(cmds, profile)
	} else if h == "globals" {
		runManageGlobals(cmds, profile)

	} else if h == "exit" {
		runExit(cmds, profile)

	} else if h == "reload" {
		runReload(cmds, profile)

	} else if h == "set" {
		runSet(cmds, profile)
	
	} else if h == "run" {
		runRunScript(cmds, profile)

	} else if h == "search" {
		runSearch(cmds, profile)

	} else if h == "info" {
		runInfo(cmds, profile)

	} else if h == "lua" {
		runLua(cmds, profile)

	} else if h == "vpm" {
		runVPM(cmds, profile)

	} else if h == "back"  ||  h == "b" {
		// Removes the current selected script
		runBack(cmds, profile)

	} else if h == "options"  ||  h == "o" {
		// Removes the current selected script
		runOptions(cmds, profile)
	
	} else if h == "use" {
		// Uses a script
		runUse(cmds, profile)

	} else if h == "banner" {
		// Uses a script
		runBanner(cmds, profile)
			
	} else {
		utils.PrintErr("Not a valid command or missing a selected script. Type `help`.")
	}*/
}



/* 
	Register the default commands
	type Command struct {
		Call 	func([]string, *Profile) int // Callback entrypoint
		Usage 	func() // help function callback
		Desc 	string // hight level description.
		Promp 	[][]string // Prompt help and auto-complete for subcmds
	}
*/
func loadFunctions() {
	 Mapping["help"] = &Command{
		Call: runHelp,
		Usage: nil,
		Desc: "Show help menu.",
		Promp: nil,
	}

	Mapping["bash"] = &Command{
		Call: runBash,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["import"] = &Command{
		Call: runImport,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["export"] = &Command{
		Call: runExport,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["globals"] = &Command{
		Call: runManageGlobals,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["run"] = &Command{
		Call: runRunScript,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["globals"] = &Command{
		Call: runManageGlobals,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["exit"] = &Command{
		Call: runExit,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["reload"] = &Command{
		Call: runReload,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["search"] = &Command{
		Call: runSearch,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["info"] = &Command{
		Call: runInfo,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["options"] = &Command{
		Call: runOptions,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["lua"] = &Command{
		Call: runLua,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}


	Mapping["back"] = &Command{
		Call: runBack,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["use"] = &Command{
		Call: runUse,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["banner"] = &Command{
		Call: runBanner,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}

	Mapping["vpm"] = &Command{
		Call: runVPM,
		Usage: nil,
		Desc: "Execute Bash.",
		Promp: nil,
	}
}