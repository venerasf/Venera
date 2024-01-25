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

// Will keep it global for now.
var Mapping = make(map[string]Command)

/*
	TODO: Refactor the function completely.
	Maybe use a functional approach for function mapping.
		Change from `cmds` to `args`.
*/
func (profile *Profile) Execute(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmds := strings.Split(cmd, " ")
	length := len(cmds)

  
	// Validates length
	if (length == 0) {
		utils.PrintErr("Too few arguments. Try `help cmd`.")
		return
	}

	/*switch cmds[0] {
		case "help":
			if length >= 2 {
				callUsage(Mapping[cmds[1]].Usage,profile)
			} else {
				CmdHelp()
			}
		default:
			err := Mapping[cmds[0]].Call(cmds)
			if err != nil {
				utils.PrintErr(err.Error())
			}
	}*/

	// header
	h := cmds[0]

	// Generic commands
    if h == "help" {
		runHelp(cmds, profile)

	} else if h == "bash"{
		// Starts a bash shell
		utils.GetBash()
	} else if h == "import" {
		runImport(cmds, profile)
	} else if h == "export" {
		runExport(cmds, profile)
	} else if h == "globals" {
		runManageGlobals(cmds, profile)

	} else if h == "exit"  ||  h == "e"  ||  h == "quit" {
		runExit(cmds, profile)

	} else if h == "reload" {
		runReload(cmds, profile)

	} else if h == "set" {
		runSet(cmds, profile)
	
	} else if h == "r" ||  h == "run"  ||  h == "exploit" {
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
	}
}

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

func runHelp(cmds []string,profile *Profile) int {
	CmdHelp()
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

func runExport(cmds []string,profile *Profile) {
	// Exports a script
	if len(cmds) != 3 {
		utils.PrintErr("Invalid arguments.")
	} else {
		profile.SCExportScript(cmds[1], cmds[2])
	}
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

func runReload(cmds []string,profile *Profile) {
	if len(cmds) != 2 {
		utils.PrintErr("Invalid args.")
		return
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
}


func runManageGlobals(cmds []string, profile *Profile) {
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
}

// the funcs um this line isnt in defcallbacks
// -------------------------------------------

// call usage of a script
func callUsage(usage func(), p *Profile, cmds []string) {
	usage()
}

func loadFunctions() {
	
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
