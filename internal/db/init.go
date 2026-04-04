package db

import (
	"database/sql"
	"embed"
	"os"
	"venera/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed key
var embedKey embed.FS

func DBInit(homeDir string) DBDef {
	fname := homeDir + "/.venera/database.db"
	_, err := os.Open(fname)
	if err != nil {
		utils.LogMsg(homeDir+"/.venera/message.log", 0, "core", "Creating database")
		println("[+]- Creating database")
		_, err := os.Create(fname)
		if err != nil {
			utils.LogMsg(homeDir+"/.venera/message.log", 3, "core", "Error creating database.")
			utils.PrintErr(err.Error())
			os.Exit(1)
		}
	}

	// Create db definition
	db := new(DBDef)
	utils.LogMsg(homeDir+"/.venera/message.log", 0, "core", "Open database.")
	db.DBConn, err = sql.Open("sqlite3", fname)
	if err != nil {
		utils.LogMsg(homeDir+"/.venera/message.log", 3, "core", "Error while open db func.")
		utils.PrintErr(err.Error())
		os.Exit(1)
	}
	db.dbCreateDs()
	return *db
}

/*
	TODO: make logpath (from utils.LogMsg()) relative.
*/
func (db *DBDef) dbCreateDs() {
	sttm, err := db.DBConn.Prepare(`
	CREATE TABLE IF NOT EXISTS global (
		gid		INTEGER PRIMARY KEY AUTOINCREMENT,
		key 	TEXT UNIQUE,
		value 	TEXT
	)
	`)
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
	} else {
		if _, err := sttm.Exec(); err != nil {
			utils.PrintErr("Failed to create global table: " + err.Error())
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
		}
		sttm.Close()
	}

	sttm, err = db.DBConn.Prepare(`
	CREATE TABLE IF NOT EXISTS Pubkey (
		gid		INTEGER PRIMARY KEY AUTOINCREMENT,
		Author	TEXT UNIQUE,
		Key 	TEXT
	)
	`)
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
	} else {
		if _, err := sttm.Exec(); err != nil {
			utils.PrintErr("Failed to create Pubkey table: " + err.Error())
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
		}
		sttm.Close()
	}

	sttm, err = db.DBConn.Prepare(`
	CREATE TABLE IF NOT EXISTS script (
		sid		INTEGER PRIMARY KEY AUTOINCREMENT,
		hash	VARCHAR(32) UNIQUE,
		path 	TEXT UNIQUE,
		tags	TEXT,
		version REAL,
		description TEXT,
		date DATETIME
	)
	`)
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
	} else {
		if _, err := sttm.Exec(); err != nil {
			utils.PrintErr("Failed to create script table: " + err.Error())
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
		}
		sttm.Close()
	}

	// default root key must be changed and dynamic
	keyBytes, err := embedKey.ReadFile("key/root_pub")
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
	}
	//print(keyBytes)
	keyPack, err := utils.GetKeyFromPack(keyBytes)
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
	}

	sttm, err = db.DBConn.Prepare("INSERT INTO pubkey (Author,Key) VALUES (?,?);")
	if err != nil {
		utils.PrintErr(err.Error())
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
	} else {
		if _, err := sttm.Exec(keyPack.Email, keyPack.Key); err != nil {
			// Key might already exist, log but don't fail
			utils.LogMsg("~/venera/message.log", 1, "core", "Key insert: "+err.Error())
		}
		sttm.Close()
	}
}

func (db *DBDef) DBStoreGlobal(key string, value string) {
	// validate if key exists
	var v string = ""
	row := db.DBConn.QueryRow("SELECT value FROM global WHERE key = ?;", key)
	row.Scan(&v)  // Ignore error - empty result is expected if key doesn't exist

	if v != "" {
		// if key exists we update it
		sttm, err := db.DBConn.Prepare("UPDATE global SET value = ? WHERE key = ?;")
		if err != nil {
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
			utils.PrintErr(err.Error())
			return
		}
		defer sttm.Close()
		
		if _, err := sttm.Exec(value, key); err != nil {
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
			utils.PrintErr("Failed to update global: " + err.Error())
		}
	} else {
		// if not set assing the velue
		sttm, err := db.DBConn.Prepare(`
			INSERT INTO global (key, value) VALUES (?,?);
		`)
		if err != nil {
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
			utils.PrintErr(err.Error())
			return
		}
		defer sttm.Close()
		
		if _, err := sttm.Exec(key, value); err != nil {
			utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
			utils.PrintErr("Failed to insert global: " + err.Error())
		}
	}
}

/*
	DBLoadIntoGlobals: loads the data from database into a map.
	Probably it is gonna be moved to outer package for the case of conflicts from cycling.
*/
func (db *DBDef) DBLoadIntoGlobals() map[string]string {
	g := make(map[string]string)
	row, err := db.DBConn.Query("SELECT key, value FROM global;")
	if err != nil {
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
		utils.PrintErr("Failed to load globals from database: " + err.Error())
		return g  // Return empty map instead of panicking
	}
	defer row.Close()
	
	for row.Next() {
		var k, v string
		if err := row.Scan(&k, &v); err != nil {
			utils.LogMsg("~/venera/message.log", 1, "core", "Error scanning row: "+err.Error())
			continue  // Skip this row and continue
		}
		g[k] = v
	}
	
	// Check for errors during iteration
	if err := row.Err(); err != nil {
		utils.LogMsg("~/venera/message.log", 1, "core", "Error iterating rows: "+err.Error())
	}
	
	return g
}

func (db *DBDef) DBRemoveGlobals(key string) error {
	sttm, err := db.DBConn.Prepare("DELETE FROM global WHERE key = ?;")
	if err != nil {
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
		return err
	}
	defer sttm.Close()
	
	_, err = sttm.Exec(key)
	if err != nil {
		utils.LogMsg("~/venera/message.log", 3, "core", err.Error())
		return err
	}
	return nil
}
