package src

import (
	"os"
	"strings"
	"venera/src/wlua"
	"github.com/cheynewallace/tabby"
	//"github.com/c-bata/go-prompt"
)

func (p *Profile) Execute(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmds := strings.Split(cmd, " ")


	if cmd == "help" {
		CmdHelp()
	} else if cmds[0] == "options" {
		if p.SSet || p.Chain {
			wlua.VarsList()
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmds[0] == "search" || cmds[0] == "s" {
		p.SCListScripts(cmds)


	} else if cmds[0] == "info" {
		if p.SSet {
			wlua.MetaShow()
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmds[0] == "elf" {
		PrintSuccs("Elf")


	} else if cmds[0] == "reload" {
		ReloadScript(p)


	} else if cmds[0] == "global" || cmds[0] == "globals" || cmds[0] == "g" {
		p.ListGlobals()

	
	} else if cmds[0] == "import" && len(cmds) == 3 {
		p.SCImportScript(cmds[1], cmds[2])
	} else if cmds[0] == "export" && len(cmds) == 3 {
		p.SCExportScript(cmds[1], cmds[2])


	} else if cmds[0] == "back" || cmds[0] == "b" {
		if p.SSet || p.Chain {
			FreeScript(p)
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmds[0] == "lua" && len(cmds) >= 2 {
		if p.SSet {
			wlua.LuaExecString(p.State, strings.Join(cmds[1:], " "))
		} else {
			PrintErr("No module setted. Type `help`.")
		}

	
	// Set global variable
	/*}  else if cmds[0] == "set" && len(cmds) >= 4 {
		if cmds[1] == "global" || cmds[1] == "g" || cmds[1] == "globals" {
			p.SetGlobals(cmds[2], strings.Join(cmds[3:]," "))
		}*/


	// Set normal variable
	} else if cmds[0] == "set" && len(cmds) >= 3 {
		if (cmds[1] == "global" || cmds[1] == "g" || cmds[1] == "globals") && len(cmds) >= 4 {
			p.SetGlobals(cmds[2], strings.Join(cmds[3:], " "))
		} else {
			if p.SSet {
				//print("aaaaa")
				wlua.SetVarValue(p.State, cmds[1], strings.Join(cmds[2:], " "))
			} else {
				if p.Chain {
					PrintErr("Use global variable.")
				} else {
					PrintErr("No module setted. Type `help`.")
				}
			}
		}


	// spawn bash, useless func
	} else if cmd == "bash" {
		GetBash()


	// just for tests
	} else if cmds[0] == "setp" && len(cmds) == 2 {
		p.Prompt = "[" + cmds[1] + "]>> "
		LivePrefixState.LivePrefix = p.Prompt
		LivePrefixState.IsEnable = true
		return

	
	/* TODO:
	join "use" and "use tag" blocks in one function.
	*/
	// Use a script
	} else if cmds[0] == "use" && len(cmds) == 2 {
		if !p.SSet {
			useScript(p, cmds)
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	// Use a set of scripts based on tags
	} else if cmds[0] == "use" && len(cmds) >= 3 && (cmds[1] == "tags" || cmds[1] == "tag" || cmds[1] == "t") {
		if !p.SSet {
			useScriptTAG(p, cmds)
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmd == "run" || cmd == "exploit" || cmd == "r" {
		if !p.Chain {
			runScript(p)
			//print("aaa")
		} else {
			runChain(p)
		}

	} else if cmds[0] == "exit" || cmds[0] == "e" || cmds[0] == "quit" {
		HandleExit()
		os.Exit(0)
	} else if cmds[0] == "banner" {
		Banner()
	} else {
		PrintErr("Not a command. Type `help`.")
	}
}

// Load script
func useScript(p *Profile, cmds []string) {
	p.Script = cmds[1]                     // Set script as passed over cmd
	profile := *p                          // Take off pointer
	pl := wlua.LuaProfile(profile)         // Convert Profile to LuaProfile
	p.State, p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		PrintErr("Error loading script/module.")
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
func ReloadScript(p *Profile) {
	aux := p.Script

	PrintSuccs("Freeing memory.")
	// Free script
	p.State.Close()
	p.SSet = false
	p.Script = ""
	wlua.LuaFreeScript()
	p.Prompt = "[*]>> "
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true

	// load script
	PrintSuccs("Loading " + aux)
	p.Script = aux                         // Set script as passed over cmd
	profile := *p                          // Take off pointer
	pl := wlua.LuaProfile(profile)         // Convert Profile to LuaProfile
	p.State, p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		PrintErr("Error loading script/module.")
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
		PrintErr("Error loading tags, no script found.")
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
		PrintErr("Error loading tags, no script found.")
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
