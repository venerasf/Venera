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
	print("\n")
	t.Print()
	print("\n")
	println("    Press `TAB` to rotate suggestions.")
	println("    Press `arrow key` to pass suggentions or history.")
	println("    Press `CTRL-D` to exit.")
	print("\n")
}


func Banner() {
	stb := ""
	if Stable {
		stb = "Stable"
	} else {
		stb = "NotStable"
	}
	fmt.Printf("### Venera %.2f-%s ###\n",Version,stb)
	/*print(`
	__    _ ____________________________________
	\ \  | |   ___   _  _    ___   _  _   __ _  |
	 \ \ | |  / _ \ | \| |  / _ \ | |//  / _' | |
	  \ \| | |  __/ |  \ | |  __/ | |/  | (_| | |
	   \___|  \___| |_| _|  \___| |_|    \__,_| |
	   -----------------------------------------+
	   Recon Mission: github.com/farinap5/venera
	`)*/
}