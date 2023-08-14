local regex = require("regexp")

-- Metadata
METADATA = {
    AUTHOR = {"author1 <null>"},
    VERSION = "1.0",
    TAGS = {"example","regex"},
    INFO = [[Regex example]]
}

-- Arguments/Variables needed to execute script
VARS = {
    DATA = {VALUE="root:x:0:0:root:/root:/bin/bash", NEEDED="yes", DESCRIPT="Text pattern."},
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    local r,e = regex.match("root:[x*]:0:0:", VARS.DATA.VALUE)
    if e then error(e) end

    if r == true then
        PrintSuccsln("found")
    end
end