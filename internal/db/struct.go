package db

import "database/sql"

type DBDef struct {
	DBConn *sql.DB
}

// Close closes the database connection
func (db *DBDef) Close() error {
	if db.DBConn != nil {
		return db.DBConn.Close()
	}
	return nil
}
