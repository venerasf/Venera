local http = require("http")
local client = http.client()

-- Metadata
METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    TAGS = {"example","http"},
    INFO = [[HTTP requests with lua-go]]
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
    local request = http.request(VARS.METHOD.VALUE, VARS.URL.VALUE)
    local result, err = client:do_request(request)
    
    if err then
        error(err)
    end

    PrintSuccsln(result.code)
    PrintSuccsln(result.body)
end