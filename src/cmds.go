package src

import (
	"os"
	"strings"
	"venera/src/wlua"
	"venera/src/utils"
	"github.com/cheynewallace/tabby"
	//"github.com/c-bata/go-prompt"
)

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
		return; 
	}

	// header
	h := cmds[0] 

	// Generic commands
    if h == "help" {
		// Displays help pannel
		CmdHelp()



	} else if h == "bash"{
		// Starts a bash shell
		utils.GetBash()



	} else if h == "import" {
		// Imports a script
		if length != 3 {
			utils.PrintErr("Invalid arguments.")
		} else {
			profile.SCImportScript(cmds[1], cmds[2])
		}



	} else if h == "export" {
		// Exports a script
		if length != 3 {
			utils.PrintErr("Invalid arguments.")
		} else {
			profile.SCExportScript(cmds[1], cmds[2])
		}



	} else if h == "globals" {
		// Lists global variables
		profile.ListGlobals()



	} else if h == "exit"  ||  h == "e"  ||  h == "quit" {
		// Exits the program
		HandleExit()
		os.Exit(0)



	} else {
		// Not generic commands
		if profile.SSet || profile.Chain {
			// If a script is set
			if h == "set" {
				// Sets a variable for the script
				if length < 3 {
					utils.PrintErr("Invalid arguments.")
				} else {
					if (cmds[1] == "global" || cmds[1] == "g" || cmds[1] == "globals") {
						//continuar
						utils.PrintSuccs("Setting global variable")
						profile.SetGlobals(cmds[2], strings.Join(cmds[3:], " "))
					} else {
						wlua.SetVarValue(profile.State, cmds[1], strings.Join(cmds[2:], " "))
					}
				}



			} else if h == "run"  ||  h == "r"  ||  h == "exploit" {
				// Runs the script
				if profile.Chain {
					runChain(profile)
				} else {
					runScript(profile)
				}



			} else if h == "back"  ||  h == "b" {
				// Removes the current selected script
				FreeScript(profile)



			} else if h == "options" {
				// Displays lua's variables
				wlua.VarsList()



			} else if h == "lua" {
				// Executes lua code
				if length < 2 {
					utils.PrintErr("Invalid arguments.")
				} else {
					wlua.LuaExecString(profile.State, strings.Join(cmds[1:], " "))
				}



			} else if h == "info" {
				// Displays information
				wlua.MetaShow()



			} else if h == "reload" {
				// Reloads the selected script
				profile.ReloadScript()



			} else {
				// No commands recognized
				utils.PrintErr("Not a valid command or unable to call it because a script is selected.")



			}

		} else {
			// If there's no script set
			if h == "search"  ||  h == "s" {
				// Searches a script
				profile.SCListScripts(cmds)



			} else if h == "use" {
				// Uses a script
				if length < 2 {
					utils.PrintErr("Invalid arguments.")
				} else {
					if (cmds[1] == "tags" || cmds[1] == "tag" || cmds[1] == "t") {
						utils.PrintSuccs("Using tag context")
						if length < 3 {
							utils.PrintErr("Invalid Arguments.")
						} else {
							useScriptTAG(profile, cmds)
						}
					} else {
						useScript(profile, cmds)
					}
				}



			} else {
				// No commands found
				utils.PrintErr("Not a valid command or missing a selected script. Type `help`.")



			}
		}
	} // end

	found := false
	// misc commands
	if cmds[0] == "elf" {
		found = true
		utils.PrintSuccs("Elf")

	} else if cmds[0] == "setp" && len(cmds) == 2 {
		found = true
		profile.Prompt = "[" + cmds[1] + "]>> "
		LivePrefixState.LivePrefix = profile.Prompt
		LivePrefixState.IsEnable = true

	} else if cmds[0] == "banner" {
		found = true
		Banner()
	}

	if found { utils.PrintSuccs("Yet it still did something.") }
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

	p.Prompt = "(" + cmds[1] + ")>> " // Change prompt
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
