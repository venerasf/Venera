package pacman

import (
	"database/sql"
	"fmt"
	"venera/internal/db"
	"venera/internal/utils"
)

// the signature block is explained during the sign.go file
// more references in https://venera.farinap5.com/6-venera-package-manager.html

func DownloadScript(dbc *db.DBDef, pack Pack, vnrhome string, i int) int {
	// pack.Target[i] is the remote remote script
	sigStatus := "\u001B[1;32mSigned\u001B[0;0m"

	// Verify if i have the script downloaded, if so, whats the version
	storedS, err := SelectScript(dbc, pack.Target[i])

	/*
		If the script does not exist in the database, it will be installed
		normally.
	*/
	if err != nil && err == sql.ErrNoRows {
		// download
		data, err := DownloadData(pack.Target[i].Path)
		if err != nil {
			utils.PrintErr(err.Error())
			return 3
		}
		matchSignature := VerifySignatureScript(data, pack.Target[i].Hash)
		if !matchSignature {
			utils.PrintAlert("Signature Does Not Match!")
			sigStatus = "\u001B[1;31mSignature error!\u001B[0;0m"
		}
		RegisterScript(dbc, pack.Target[i])
		r := installer(data, vnrhome, pack.Target[i].Script)
		if r == 3 {
			utils.PrintAlert("error.")
			return 3
		} else {
			utils.PrintSuccs(pack.Target[i].Script + " installed. " + sigStatus)
			return 0
		}
	}

	s := pack.Target[i].Script
	if pack.Target[i].Hash == storedS.Hash {
		utils.PrintAlert(s + " already installed and up to date!")
		return 0

	} else if pack.Target[i].Version == storedS.Version {
		utils.PrintAlert(s + " is up to date!")
		return 0

	} else if pack.Target[i].Version > storedS.Version {
		// download
		data, err := DownloadData(pack.Target[i].Path)
		if err != nil {
			utils.PrintErr(err.Error())
			return 3
		}
		matchSignature := VerifySignatureScript(data, pack.Target[i].Hash)
		if !matchSignature {
			utils.PrintAlert("Signature Does Not Match!")
			sigStatus = "\u001B[1;31mSignature error!\u001B[0;0m"
		}

		UpdateScript(dbc, pack.Target[i])
		r := installer(data, vnrhome, pack.Target[i].Script)
		if r == 3 {
			utils.PrintAlert("error.")
			return 3
		} else {
			utils.PrintSuccs(
				fmt.Sprintf("%s updated to %.2f. %s.", pack.Target[i].Script, pack.Target[i].Version, sigStatus),
			)
			return 0
		}
	}
	return 3
}
