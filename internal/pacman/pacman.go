package pacman

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"venera/internal/db"
	"venera/internal/utils"
)


func validateTarget(pack Pack) int {
	if len(pack.Target) == 0 || pack.Target == nil {
		return 2
	} else {
		return 0
	}
}

func getPack(repo string) Pack {
	pack := Pack{}
	yamlBytes, err := DownloadData(repo)
	if err != nil {
		utils.PrintErr(err.Error())
		return pack
	}

	err = yaml.Unmarshal(yamlBytes, &pack)
	if err != nil {
		utils.PrintErr(err.Error())
		return pack
	}
	return pack
}

func search(repo, pattern string) {
	// Retrive the map package
	pack := getPack(repo)
	utils.PrintSuccs("Requesting " + repo + "\n")
	if validateTarget(pack) != 0 {
		utils.PrintAlert("No data to show.")
	}

	utils.PrintSuccs(fmt.Sprintf("%d scripts found.", len(pack.Target)))

	c := 0
	for i := range pack.Target {
		if strings.Contains(pack.Target[i].Description, pattern) ||
			strings.Contains(pack.Target[i].Script, pattern) {
			c++
			if i > 0 {
				print("-----------------------\n")
			}
			fmt.Printf("Script: 	%s\n", pack.Target[i].Script)
			fmt.Printf("Version:	%.2f\n", pack.Target[i].Version)
			fmt.Printf("Description:%s\n", pack.Target[i].Description)
			fmt.Printf("Tags:		")
			for j := range pack.Target[i].Tags {
				if j != 0 {
					print(", ")
				}
				print(pack.Target[i].Tags[j])
			}
			print("\n")
		}
	}
	utils.PrintSuccs(fmt.Sprintf("%d scripts.", c))
}

func installer(data []byte, vnrhome string, scriptPath string) int {
	returnInfo := 0
	// Normalize path
	scriptPath = strings.TrimPrefix(scriptPath, "/")
	pathSplit := strings.Split(scriptPath, "/")
	path := strings.Join(pathSplit[:len(pathSplit)-1], "/")

	_, err := os.Stat(vnrhome + "/" + scriptPath)
	if err == nil {
		returnInfo = 1
	}

	/*TODO:
	Change the permissions after tests.
	*/
	if strings.Split(path, "")[0] != "/" {
		path = "/" + path
	}
	err = os.MkdirAll(vnrhome+path, 0700)
	if err != nil {
		utils.PrintErr(err.Error())
		return 3
	}

	file, err := os.Create(vnrhome + "/" + scriptPath)
	if err != nil {
		utils.PrintErr(err.Error())
		return 3
	}

	file.Write(data)
	file.Close()

	return returnInfo
}

func sync(dbc *db.DBDef, repo, vnrhome string) int {
	utils.PrintSuccs("Requesting " + repo + "\n")
	pack := getPack(repo)

	v := validateTarget(pack)
	if v != 0 {
		return v
	}

	for i := range pack.Target {
		utils.PrintAlert("Installing " + pack.Target[i].Script)

		DownloadScript(dbc, pack, vnrhome, i)
	}
	return 0
}

/*
It is currently verifying the signature just from the main .yaml.
Script isn't verified by itself.

Usefull when you configured a new package repo.
*/
func verifySign(repo string, signRepo string, database *db.DBDef) {
	yamlBytes, err := DownloadData(repo)
	if err != nil {
		utils.PrintErr(err.Error())
	}
	signBytes, err := DownloadData(signRepo)
	if err != nil {
		utils.PrintErr(err.Error())
	}
	VerifySignaturePack(yamlBytes, signBytes, *database)
}

func installCommand(dbc *db.DBDef, repo string, args []string, vnrhome string) int {
	utils.PrintSuccs("Requesting " + repo + "\n")
	pack := getPack(repo)
	for i := range pack.Target {
		if pack.Target[i].Script == args[2] {
			DownloadScript(dbc, pack, vnrhome, i)
		}
	}
	return 0
}

/*
VPMGetRemotePack is the entrypoint for using Venera Package Manager

The following exemplifies the way to call it.
pacman.VPMGetRemotePack(
	profile.Globals["repo"],  	http://r.venera.farinap5.com/package.yaml
	profile.Globals["root"],  	root where to place scripts
	profile.Globals["sign"],  	http://r.venera.farinap5.com/package.sgn
	cmds, 					  	The command like "install", "sync"...
	*profile.Database,        	Database interface
	profile.Globals["vpmvs"], 	If verification is on or off
	profile.Globals["logfile"],	Log path
)

*/

func VPMGetRemotePack(repo string, vnrhome string, signRepo string, args []string, database db.DBDef, verify string, logfile string) int {
	if len(args) < 2 {
		utils.PrintAlert("Type `help vpm`.")
		return 1
	}
	
	switch args[1] {
	case "search":
		if len(args) < 3 {
			utils.PrintAlert("search needs more arguments.")
		} else {
			search(repo, args[2])
			return 0
		}

	case "install":
		if len(args) < 3 {
			utils.PrintAlert("search needs more arguments.")
		} else {
			utils.LogMsg(logfile, utils.INF ,"vmp","install from "+repo+" requested.")
			installCommand(&database, repo, args, vnrhome)
		}

	case "sync":
		utils.LogMsg(logfile,0,"vmp","sync with "+repo+" requested.")
		n := sync(&database, repo, vnrhome)
		if n != 0 {
			utils.LogMsg(logfile,1,"vmp","Sync error reported.")
		}

	case "verify":
		verifySign(repo, signRepo, &database)
		return 0

	default:
		utils.PrintAlert("Type `help vpm`.")
	}

	return 1
}
