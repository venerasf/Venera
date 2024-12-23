package pacman

import (
	"errors"
	"strings"
	"venera/internal/db"
)

func GetKeyByEmail(mail string, db *db.DBDef) ([]byte, error) {
	var pkey string
	var err error
	db.DBConn.QueryRow("SELECT key FROM Pubkey WHERE Author = ?;", mail).Scan(&pkey)
	if len(pkey) == 0 {
		err = errors.New("No key for email: " + mail)
	}
	return []byte(pkey), err
}

func InsertScript(dbc *db.DBDef, t Target) {
	sttm, err := dbc.DBConn.Prepare(`
		INSERT INTO script 
			(hash, path, tags, version, description, date)
		VALUES
			( ?, ?, ?, ?, ?, datetime());
		`)
	if err != nil {
		panic(err.Error())
	}
	_, err = sttm.Exec(t.Hash, t.Script, strings.Join(t.Tags, ":"), t.Version, t.Description)

	if err != nil {
		panic(err.Error())
	}
}

func SelectScript(dbc *db.DBDef, t Target) (Target, error) {
	row := dbc.DBConn.QueryRow(`
		SELECT 
			hash, path, tags, version, description
		FROM script WHERE path=?;
		`, t.Script)

	storeTarget := Target{}
	var tags string
	err := row.Scan(
		&storeTarget.Hash,
		&storeTarget.Path,
		&tags,
		&storeTarget.Version,
		&storeTarget.Description,
	)

	storeTarget.Tags = strings.Split(tags, ":")
	return storeTarget, err
}

func UpdateScript(dbc *db.DBDef, t Target) {
	sttm, err := dbc.DBConn.Prepare(`
	UPDATE script SET
		hash=? path=? tags=? version=? description=? date=datetime())
	VALUES
	`)

	if err != nil {
		panic(err.Error())
	}
	_, err = sttm.Exec(t.Hash, t.Script, strings.Join(t.Tags, ":"), t.Version, t.Description)

	if err != nil {
		panic(err.Error())
	}
}