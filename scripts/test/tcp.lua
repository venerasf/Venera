local tcp = require("tcp")

Metadata = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    CATS = {"example","tcp","scanner"},
    INFO = [[TCP conn with go-lua]]
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
    local host = Vars.RHOST.VALUE..":"..Vars.RPORT.VALUE
    local conn, err = tcp.open(host)
    err = conn:write("GET /\n\n")
    if err then error(err) end
    local result, err = conn:read(64*1024)
    print(result)
end