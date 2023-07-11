package src

import (
	"log"
	"os/user"
	"venera/src/db"
)

var Version	float32
var Stable 	bool

func Start(v float32, stb bool) {
	Version = v
	Stable = stb

	user, err := user.Current()
	if err != nil {
		log.Println(err.Error())
	}

	// Init profile
	profile := new(Profile)
	// Init database definition
	var dbdef db.DBDef
	// set scripts folder
	profile.BPath = "scripts/"

	// Test vnr home directory
	if db.TestVeneraDir(user.Username) == nil {
		// Start database from home dir
		dbdef = db.DBInit(user.Username)
	}
	profile.Database = &dbdef
	
	/*
	Store global (or refector it) just if it is not setted yet.
	If setted it will be updated. Maybe put in the first interation
	setup.
	*/
	dbdef.DBStoreGlobal("chain","on")
	dbdef.DBStoreGlobal("VERBOSE","true")
	dbdef.DBStoreGlobal("myscripts","myscripts/")
	dbdef.DBStoreGlobal("user", user.Username)

	// Load persistent global variables ad init map.
	profile.Globals = dbdef.DBLoadIntoGlobals()

	profile.InitCLI()
}