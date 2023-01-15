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
		SCListScripts(cmds)


	} else if cmds[0] == "info" {
		if p.SSet {
			wlua.MetaShow()
		} else {
			PrintErr("No module setted. Type `help`.")
		}


	} else if cmds[0] == "back" || cmds[0] == "b" {
		if p.SSet {
			p.State.Close()
			p.Prompt = "[*]>> "
			LivePrefixState.LivePrefix = p.Prompt
			LivePrefixState.IsEnable = true
			p.Script = ""
			p.SSet = false
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