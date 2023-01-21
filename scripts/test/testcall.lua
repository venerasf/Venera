-- Metadata
METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    TAGS = {"example","call"},
    INFO = [[Script to test calls]]
}

-- Arguments/Variables needed to execute script
VARS = {
    URL = {VALUE="http://example.com", NEEDED="yes", DESCRIPT="URL"},
    METHOD = {VALUE="GET", NEEDED="yes", DESCRIPT="METHOD"}
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    Call("scripts/test/http.lua")
end