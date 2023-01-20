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
	t.AddLine("run","Run a script/module")
	t.AddLine("set","Set value for a ver")
	t.AddLine("lua","Run Lua code in running mod")
	t.AddLine("import","Import a (to edited) script")
	t.AddLine("export","Export a script (to edit)")
	print("\n")
	t.Print()
	print("\n")
	println("BASIC NAVEGATION:")
	println("    Press `TAB` to rotate suggestions.")
	println("    Press `arrow key` to pass suggentions or history.")
	println("    Press `CTRL-d` to exit.")
	println("    Press `CTRL-l` to clear promp.")

	print("\n")
	println("SEARCHING:")
	println("    `serach` list scripts/modules.")
	println("    `serach match <key>` list witch maches patter.")
	println("    `serach match:path <key>` list path matching.")
	println("    `serach match:description <key>` list description matching.")
	println("    `s m <key>` filter in collapsed format.")
	println("        `s m:p <key>` filter by path.")
	println("        `s m:d <key>` filter by description.")

	print("\n")
	println("SET VARIABLE:")
	println("    `set RHOST <value>` Configure var for an in use test.")
	println("    `set global RHOST <value>` Configure var to a chain of tests.")
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