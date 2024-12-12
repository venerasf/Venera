package db

import (
	"os"
	"venera/internal/utils"
)

func createVeneraDir(homeDir string) error {
	println("It seems like your first time running it!")
	println("After all finished setup process, please type `vpm sync` to synchronize packages.")
	utils.PrintSuccs("Setting up ~/.venera")
	return os.Mkdir(homeDir+"/.venera", os.ModePerm)
}

/*
	TestVeneraDir will test the directory where venera will store everything
*/
func TestVeneraDir(homeDir string) error {
	_, folderExist := os.Stat(homeDir+"/.venera")
	/*
		Many problems can occur but for now this unique will be validated.
	*/
	if folderExist != nil {
		err := createVeneraDir(homeDir)
		if err != nil {
			utils.PrintAlert("Fatal error: "+err.Error())
			return err
		}
	}
	return folderExist
}