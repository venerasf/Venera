local cmd = require("cmd")

-- Metadata
METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>"},
    VERSION = "0.1",
    TAGS = {"example","cmd"},
    INFO = [[OD command execution with lua-go]]
}

-- Arguments/Variables needed to execute script
VARS = {
    CMD = {VALUE="ls", NEEDED="yes", DESCRIPT="Command"}
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    local result, err = cmd.exec(VARS.CMD.VALUE)
    if err then error(err) end
    PrintSuccs(result.status)
    print(result.stdout)
end