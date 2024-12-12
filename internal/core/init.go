package core

import (
	"log"
	"os/user"
	"venera/internal/db"
	"venera/internal/utils"
	"venera/internal/types"
)

var Version	float32
var Stable 	bool


func SetDefaultGlobals(dbdef *db.DBDef, user *user.User) {
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
}

func Start(v float32, stb bool) {
	Version = v
	Stable = stb

	user, err := user.Current()
	if err != nil {
		log.Println(err.Error())
	}

	// Init profile
	profile := new(types.Profile)
	// Init database definition its a pointer.
	var dbdef db.DBDef
	// set scripts folder
	//profile.BPath = "scripts/" // now taken from globals[root]

	// Test vnr home directory
	vnrdir := db.TestVeneraDir(user.HomeDir)
	dbdef = db.DBInit(user.HomeDir)
	
	if vnrdir != nil {
		SetDefaultGlobals(&dbdef, user)
	}

	// profile receives the database, so it can perform actions anywhere
	profile.Database = &dbdef
	
	// Load persistent global variables to the map.
	// It can be taken typing `globals` on prompt.
	profile.Globals = dbdef.DBLoadIntoGlobals()
	utils.LogMsg(profile.Globals["logfile"],0,"core","Startup initialized.")

	// Init prompt
	InitCLI(profile)
}