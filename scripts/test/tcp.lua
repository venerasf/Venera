local tcp = require("tcp")

METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    TAGS = {"example","tcp","scanner"},
    INFO = [[TCP conn with go-lua]]
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
    local host = VARS.RHOST.VALUE..":"..VARS.RPORT.VALUE
    local conn, err = tcp.open(host)
    err = conn:write("GET /\n\n")
    if err then error(err) end
    local result, err = conn:read(64*1024)
    print(result)
end