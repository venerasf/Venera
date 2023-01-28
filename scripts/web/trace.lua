-- Script for Venera Framework
-- https://github.com/farinap5/Venera

local http = require("http")
local client = http.client()

-- Metadata
METADATA = {
    AUTHOR = {"farinap5 <null>"},
    VERSION = "1.0",
    TAGS = {"method","http","scanner","xst","trace"},
    INFO = [[Test Cross Site Tracing (xst) with TRACE method enabled.

It is the same as doing by command line with:
curl http://target.com/ -X TRACE -H "Any: abc"
]]
}

-- Arguments/Variables needed to execute script
VARS = {
    URL = {VALUE="http://example.com", NEEDED="yes", DESCRIPT="URL to test XST."},
    VERBOSE = {VALUE="false", NEEDED="no", DESCRIPT="Verbose output."}
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    local randHeader = RandomString(6,"A-Z")
    local randValue  = RandomString(6,"a-z")

    local request = http.request("TRACE", VARS.URL.VALUE)
    request:header_set(randHeader, randValue)
    local r, err = client:do_request(request)
    
    if VARS.VERBOSE.VALUE == "true" then
        print(r.body)
    end

    if r.body:find(randValue,1,true) then
        PrintSuccsln("TRACE method enable.")
    else
        if VARS.VERBOSE.VALUE == "true" then
            PrintErrln("TRACE method not enable")
        end 
    end
end