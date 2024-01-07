/*
	The package is signed with the creator's private key.
	Each script has the hash computed and it is written to the package.

	First we need to verify the package signature and assume that the script hashes are valid too.

	For each downloaded script, we must match the hash digested with the downloaded bytes. The reference of the 
	script in the package has the atribute "hash" that holds the md5 computed during the package compilation.
	
	https://venera.farinap5.com/6-venera-package-manager.html
*/


package pacman

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"strings"
	"venera/src/db"
	"venera/src/utils"
)

type SignPack struct {
	Author string `json:"Author"`
	Date   string `json:"Date"`
	Sign   string `json:"Sign"`
}

func VerifyPk(r io.Reader,pemEncd []byte, bsign []byte) bool {
	blk, _ := pem.Decode(pemEncd)
	x509Encd := blk.Bytes
	publicKey, err := x509.ParsePKIXPublicKey(x509Encd) // generic key
	if err != nil {
			println(err.Error())
	}
	pk := publicKey.(*ecdsa.PublicKey)
	h := sha256.New()
	_, err = io.Copy(h, r)
	if err != nil {
			fmt.Println(err.Error())
	}
	hash := h.Sum(nil)
	return ecdsa.VerifyASN1(pk, hash, bsign)
}

func GetKeyByEmail(mail string, db *db.DBDef) ([]byte,error) {
	var pkey string
	var err error
	db.DBConn.QueryRow("SELECT key FROM Pubkey WHERE Author = ?;", mail).Scan(&pkey)
	if len(pkey) == 0 {
		err = errors.New("No key for email: "+mail)
	}
	return []byte(pkey), err
}

func VerifySignaturePack(pack []byte, Signp []byte, db db.DBDef) bool {
	p := SignPack{}
	json.Unmarshal(Signp, &p)
	mail := strings.Split(
		strings.Split(p.Author, "<")[1],
		">",
	)[0]
	utils.PrintSuccs("Getting key for author: "+mail)
	pemkey,err := GetKeyByEmail(mail, &db)
	if err != nil {
		utils.PrintErr(err.Error())
	}

	sign, err := base64.StdEncoding.DecodeString(string(p.Sign))
	if err != nil {
		println(err.Error())
	}
	v := VerifyPk(
		bytes.NewReader(pack),
		pemkey,
		sign,
	)
	utils.PrintSuccs("Sign date: "+p.Date)
	if v {
		utils.PrintSuccs("Signed by trusted author: "+mail)
		return true
	} else {
		utils.PrintErr("Invalid signature");
		return false
	}
}


func VerifySignatureScript(data []byte, hash string) bool {
	m := md5.New()
	m.Write(data)
	h := m.Sum(nil)

	if hex.EncodeToString(h) == hash {
		return true
	} else {
		return false
	}
}