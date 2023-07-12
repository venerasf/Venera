// This file imports and exports scripts/modules
// to directory scripts/myscripts

package src

import (
	"io/ioutil"
	"os"
	"venera/src/utils"
)

// Import script from somewhere to inside scripts file
func (p Profile)SCImportScript(pathFrom string, pathTo string) {
	cont, err := ioutil.ReadFile(pathFrom)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}
	file, err :=  os.Create(p.Globals["myscripts"] +pathTo)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}

	_, err = file.Write(cont)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}
	err = file.Close()
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}
}

// export a script
func (p Profile)SCExportScript(pathFrom string, pathTo string) {
	cont, err := ioutil.ReadFile(pathFrom)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}
	fileTo, err :=  os.Create(pathTo)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}

	_, err = fileTo.Write(cont)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}
	err = fileTo.Close()
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}
}