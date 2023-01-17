local cmd = require("cmd")

-- Metadata
Metadata = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    CATS = {"example","encoding"},
    INFO = [[OD command execution with lua-go]]
}

-- Arguments/Variables needed to execute script
Vars = {
    CMD = {VALUE="ls", NEEDED="yes", DESCRIPT="Command"}
}

function Init()
    Meta(Metadata) -- Load metadata 
    LoadVars(Vars) -- Load variables
end

function Main()
    local result, err = cmd.exec(Vars.CMD.VALUE)
    if err then error(err) end
    PrintSuccs(result.status)
    print(result.stdout)
end