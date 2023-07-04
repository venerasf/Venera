package db

import (
	"os"
)

func createVeneraDir(uname string) error {
	println("It seems like your first time running it!")
	println("[+]- Setting up ~/.venera")
	return os.Mkdir("/home/"+uname+"/.venera", os.ModePerm)
}

//
func TestVeneraDir(uname string) error {
	_, err := os.Stat("/home/"+uname+"/.venera")
	if err != nil {
		err = createVeneraDir(uname)
		if err != nil {
			println("Fatal error: "+err.Error())
			return err
		}
	}
	return nil
}