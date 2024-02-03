package src

import (
	"os"
	"strings"
	"venera/src/pacman"
	"venera/src/utils"
	"venera/src/wlua"

	"github.com/cheynewallace/tabby"
	//"github.com/c-bata/go-prompt"
)

func runBanner(cmds []string,profile *Profile) int {
	Banner()
	return 0
}

func runUse(cmds []string,profile *Profile) int {
	if !profile.SSet && !profile.Chain {
		if len(cmds) < 2 {
			utils.PrintErr("Invalid arguments.")
		} else {
			if (cmds[1] == "tags" || cmds[1] == "tag" || cmds[1] == "t") {
				utils.PrintSuccs("Using tag context")
				if len(cmds) < 3 {
					utils.PrintErr("Invalid Arguments.")
				} else {
					useScriptTAG(profile, cmds)
				}
			} else {
				useScript(profile, cmds)
			}
		}
	} else {
		utils.PrintErr("Free the script with `back` before using another .")
	}
	return 0
}

// Exit from a script
func runOptions(cmds []string,profile *Profile) int {
	if profile.SSet || profile.Chain {
		wlua.VarsList()
	} else {
		utils.PrintErr("Must have script setted.")
	}
	return 0
}

// Exit from a script
func runBack(cmds []string,profile *Profile) int {
	if profile.SSet || profile.Chain {
		FreeScript(profile)
	} else {
		utils.PrintErr("Must have script setted.")
	}
	return 0
}

func runInfo(cmds []string,profile *Profile) int {
	if profile.SSet || profile.Chain {
		// Displays information
		if profile.Chain {
			profile.SCInfoForChaining()
		} else {
			wlua.MetaShow()	
		}
	} else {
		utils.PrintErr("Must have script setted.")
	}
	return 0
}

func runSearch(cmds []string,profile *Profile) int {
	// Searches a script
	profile.SCListScripts(cmds)
	return 0
}

func runRunScript(cmds []string,profile *Profile) int {
	if profile.SSet || profile.Chain {
		// Runs the script
		if profile.Chain {
			runChain(profile)
		} else {
			runScript(profile)
		}
	} else {
		utils.PrintErr("Must have script setted.")
	}
	return 0
}

func runSet(cmds []string,profile *Profile) int {
	if profile.SSet || profile.Chain {
		// Sets a variable for the script
		if len(cmds) < 3 {
			utils.PrintErr("Invalid arguments.")
		} else {
			wlua.SetVarValue(profile.State, cmds[1], strings.Join(cmds[2:], " "))
		}
	} else {
		utils.PrintErr("Must have script setted.")
	}
	return 0
}

func runHelp(cmds []string, p *Profile) int {
	if len(cmds) >= 2 {
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
	return 0
}

func runVPM(cmds []string,profile *Profile) int {
	// TODO: Must not be executed with scrpt setted
	// TODO: Validate return code
	return pacman.VPMGetRemotePack(
		profile.Globals["repo"],
		profile.Globals["root"],
		profile.Globals["sign"],
		cmds, 
		*profile.Database,
		profile.Globals["vpmvs"],
	)
}

func runExport(cmds []string,profile *Profile) int {
	// Exports a script
	if len(cmds) != 3 {
		utils.PrintErr("Invalid arguments.")
	} else {
		profile.SCExportScript(cmds[1], cmds[2])
	}
	return 0
}

func runImport(cmds []string,profile *Profile) int {
	// Imports a script
	if len(cmds) != 3 {
		utils.PrintErr("Invalid arguments.")
	} else {
		profile.SCImportScript(cmds[1], cmds[2])
	}
	return 0
}

func runBash(cmds []string,profile *Profile) int {
	// Starts a bash shell
	utils.GetBash()
	return 0
}

func runExit(cmds []string,profile *Profile) int {
	// Exits the program
	HandleExit()
	os.Exit(0)
	return 0 // wont run
}

func runLua(cmds []string,profile *Profile) int {
	if profile.SSet || profile.Chain {
		length := len(cmds)
		// Executes lua code
		if length < 2 {
			utils.PrintErr("Invalid arguments.")
		} else {
			wlua.LuaExecString(profile.State, strings.Join(cmds[1:], " "))
		}
	} else {
		utils.PrintErr("Must have script setted.")
	}
	return 0
}

func runReload(cmds []string,profile *Profile) int {
	if len(cmds) != 2 {
		utils.PrintErr("Invalid args.")
		return 1
	}

	if cmds[1] == "script" || cmds[1] == "s" {
		// Reloads the selected script
		if profile.SSet { 
			profile.ReloadScript()
		} else {
			utils.PrintErr("Not a valid command out from script.")
		}
	} else if cmds[1] == "root" {
		profile.SCLoadScripts()
	}
	return 0
}


func runManageGlobals(cmds []string, profile *Profile) int {
	length := len(cmds)
	if length == 3 && cmds[1] == "rm" {
		profile.Database.DBRemoveGlobals(cmds[2])
		profile.Globals = nil
		profile.Globals = profile.Database.DBLoadIntoGlobals()
		// may be changed to >= 4
	} else if length == 4 && cmds[1] == "set" {
		profile.SetGlobals(cmds[2], strings.Join(cmds[3:], " "))
	} else {
		// Lists global variables
		profile.ListGlobals()
	}
	return 0
}


// Load script
func useScript(p *Profile, cmds []string) {
	p.Script = cmds[1]                     // Set script as passed over cmd
	profile := *p                          // Take off pointer
	pl := wlua.LuaProfile(profile)         // Convert Profile to LuaProfile
	p.State, p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		utils.PrintErr("Error loading script/module.")
		return
	}

	// hide the root path and extension when prompting the script path
	promptedPath := utils.HideBasePath(p.Globals["root"], cmds[1])
	promptedPath = utils.HideLuaExtension(promptedPath)

	p.Prompt = "(" + promptedPath + ")>> " // save new prompt

	// set new prompt
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true
}



func runScript(p *Profile) {
	if p.SSet {
		wlua.LuaRunUniq(p.State)
	} else {
		println("No Script")
	}
}

// Erase everything of a script from the memory
func FreeScript(p *Profile) {
	if p.Chain {
		//p.State.Close()
		p.Chain = false
		//print("cleaning\n")
	}

	p.SSet = false
	p.Chain = false
	p.Script = ""
	wlua.LuaFreeScript()
	p.Prompt = "[*]>> "
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true
}

// This function will reload a script
func (p *Profile)ReloadScript() {
	aux := p.Script

	utils.PrintSuccs("Freeing memory.")
	// Free script
	p.State.Close()
	p.SSet = false
	p.Script = ""
	wlua.LuaFreeScript()
	p.Prompt = "[*]>> "
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true

	// load script
	utils.PrintSuccs("Loading " + aux)
	p.Script = aux                         // Set script as passed over cmd
	profile := *p                          // Take off pointer
	pl := wlua.LuaProfile(profile)         // Convert Profile to LuaProfile
	p.State, p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		utils.PrintErr("Error loading script/module.")
		return
	}

	p.Prompt = "(" + aux + ")>> " // Change prompt
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true
}

// ################################ Global variables ################################
// / Set globals
func (p *Profile) SetGlobals(key string, value string) {
	p.Globals[key] = value
	p.Database.DBStoreGlobal(key, value)
}

func (p Profile) ListGlobals() {
	t := tabby.New()
	t.AddHeader("VARIABLE", "VALUE")
	for key, value := range p.Globals {
		t.AddLine(key, value)
	}
	print("\n")
	t.Print()
	print("\n")
}

func useScriptTAG(p *Profile, cmds []string) {
	var scriptslist []string

	var scriptScanner []ScriptTAGInfo
	aux := SCTAG

	/*
		When using scripts based on tags the script cant be hard to configure
		or has a complex stucture like asking prompts from the user, se scripts to
		be used with tags are specified with the tag "scanner".

		First its important to get all those scripts that have "scanner" tag,
		and after that we match the tags
	*/
	for _, sti := range aux {
		for i := range sti.Tag {
			if sti.Tag[i] == "scanner" {
				scriptScanner = append(scriptScanner, sti)
			}
		}
	}
	if len(scriptScanner) == 0 {
		utils.PrintErr("Error loading tags, no script found.")
		return
	}

	for _, sti := range scriptScanner {
		for i := range sti.Tag {
			for _, j := range cmds[2:] {
				if sti.Tag[i] == j {
					scriptslist = append(scriptslist, sti.Path)
					break
				}
			}
		}
	}
	if len(scriptslist) == 0 {
		utils.PrintErr("Error loading tags, no script found.")
		return
	}

	p.Scriptslist = scriptslist
	profile := *p // Take off pointer
	pl := wlua.LuaProfile(profile)
	wlua.GetVarsToChainTAGS(pl)
	//wlua.PopulateLoadVarsFromGlobals(pl)

	p.Prompt = "(" + JoinTgs(cmds[2:]) + ")>> " // Change prompt
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true
	p.Chain = true
}

func runChain(p *Profile) {
	profile := *p // Take off pointer
	pl := wlua.LuaProfile(profile)
	wlua.LuaRunChaining(pl)
	//p.Chain = true

	/*for _,i := range (p.Scriptslist) {
		println(i)
	} */
}
