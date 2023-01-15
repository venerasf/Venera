local tcp = require("tcp")

-- Metadata
Metadata = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    CATS = {"example","XSS","scanner"},
    INFO = [[Other Test]]
}

-- Arguments/Variables needed to execute script
Vars = {
    RHOST = {VALUE="0.0.0.0", NEEDED="yes", DESCRIPT="Any"},
    RPORT = {VALUE="4444",    NEEDED="yes", DESCRIPT="Any"},
}

function Init()
    Meta(Metadata) -- Load metadata 
    LoadVars(Vars) -- Load variables
end

function Main()
    PrintSuccs("OK")
end