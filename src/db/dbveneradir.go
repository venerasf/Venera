package db

import (
	"os"
)

func createVeneraDir(homeDir string) error {
	println("It seems like your first time running it!")
	println("[+]- Setting up ~/.venera")
	return os.Mkdir(homeDir+"/.venera", os.ModePerm)
}

//
func TestVeneraDir(homeDir string) error {
	_, err := os.Stat(homeDir+"/.venera")
	/*
		Many problems can occur but for now this unique.
	*/
	if err != nil {
		err = createVeneraDir(homeDir)
		if err != nil {
			println("Fatal error: "+err.Error())
			return err
		}
	}
	return nil
}