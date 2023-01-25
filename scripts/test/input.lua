-- Metadata
METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>",
                "Author2 <author2@mail.com>",
                "Author3 <author3@mail.com>"
            },
    VERSION = "0.1",
    TAGS = {"example","code","lua"},
    INFO = [[Input function]]
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
    local x = Input("aaaa> ")
    PrintSuccs(x)
end