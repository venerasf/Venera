package utils

import "encoding/json"


type KeyPack struct {
	Key   string `json:"key"`
	Email string `json:"email"`
	// Expr  string `json:"expr"`
}

/*
	GetKeyFromPack receives the json and returns the keyPack (key and mail).
*/
func GetKeyFromPack(data []byte) (KeyPack, error) {
	pack := new(KeyPack)
	err := json.Unmarshal(data, pack)
	if err != nil {
		return KeyPack{},err
	}
	return *pack, nil
}