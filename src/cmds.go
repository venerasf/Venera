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
		println("Help")
	} else if cmds[0] == "options" {
		if p.SSet {
			wlua.VarsList()
		}
	} else if cmds[0] == "info" {
		if p.SSet {
			wlua.MetaShow()
		}
	}  else if cmds[0] == "set" && len(cmds) == 3 {
		if p.SSet {
			wlua.SetVarValue(p.State,cmds[1],cmds[2])
		}
	} else if cmd == "bash" {
		GetBash()
	} else if cmds[0] == "setp" && len(cmds) == 2 {
		p.Prompt = "["+cmds[1]+"]>> "
		LivePrefixState.LivePrefix = p.Prompt
		LivePrefixState.IsEnable = true
		return
	} else if cmds[0] == "use" && len(cmds) == 2 {
		useScript(p,cmds)
	} else if cmd == "run" || cmd == "exploit" {
		runScript(p)
	} else {
		println("aaa")
	}
}

// Load script
func useScript(p *Profile, cmds []string) {
	p.Script = cmds[1] // Set script as passed over cmd
	profile := *p // Take off pointer
	pl := wlua.LuaProfile(profile) // Convert profile to LuaProfile
	p.State,p.SSet = wlua.LuaInitUniq(pl) // Init script
	if !p.SSet {
		println("Error")
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