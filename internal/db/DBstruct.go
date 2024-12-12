package db

import "database/sql"

type DBDef struct {
	DBConn *sql.DB
}