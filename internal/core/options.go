package core

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cheynewallace/tabby"
)

func CmdHelp() {
	print("\n")
	t := tabby.New()
	t.AddHeader("GENERIC COMMAND", "DESCRIPTION")
	t.AddLine("help", "Show help menu. Use `help <cmd>`.") //
	t.AddLine("bash", "Spawns a shell")                    //
	t.AddLine("import", "Import a (edited) script")        //
	t.AddLine("export", "Export a script (to edit)")       //
	t.AddLine("globals", "Show global variables")          //
	t.AddLine("vpm", "Venera package manager")             //
	t.AddLine("exit", "Exits the prompt")                  //

	t.AddLine("search", "Searches a script/module") //
	t.AddLine("use", "Load a script/module\n")      //

	t.AddHeader("SCRIPT COMMAND", "DESCRIPTION")               //
	t.AddLine("set", "Set value for a variable")               //
	t.AddLine("run", "Run a script/module")                    //
	t.AddLine("back", "Exit module/script")                    //
	t.AddLine("options", "Show variables of script/module")    //
	t.AddLine("lua", "Run Lua code in running script")         //
	t.AddLine("info", "Info/metadata about script/module")     //
	t.AddLine("reload", "Reloads the current script/module")   //
	t.Print()
	print("\n")
	/*println("BASIC NAVEGATION:")
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

	print("\n")
	println("USE SCRIPT/MODULE:")
	println("    `use path/to/script.lua` Configure a script.")
	println("    `use tags http sql` Set scripts matching with tags.")

	print("\n")
	println("SET VARIABLE:")
	println("    `set RHOST <value>` Configure variable in a script.")
	println("    `global set RHOST <value>` Configure variable to a chain of scripts.")
	print("\n")
	print("\n")*/
}

func usageSet(cmds []string) {
	print("\n")
	println("SET:")
	println("    `set` Assign value to a variable within the context of a script.")
	println("    `set <key> <value>` Set a variable.")
	print("\n")
}

func usageGlobal(cmds []string) {
	print("\n")
	println("GLOBALS:")
	println("    `globals` Show global variables.")
	println("    `globals set <key> <value>` Set a global variable.")
	println("    `globals rm <key>` Remove a global variable.")
	println("    When a script is loaded, like typing `use`, if the script has a variable with the same name of a global variable, it receives automatically the value from globals.")
	print("\n")
}

func usageUse(cmds []string) {
	print("\n")
	println("USE:")
	println("    `use <script>` Load a script.")
	println("    When executed, the `Init()` function is called for default routines.")
}

func usageHelp(cmds []string) {
	println("HELP:")
	println("    `help` Show help menu.")
	println("    `help <cmd>` Show help menu for that command.")
	println("    `help <cmd> arg1 arg2` Arguments are accepted if implemented for that command.")
	print("\n")
}

func usageReload(cmds []string) {
	print("\n")
	println("RELOAD:")
	println("    `reload script` Free memory and load script again.")
	println("    `reload root` Will reload the root directory.")
	println("    May be used when you set a new root directory from where Venera loads the scripts.")
	print("\n")
}

func usageSearch(cmds []string) {
	print("\n")
	println("SEARCHING:")
	println("    `search` list scripts.")
	println("    `search match <key>` list matching patterns.")
	println("    `search match:path <key>` list path matching.")
	println("    `search match:description <key>` list description matching.")
	println("    `search tag <tag1 tag2...>` list matching tags.")
	print("\n")
}

func usageVPM(cmds []string) {
	print("\n")
	println("Venera Package Manager:")
	println("    `vpm` Call vpm commands.")
	println("    `vpm search <pattern>` List matching substring.")
	println("    `vpm install <script>` Install a script.")
	println("       Usage:`vpm install /path/to/the/script.lua`")
	println("    `vpm sync` Sincronize with remote repository.")
	println("    `vpm verify` Verify the signature of the configured remote repository.")
	print("\n")
}

func Banner() {
	stb := ""
	if Stable {
		stb = "Stable"
	} else {
		stb = "NotStable"
	}
	// fmt.Printf("### Venera %.2f-%s ###\nType `help`\n\n",Version,stb)

	r := rand.NewSource(time.Now().UnixNano())
	rn := rand.New(r)
	x := rn.Intn(6)
	if x == 0 {
		fmt.Printf(`
	__    _ ____________________________________
	\ \  | |   ___   _  _    ___   _  _   __ _  |
	 \ \ | |  / _ \ | \| |  / _ \ | |//  / _' | |
	  \ \| | |  __/ |  \ | |  __/ | |/  | (_| | |
	   \___|  \___| |_| _|  \___| |_|    \__,_| |
	   -----------------------------------------+
	   Recon Mission: github.com/venerasf/Venera
	   Read the docs https://venera.farinap5.com/
                type 'help'      %.2f-%s

`, Version, stb)
	} else if x == 1 {
		fmt.Printf(`
	__     __                        
	\ \   / /__ _ __   ___ _ __ __ _ 
	 \ \ / / _ \ '_ \ / _ \ '__/ _' |
	  \ V /  __/ | | |  __/ | | (_| |
	   \_/ \___|_| |_|\___|_|  \__,_|	   
	Recon Mission: github.com/venerasf/Venera
	Read the docs https://venera.farinap5.com/
        type 'help'      %.2f-%s

`, Version, stb)
	} else if x == 2 {
		fmt.Printf(`
	██╗   ██╗███████╗███╗   ██╗███████╗██████╗  █████╗ 
	██║   ██║██╔════╝████╗  ██║██╔════╝██╔══██╗██╔══██╗
	██║   ██║█████╗  ██╔██╗ ██║█████╗  ██████╔╝███████║
	╚██╗ ██╔╝██╔══╝  ██║╚██╗██║██╔══╝  ██╔══██╗██╔══██║
	 ╚████╔╝ ███████╗██║ ╚████║███████╗██║  ██║██║  ██║
	  ╚═══╝  ╚══════╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝	
	   Recon Mission: github.com/venerasf/Venera
	   Read the docs https://venera.farinap5.com/
                type 'help'      %.2f-%s

`, Version, stb)
	} else if x == 3 {
		// print("\u001B[1;31m")
		fmt.Printf(`
	██▒   █▓▓█████  ███▄    █ ▓█████  ██▀███   ▄▄▄      
	▓██░   █▒▓█   ▀  ██ ▀█   █ ▓█   ▀ ▓██ ▒ ██▒▒████▄    
	 ▓██  █▒░▒███   ▓██  ▀█ ██▒▒███   ▓██ ░▄█ ▒▒██  ▀█▄  
	  ▒██ █░░▒▓█  ▄ ▓██▒  ▐▌██▒▒▓█  ▄ ▒██▀▀█▄  ░██▄▄▄▄██ 
	   ▒▀█░  ░▒████▒▒██░   ▓██░░▒████▒░██▓ ▒██▒ ▓█   ▓██▒
	   ░ ▐░  ░░ ▒░ ░░ ▒░   ▒ ▒ ░░ ▒░ ░░ ▒▓ ░▒▓░ ▒▒   ▓▒█░
	   ░ ░░   ░ ░  ░░ ░░   ░ ▒░ ░ ░  ░  ░▒ ░ ▒░  ▒   ▒▒ ░
		 ░░     ░      ░   ░ ░    ░     ░░   ░   ░   ▒   
		  ░     ░  ░         ░    ░  ░   ░           ░  ░
		 ░                                               	
	   Recon Mission: github.com/venerasf/Venera
	   Read the docs https://venera.farinap5.com/
                type 'help'      %.2f-%s

`, Version, stb)
		// print("\u001B[0;0m")
	} else if x == 4 {
		fmt.Printf(`
	 _____                        _____                                 _   
	|  |  |___ ___ ___ ___ ___   |   __|___ ___ _____ ___ _ _ _ ___ ___| |_ 
	|  |  | -_|   | -_|  _| .'|  |   __|  _| .'|     | -_| | | | . |  _| '_|
	 \___/|___|_|_|___|_| |__,|  |__|  |_| |__,|_|_|_|___|_____|___|_| |_,_|		
	   Recon Mission: github.com/venerasf/Venera
	   Read the docs https://venera.farinap5.com/
                type 'help'      %.2f-%s

`, Version, stb)
	} else if x == 5 {
		fmt.Printf(`
	╦  ╦┌─┐┌┐┌┌─┐┬─┐┌─┐  ╔═╗┬─┐┌─┐┌┬┐┌─┐┬ ┬┌─┐┬─┐┬┌─
	╚╗╔╝├┤ │││├┤ ├┬┘├─┤  ╠╣ ├┬┘├─┤│││├┤ ││││ │├┬┘├┴┐
	 ╚╝ └─┘┘└┘└─┘┴└─┴ ┴  ╚  ┴└─┴ ┴┴ ┴└─┘└┴┘└─┘┴└─┴ ┴	
	   Recon Mission: github.com/venerasf/Venera
	   Read the docs https://venera.farinap5.com/
                type 'help'      %.2f-%s

`, Version, stb)
	}
}
