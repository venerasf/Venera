package src

import (
	"strings"
	"venera/src/wlua"
	//"github.com/c-bata/go-prompt"
)

func (p *Profile) Execute(cmd string) {
	cmd = strings.TrimSpace(cmd)
	cmds := strings.Split(cmd," ")
	
	if cmd == "help" {
		CmdHelp()
	} else if cmds[0] == "options" {
		if p.SSet {
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

	} else if cmds[0] == "reload" {
		ReloadScript(p)


	} else if cmds[0] == "import" && len(cmds) == 3{
		p.SCImportScript(cmds[1],cmds[2])
	} else if cmds[0] == "export" && len(cmds) == 3{
		p.SCExportScript(cmds[1],cmds[2])


	} else if cmds[0] == "back" || cmds[0] == "b" {
		if p.SSet {
			FreeScript(p)
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmds[0] == "lua" && len(cmds) >= 2 {
		if p.SSet {
			wlua.LuaExecString(p.State,strings.Join(cmds[1:]," "))
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	}  else if cmds[0] == "set" && len(cmds) == 3 {
		if p.SSet {
			wlua.SetVarValue(p.State,cmds[1],cmds[2])
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmd == "bash" {
		GetBash()


	} else if cmds[0] == "setp" && len(cmds) == 2 {
		p.Prompt = "["+cmds[1]+"]>> "
		LivePrefixState.LivePrefix = p.Prompt
		LivePrefixState.IsEnable = true
		return


	} else if cmds[0] == "use" && len(cmds) == 2 {
		if !p.SSet {
			useScript(p,cmds)
		} else {
			PrintErr("No module setted. Type `help`.")
		}
		

	} else if cmd == "run" || cmd == "exploit" {
		runScript(p)


	} else {
		PrintErr("Not a command. Type `help`.")
	}
}

// Load script
func useScript(p *Profile, cmds []string) {
	p.Script = cmds[1] // Set script as passed over cmd
	profile := *p // Take off pointer
	pl := wlua.LuaProfile(profile) // Convert Profile to LuaProfile
	p.State,p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		PrintErr("Error loading script/module.")
		return
	}

	p.Prompt = "("+cmds[1]+")>> " // Change prompt
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
	p.State.Close()
	p.SSet = false
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
	PrintSuccs("Loading "+aux)
	p.Script = aux // Set script as passed over cmd
	profile := *p // Take off pointer
	pl := wlua.LuaProfile(profile) // Convert Profile to LuaProfile
	p.State,p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		PrintErr("Error loading script/module.")
		return
	}

	p.Prompt = "("+aux+")>> " // Change prompt
	LivePrefixState.LivePrefix = p.Prompt
	LivePrefixState.IsEnable = true
}