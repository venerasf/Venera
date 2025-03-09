package pacman

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/cheynewallace/tabby"
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

func search(dbc *db.DBDef ,repo, pattern string) {
	// Retrieve the map package
	pack := getPack(repo)
	utils.PrintSuccs("Requesting " + repo + "\n")
	if validateTarget(pack) != 0 {
		utils.PrintAlert("No data to show.")
	}

	utils.PrintSuccs(fmt.Sprintf("%d scripts found.", len(pack.Target)))

	c := 0
	t := tabby.New()
	t.AddHeader("SCRIPT","VERSION","DESCRIPTION","TAGS","STATUS")
	for i := range pack.Target {
		if strings.Contains(pack.Target[i].Description, pattern) ||
			strings.Contains(pack.Target[i].Script, pattern) || pattern == "all" {

			scriptStatus := "any"
			localS, err := SelectScript(dbc, pack.Target[i])
			if err == sql.ErrNoRows {
				scriptStatus = "download"
			} else if localS.Version < pack.Target[i].Version {
				scriptStatus = "\u001B[1;31mupdate\u001B[0;0m"
			} else if localS.Version >= pack.Target[i].Version {
				scriptStatus = "\u001B[1;32mup to date\u001B[0;0m"
			}

			c++
			desc := ""
			if len(pack.Target[i].Description) > 30 {
				desc = pack.Target[i].Description[:30]+"..."
			} else {
				desc = pack.Target[i].Description
			}
			t.AddLine(
				pack.Target[i].Script,
				pack.Target[i].Version,
				desc,
				strings.Join(pack.Target[i].Tags, ":"),
				scriptStatus,
			)

			/*
			if i > 0 {
				print("-----------------------\n")
			}
			fmt.Printf("Script: 	%s\n", pack.Target[i].Script)
			fmt.Printf("Version:	%.2f\n", pack.Target[i].Version)
			fmt.Printf("Description:%s\n", pack.Target[i].Description)
			fmt.Printf("Tags:		%s\n", strings.Join(pack.Target[i].Tags, ":"))
			fmt.Printf("Status:		%s\n", scriptStatus)
			*/
		}
	}
	print("\n")
	t.Print()
	print("\n")
	utils.PrintSuccs(fmt.Sprintf("%d scripts listed.", c))
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
		DownloadScript(dbc, pack, vnrhome, i)
	}
	return 0
}

/*
It is currently verifying the signature just from the main .yaml.
Script isn't verified by itself.

Useful when you configured a new package repo.
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
			utils.PrintAlert("vpm needs more arguments.")
		} else {
			search(&database, repo, args[2])
			return 0
		}

	case "install":
		if len(args) < 3 {
			utils.PrintAlert("vpm needs more arguments.")
		} else {
			utils.LogMsg(logfile, utils.INF ,"vmp","install from "+repo+" requested.")
			installCommand(&database, repo, args, vnrhome)
		}

	case "sync":
		n := sync(&database, repo, vnrhome)
		if n != 0 {
			utils.LogMsg(logfile,utils.ERR,"vmp","sync error reported for repo " + repo)
		} else {
			utils.LogMsg(logfile,utils.INF,"vmp","sync with " + repo + " requested")
		}

	case "verify":
		verifySign(repo, signRepo, &database)
		return 0

	case "key":
		if len(args) < 3 {
			utils.PrintAlert("vpm needs more arguments.")
			return 1
		}

		if len(args) < 4 && (args[2] == "i" || args[2] == "import") {
			err := RegisterKeyFromFile(&database, args[3])
			if err != nil {
				utils.PrintAlert(err.Error())
				return 1
			}
			utils.PrintSuccs("New key imported.")
			utils.LogMsg(logfile, 0, "vmp", "imported key from file " + args[3])

		} else if len(args) > 3 && args[2] == "del" {
			err := DeleteRegisKey(&database, args[3])
			if err != nil {
				utils.PrintAlert(err.Error())
				return 1
			}
			utils.PrintSuccs("Key deleted ", args[3])
			utils.LogMsg(logfile, 0, "vmp", "Key " + args[3] + " deleted.")

		} else if args[2] == "s" || args[2] == "show" {
			ShowKeys(&database)
		} else {
			utils.PrintAlert("Type `help vpm`.")
		}

	default:
		utils.PrintAlert("Type `help vpm`.")
	}

	return 1
}
