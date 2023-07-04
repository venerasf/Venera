package src

import (
	"log"
	"os/user"
	"venera/src/db"
)

var Version	float32
var Stable 	bool

func Start(v float32, stb bool) {
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
		// Start database
		dbdef = db.DBInit(user.Username)
	}

	Version = v
	Stable = stb
	
	// Store global (or refector it) just if it is not setted yet
	dbdef.DBStoreGlobal("chain","on")
	dbdef.DBStoreGlobal("VERBOSE","true")
	dbdef.DBStoreGlobal("user", user.Username)

	// Load persistent global variables ad init map.
	profile.Globals = dbdef.DBLoadIntoGlobals()

	profile.InitCLI()
}