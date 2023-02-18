package src

import (
	"fmt"

	"github.com/cheynewallace/tabby"
)

func CmdHelp() {
	t := tabby.New()
	t.AddHeader("COMMAND","DESCRIPTION")
	t.AddLine("help","Show help menu")
	t.AddLine("bash","Spawn shell")

	t.AddLine("use","Load a script/module")
	t.AddLine("back","Exit module/script")
	t.AddLine("options","Show variables of script/module")
	t.AddLine("info","Info/metadata about script/module")
	t.AddLine("globals","Show global variables")
	t.AddLine("run","Run a script/module")
	t.AddLine("set","Set value for a variable")
	t.AddLine("lua","Run Lua code in running script")
	t.AddLine("import","Import a (to edited) script")
	t.AddLine("export","Export a script (to edit)")
	t.AddLine("exir","Exit from prompt")
	print("\n")
	t.Print()
	print("\n")
	println("BASIC NAVEGATION:")
	println("    Press `TAB` to rotate suggestions.")
	println("    Press `arrow key` to pass suggentions or history.")
	println("    Press `CTRL-d` to exit.")
	println("    Press `CTRL-l` to clear prompt.")

	print("\n")
	println("SEARCHING:")
	println("    `search` list scripts/modules.")
	println("    `search match <key>` list matching patterns.")
	println("    `search match:path <key>` list path matching.")
	println("    `search match:description <key>` list description matching.")
	println("    `search tag <tag1 tag2...>` list tags matching.")
	println("    `s m <key>` filter in collapsed format.")
	println("        `s m:p <key>` filter by path.")
	println("        `s m:d <key>` filter by description.")
	println("        `s t <tag1 tag2...>` filter by tags.")

	print("\n")
	println("USE SCRIPT/MODULE:")
	println("    `use path/to/script.lua` Configure a script.")
	println("    `use tags http sql` Set scripts matching with tags.")

	print("\n")
	println("SET VARIABLE:")
	println("    `set RHOST <value>` Configure variable in a script.")
	println("    `set global RHOST <value>` Configure variable to a chain of scripts.")
	print("\n")
	print("\n")
}


func Banner() {
	stb := ""
	if Stable {
		stb = "Stable"
	} else {
		stb = "NotStable"
	}
	//fmt.Printf("### Venera %.2f-%s ###\nType `help`\n\n",Version,stb)
	fmt.Printf(`
	__    _ ____________________________________
	\ \  | |   ___   _  _    ___   _  _   __ _  |
	 \ \ | |  / _ \ | \| |  / _ \ | |//  / _' | |
	  \ \| | |  __/ |  \ | |  __/ | |/  | (_| | |
	   \___|  \___| |_| _|  \___| |_|    \__,_| |
	   -----------------------------------------+
	   Recon Mission: github.com/farinap5/venera
                type 'help'      %.2f-%s

`,Version,stb)
}
