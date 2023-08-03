package pacman

import (
	"fmt"
	"os"
	"venera/src/utils"
	"strings"

	"gopkg.in/yaml.v2"
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
	utils.PrintSuccs("Requesting "+repo+"\n")
	if validateTarget(pack) != 0 {
		utils.PrintAlert("No data to show.")
	}

	utils.PrintSuccs(fmt.Sprintf("%d scripts found.",len(pack.Target)))

	for i := range(pack.Target) {
		if strings.Contains(pack.Target[i].Description,pattern) ||
		strings.Contains(pack.Target[i].Script,pattern) {
			if i > 0 {
				print("-----------------------\n")
			}
			fmt.Printf("Script: 	%s\n",pack.Target[i].Script)
			fmt.Printf("Version:	%f\n",pack.Target[i].Version)
			fmt.Printf("Decription:	%s\n",pack.Target[i].Description)
			fmt.Printf("Tags:		",)
			for j := range(pack.Target[i].Tags) {
				if j != 0 {
					print(", ")
				}
				print(pack.Target[i].Tags[j])
			}
			print("\n")
		}
	}
}

func installer(data []byte, vnrhome string, scriptPath string) int {
	scriptPath = strings.TrimPrefix(scriptPath,"/")
	pathSplit := strings.Split(scriptPath,"/")
	path := strings.Join(pathSplit[:len(pathSplit)-1],"/")
	/*TODO:
		Change the permissions after tests.
	*/
	err := os.MkdirAll(vnrhome+path,0700)
	if err != nil {
		utils.PrintErr(err.Error())
		return 3
	}
	file,err := os.Create(vnrhome+"/"+scriptPath)
	if err != nil {
		utils.PrintErr(err.Error())
		return 3
	}
	file.Write(data)
	file.Close()
	return 0
}

func sync(repo, vnrhome string) int {
	utils.PrintSuccs("Requesting "+repo+"\n")
	pack := getPack(repo)

	v := validateTarget(pack)
	if v != 0 {
		return v
	}

	for i := range(pack.Target) {
		utils.PrintAlert("Intalling "+pack.Target[i].Script)
		b, err := DownloadData(pack.Target[i].Script)
		if err != nil {
			utils.PrintErr("Error downloading script:"+err.Error())
		} else {
			r := installer(b, vnrhome, pack.Target[i].Script)
			if r == 0 {
				utils.PrintSuccs(pack.Target[i].Script+" installed.")
			} else if r == 1 {
				utils.PrintSuccs(pack.Target[i].Script+" updated.")
			} else if r == 3 {
				utils.PrintAlert("error.")
			}
		}
	}
	return 0
}

func VPMGetRemotePack(repo string, vnrhome string, args []string) int {
	if len(args) > 2 && args[1] == "search" {
		search(repo, args[2])
	} else if len(args) > 2 && args[1] == "install" {
		utils.PrintSuccs("Requesting "+repo+"\n")
		pack := getPack(repo)
		for i := range(pack.Target) {
			if pack.Target[i].Script == args[2] {
				data,err := DownloadData(pack.Target[i].Path)
				if err != nil {
					utils.PrintErr(err.Error())
				} else {
					r := installer(data, vnrhome, pack.Target[i].Script)
					if r == 0 {
						utils.PrintSuccs(pack.Target[i].Script+" installed.")
						return 0
					} else if r == 1 {
						utils.PrintSuccs(pack.Target[i].Script+" updated.")
						return 1
					} else if r == 2 {
						utils.PrintAlert("No data script found.")
						return 2
					} else if r == 3 {
						utils.PrintAlert("error.")
						return 3
					}
				}
			}
		}
	} else if len(args) > 1 && args[1] == "sync" {
		sync(repo, vnrhome)
	} else {
		utils.PrintAlert("No arg")
	}
	return 0
}
