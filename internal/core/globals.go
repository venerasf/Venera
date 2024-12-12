package core

import (
	"strings"
	"venera/internal/types"

	"github.com/cheynewallace/tabby"
)

// ################################ Global variables ################################
// / Set globals
func SetGlobals(p *types.Profile, key string, value string) {
	p.Globals[key] = value
	p.Database.DBStoreGlobal(key, value)
}

func ListGlobals(p types.Profile) {
	t := tabby.New()
	t.AddHeader("VARIABLE", "VALUE")
	for key, value := range p.Globals {
		t.AddLine(key, value)
	}
	print("\n")
	t.Print()
	print("\n")
}

func runManageGlobals(cmds []string, profile *types.Profile) int {
	length := len(cmds)
	if length == 3 && cmds[1] == "rm" {
		profile.Database.DBRemoveGlobals(cmds[2])
		profile.Globals = nil
		profile.Globals = profile.Database.DBLoadIntoGlobals()
		// may be changed to >= 4
	} else if length == 4 && cmds[1] == "set" {
		SetGlobals(profile, cmds[2], strings.Join(cmds[3:], " "))
	} else {
		// Lists global variables
		ListGlobals(*profile)
	}
	return 0
}