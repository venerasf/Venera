package pacman

import (
	"errors"
	"strings"
	"venera/internal/db"
	"venera/internal/utils"
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

// Register new key
func RegisterKey(dbc *db.DBDef, keypack utils.KeyPack) error {
	sttm, err := dbc.DBConn.Prepare(`
		INSERT INTO Pubkey 
			(Author, Key)
		VALUES
			( ?, ?);
	`)
	if err != nil {
		return err
	}
	_, err = sttm.Exec(keypack.Email, keypack.Key)

	if err != nil {
		return err
	}
	return nil
}

// Register a new script
func RegisterScript(dbc *db.DBDef, t Target) {
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

func GetRegisteredKeys(dbc *db.DBDef) ([]utils.KeyPack, error) {
	rows, err := dbc.DBConn.Query(`
		SELECT 
			gid, Author, Key
		FROM Pubkey;
	`)
	
	if err != nil {
		return nil,err
	}

	data := []utils.KeyPack{}
	for rows.Next() {
		var id int
		var a,k string
		err := rows.Scan(&id, &a, &k)
		if err != nil {return nil,err}
		data = append(data, utils.KeyPack{Id: id, Email: a, Key: k})
	}

	return data, nil
}