
-- Metadata
METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    TAGS = {"example","XSS","scanner"},
    INFO = [[Other Test]]
}

-- Arguments/Variables needed to execute script
VARS = {
    RHOST = {VALUE="0.0.0.0", NEEDED="yes", DESCRIPT="Any"},
    RPORT = {VALUE="4444",    NEEDED="yes", DESCRIPT="Any"},
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    PrintSuccs("OK")
end