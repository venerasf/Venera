package src

import (
	"log"
	"os/user"
	"venera/src/db"
	"venera/src/utils"
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
	// Init database definition its a pointer.
	var dbdef db.DBDef
	// set scripts folder
	//profile.BPath = "scripts/" // now taken from globals[root]

	// Test vnr home directory
	if db.TestVeneraDir(user.HomeDir) == nil {
		// Start database from home dir
		dbdef = db.DBInit(user.HomeDir)
	}
	// profile receives the database, so it can perform actions anywhere
	profile.Database = &dbdef
	
	/*
	Store global (or reset it) just if it is not setted yet.
	If setted it will be updated. Maybe put in the first interation
	setup.

	see the https://farinap5.github.io/venera/Global%20Variables/
	*/
	dbdef.DBStoreGlobal("chain","on")
	dbdef.DBStoreGlobal("VERBOSE","true")
	dbdef.DBStoreGlobal("myscripts",user.HomeDir+"/.venera/scripts/myscripts/")
	dbdef.DBStoreGlobal("logfile",user.HomeDir+"/.venera/message.log")
	dbdef.DBStoreGlobal("user", user.Username)
	dbdef.DBStoreGlobal("home", user.HomeDir)
	dbdef.DBStoreGlobal("root",user.HomeDir+"/.venera/scripts")
	dbdef.DBStoreGlobal("repo","http://r.venera.farinap5.com/package.yaml")
	dbdef.DBStoreGlobal("sign","http://r.venera.farinap5.com/package.sgn")
	dbdef.DBStoreGlobal("vpmvs","true")

	// Load persistent global variables to the map.
	// It can be taken typing `globals` on prompt.
	profile.Globals = dbdef.DBLoadIntoGlobals()
	utils.LogMsg(profile.Globals["logfile"],0,"core","Startup initialized.")

	// Init prompt
	profile.InitCLI()
}