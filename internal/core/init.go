package core

import (
	"log"
	"os/user"
	"path/filepath"
	"venera/internal/constants"
	"venera/internal/db"
	"venera/internal/types"
	"venera/internal/utils"
)

var Version float32
var Stable bool

func SetDefaultGlobals(dbdef *db.DBDef, user *user.User) {
	veneraDir := filepath.Join(user.HomeDir, constants.VeneraDirName)
	scriptsDir := filepath.Join(veneraDir, constants.ScriptsDirName)
	myScriptsDir := filepath.Join(scriptsDir, constants.MyScriptsDirName)
	logFile := filepath.Join(veneraDir, constants.LogFileName)

	dbdef.DBStoreGlobal("chain", "on")
	dbdef.DBStoreGlobal("VERBOSE", "true")
	dbdef.DBStoreGlobal("myscripts", myScriptsDir+"/")
	dbdef.DBStoreGlobal("logfile", logFile)
	dbdef.DBStoreGlobal("user", user.Username)
	dbdef.DBStoreGlobal("home", user.HomeDir)
	dbdef.DBStoreGlobal("root", scriptsDir)
	dbdef.DBStoreGlobal("repo", constants.DefaultRepoURL)
	dbdef.DBStoreGlobal("sign", constants.DefaultSignURL)
	dbdef.DBStoreGlobal("vpmvs", "true")
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
	utils.LogMsg(profile.Globals["logfile"], 0, "core", "Startup initialized.")

	// Init prompt
	InitCLI(profile)
}
