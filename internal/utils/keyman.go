package utils

import "encoding/json"


type keyPack struct {
	Key   string `json:"key"`
	Email string `json:"email"`
	// Expr  string `json:"expr"`
}

/*
	GetKeyFromPack receives the json and returns the keyPack (key and mail).
*/
func GetKeyFromPack(data []byte) (keyPack, error) {
	pack := new(keyPack)
	err := json.Unmarshal(data, pack)
	if err != nil {
		return keyPack{},err
	}
	return *pack, nil
}