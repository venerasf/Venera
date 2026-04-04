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
func RegisterScript(dbc *db.DBDef, t Target) error {
	sttm, err := dbc.DBConn.Prepare(`
		INSERT INTO script 
			(hash, path, tags, version, description, date)
		VALUES
			( ?, ?, ?, ?, ?, datetime());
		`)
	if err != nil {
		utils.PrintErr("Failed to prepare script registration: " + err.Error())
		return err
	}
	defer sttm.Close()
	
	_, err = sttm.Exec(t.Hash, t.Script, strings.Join(t.Tags, ":"), t.Version, t.Description)
	if err != nil {
		utils.PrintErr("Failed to register script: " + err.Error())
		return err
	}
	return nil
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

func UpdateScript(dbc *db.DBDef, t Target) error {
	sttm, err := dbc.DBConn.Prepare(`
	UPDATE script SET
		hash=?, path=?, tags=?, version=?, description=?, date=datetime()
	WHERE path=?
	`)

	if err != nil {
		utils.PrintErr("Failed to prepare script update: " + err.Error())
		return err
	}
	defer sttm.Close()
	
	_, err = sttm.Exec(t.Hash, t.Script, strings.Join(t.Tags, ":"), t.Version, t.Description, t.Script)
	if err != nil {
		utils.PrintErr("Failed to update script: " + err.Error())
		return err
	}
	return nil
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

func DelRegisteredKeys(dbc *db.DBDef, Email string) error {
	sttm, err := dbc.DBConn.Prepare(`DELETE FROM Pubkey WHERE Author=?;`)

	if err != nil {
		return err
	}
	_, err = sttm.Exec(Email)
	return err
}