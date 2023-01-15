local tcp = require("tcp")

Metadata = {
    Metadata = {
        AUTHOR = {"Author1 <author1@mail.com>"},
        VERSION = "0.1",
        CATS = {"example","tcp","http"},
        INFO = [[Simple HTTP made with raw tcp connection.]]
    }
}

-- Arguments/Variables needed to execute script
Vars = {
    RHOST = {VALUE="0.0.0.0", NEEDED="yes", DESCRIPT="Remote Host"},
    RPORT = {VALUE="4444",    NEEDED="yes", DESCRIPT="Remote TCP port"},
}

function Init()
    Meta(Metadata) -- Load metadata 
    LoadVars(Vars) -- Load variables
end

function Main()
    local conn, err = tcp.open(Vars.RHOST.VALUE+":"+Vars.RPORT.VALUE)
    err = conn:write("GET /\n\n")
    
    if err then error(err) end
    local result, err = conn:read(64*1024)
    print(result)
end