package db

import (
	"database/sql"
	"os"
	"venera/src/utils"

	_ "github.com/mattn/go-sqlite3"
)

/*
	TODO: change `panic()` to write the output to the log file
*/
func DBInit(homeDir string) DBDef {
	fname := homeDir+"/.venera/database.db"
	_,err := os.Open(fname)
	if err != nil {
		utils.LogMsg(homeDir+"/.venera/message.log",0,"core","Creating database")
		println("[+]- Creating database")
		_,err := os.Create(fname)
		if err != nil {
			utils.LogMsg(homeDir+"/.venera/message.log",3,"core","Error creating database.")
			panic(err.Error())
		}
	}

	// Create db definition
	db := new(DBDef)
	utils.LogMsg(homeDir+"/.venera/message.log",0,"core","Open database.")
	db.DBConn, err = sql.Open("sqlite3", fname)
	if err != nil {
		utils.LogMsg(homeDir+"/.venera/message.log",3,"core","Error while open db func.")
		panic(err.Error())
	}
	db.dbCreateDs()
	return *db
}

/* 
	TODO: change `panic()` to write the output to the log file
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
		panic(err.Error())
	} else {
		sttm.Exec()
	}
}

/* 
	TODO: change `panic()` to write the output to the log file
*/
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
			panic(err.Error())
		}
		sttm.Exec(value,key)
	} else {
		// if not set assing the velue
		sttm, err := db.DBConn.Prepare(`
			INSERT INTO global (key, value) VALUES (?,?);
		`)
		if err != nil {
			utils.LogMsg("~/venera/message.log",3,"core",err.Error())
			panic(err.Error())
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