// Package constants defines constant values used throughout the Venera Framework
package constants

import "time"

// Directory and file paths
const (
VeneraDirName    = ".venera"
DatabaseFileName = "database.db"
LogFileName      = "message.log"
ScriptsDirName   = "scripts"
MyScriptsDirName = "myscripts"
)

// HTTP timeouts and configurations
const (
HTTPDialTimeout            = 5 * time.Second
HTTPKeepAlive              = 5 * time.Second
HTTPTLSHandshakeTimeout    = 3 * time.Second
HTTPResponseHeaderTimeout  = 3 * time.Second
HTTPExpectContinueTimeout  = 1 * time.Second
)

// Repository URLs
const (
DefaultRepoURL = "http://r.venera.farinap5.com/package.yaml"
DefaultSignURL = "http://r.venera.farinap5.com/package.sgn"
)

// Display and formatting
const (
MaxInfoDisplayLength = 25  // Maximum length for info display before truncation
)

// File size limits
const (
MaxLuaFileReadSize = 10 * 1024 * 1024  // 10 MB limit for Lua Open() function
)

// File permissions
const (
VeneraDirPermissions = 0750  // Owner: rwx, Group: r-x, Other: none
LogFilePermissions   = 0640  // Owner: rw-, Group: r--, Other: none
ScriptDirPermissions = 0700  // Owner: rwx, Group: none, Other: none
)
