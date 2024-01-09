package db

import (
	"database/sql"
	"os"
	"venera/src/utils"

	_ "github.com/mattn/go-sqlite3"
)

func DBInit(homeDir string) DBDef {
	fname := homeDir+"/.venera/database.db"
	_,err := os.Open(fname)
	if err != nil {
		utils.LogMsg(homeDir+"/.venera/message.log",0,"core","Creating database")
		println("[+]- Creating database")
		_,err := os.Create(fname)
		if err != nil {
			utils.LogMsg(homeDir+"/.venera/message.log",3,"core","Error creating database.")
			utils.PrintErr(err.Error())
			os.Exit(1)
		}
	}

	// Create db definition
	db := new(DBDef)
	utils.LogMsg(homeDir+"/.venera/message.log",0,"core","Open database.")
	db.DBConn, err = sql.Open("sqlite3", fname)
	if err != nil {
		utils.LogMsg(homeDir+"/.venera/message.log",3,"core","Error while open db func.")
		utils.PrintErr(err.Error())
		os.Exit(1)
	}
	db.dbCreateDs()
	return *db
}

/* 
	TODO: make logpath (from utils.LogMsg()) relative.
*/
func (db *DBDef)dbCreateDs() {
	sttm,err := db.DBConn.Prepare(`
	CREATE TABLE IF NOT EXISTS global (
		gid		INTEGER PRIMARY KEY AUTOINCREMENT,
		key 	TEXT UNIQUE,
		value 	TEXT
	)
	`)
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log",3,"core",err.Error())
	} else {
		sttm.Exec()
	}
	sttm,err = db.DBConn.Prepare(`
	CREATE TABLE IF NOT EXISTS Pubkey (
		gid		INTEGER PRIMARY KEY AUTOINCREMENT,
		Author	TEXT UNIQUE,
		Key 	TEXT
	)
	`)
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log",3,"core",err.Error())
	} else {
		sttm.Exec()
	}

	// default root key must be changed and dynamic
	key := `-----BEGIN PUBLIC KEY-----
MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQADJX13tbFJYlQ0aWG6gHTZqJ6dLg3
/n/Z/aoUdROOrRfvKGNdgxw0IOH8EetWADU7zcZFd65+wMqV+x4iM2SsBBUAx6U2
vqaHn8ubrE+Z0GZtAAMR9Wusar4pjFS9G98XhILLPfzgZTCtY4BOpfenL+gqg/GT
euivf5/tEQVeHt9f+MQ=
-----END PUBLIC KEY-----`
	sttm,err = db.DBConn.Prepare("INSERT INTO pubkey (Author,Key) VALUES (?,?);")
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log",3,"core",err.Error())
	} else {
		sttm.Exec("elf@mail.com",key)
	}
}


func (db *DBDef)DBStoreGlobal(key string, value string) {
	// validate if key exists
	var v string = ""
	row := db.DBConn.QueryRow("SELECT value FROM global WHERE key = ?;", key)
	row.Scan(&v)

	if v != "" {
		// if key exists we update it
		sttm, err := db.DBConn.Prepare("UPDATE global SET value = ? WHERE key = ?;")
		if err != nil {
			utils.LogMsg("~/venera/message.log",3,"core",err.Error())
			utils.PrintErr(err.Error())
		}
		sttm.Exec(value,key)
	} else {
		// if not set assing the velue
		sttm, err := db.DBConn.Prepare(`
			INSERT INTO global (key, value) VALUES (?,?);
		`)
		if err != nil {
			utils.LogMsg("~/venera/message.log",3,"core",err.Error())
			utils.PrintErr(err.Error())
		}
		sttm.Exec(key,value)
	}
}


/*
	DBLoadIntoGlobals: loads the data from databaso into a map.
	Probably it is gonna be moved to outter package for the case of conflicts for cycling.
*/
func (db *DBDef)DBLoadIntoGlobals() map[string]string {
	g := make(map[string]string)
	row, err := db.DBConn.Query("SELECT key, value FROM global;")
	if err != nil {
		utils.LogMsg("~/venera/message.log",3,"core",err.Error())
		panic(err.Error())
	}
	for row.Next() {
		var k,v string
		row.Scan(&k,&v)
		g[k]=v
	}
	return g
}

func (db *DBDef)DBRemoveGlobals(key string) {
	sttm, err := db.DBConn.Prepare("DELETE FROM global WHERE key = ?;")
	if err != nil {
		utils.LogMsg("~/venera/message.log",3,"core",err.Error())
		panic(err.Error())
	}
	sttm.Exec(key)
}

