// This file imports and exports scripts/modules
// to directory scripts/myscripts

package core

import (
	"path/filepath"
	"venera/internal/types"
	"venera/internal/utils"
)

// Import script from somewhere to inside scripts file
func SCImportScript(p types.Profile, pathFrom string, pathTo string) {
	// Validate that pathTo doesn't contain directory traversal
	myscriptsDir := filepath.Clean(p.Globals["myscripts"])
	safePath, err := utils.SafeJoinPath(myscriptsDir, pathTo)
	if err != nil {
		utils.PrintErr("Invalid destination path: " + err.Error())
		return
	}

	// Copy file to safe destination
	finalPath, err := utils.CopyFile(pathFrom, safePath, 0750, 0640)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}

	utils.PrintSuccs("Script imported to: " + finalPath)
}

// export a script
func SCExportScript(p types.Profile, pathFrom string, pathTo string) {
	// Validate pathFrom is within the scripts directory
	scriptsDir := filepath.Clean(p.Globals["root"])
	safePathFrom, err := utils.ValidatePath(pathFrom, scriptsDir)
	if err != nil {
		utils.PrintErr("Invalid source path: " + err.Error())
		return
	}

	// For export, we allow writing to user-specified location
	// but we still clean the path and validate it's absolute
	cleanPathTo := filepath.Clean(pathTo)
	if !filepath.IsAbs(cleanPathTo) {
		// Convert to absolute path if relative
		cleanPathTo, err = filepath.Abs(cleanPathTo)
		if err != nil {
			utils.PrintErr("Invalid export path: " + err.Error())
			return
		}
	}

	// Copy file to export destination
	finalPath, err := utils.CopyFile(safePathFrom, cleanPathTo, 0750, 0640)
	if err != nil {
		utils.PrintErr(err.Error())
		return
	}

	utils.PrintSuccs("Script exported to: " + finalPath)
}
