package pacman

import (
	"gopkg.in/yaml.v3"
)

func VPMGetRemotePack(repo string) {
	// Retrive the map package
	yamlBytes, err := DownloadData(repo)
	if err != nil {
		panic(err.Error())
	}

	pack := Pack{}

	err = yaml.Unmarshal(yamlBytes, &pack)
	if err != nil {
		panic(err.Error())
	}

	for i := range(pack.Target) {
		path := pack.Target[i].Path
		_, err := DownloadData(path)
		if err != nil {
			panic(err.Error())
		}
		println(pack.Target[i].Path)
	}
}