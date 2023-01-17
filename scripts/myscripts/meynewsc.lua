-- Metadata
Metadata = {
    AUTHOR = {"Author1 <author1@mail.com>",
                "Author2 <author2@mail.com>",
                "Author3 <author3@mail.com>"
            },
    VERSION = "0.1",
    CATS = {"example","XSS","scanner"},
    INFO = [[IE test]]
}

-- Arguments/Variables needed to execute script
Vars = {
    RHOST = {VALUE="0.0.0.0", NEEDED="yes", DESCRIPT="Any"},
    RPORT = {VALUE="4444",    NEEDED="yes", DESCRIPT="Any"},
    TARGET = {VALUE="8.8.8.8", NEEDED="yse", DESCRIPT="anyyyy"}
}

function Init()
    Meta(Metadata) -- Load metadata 
    LoadVars(Vars) -- Load variables
end

function Main()
    PrintSuccs("OK")
    PrintSuccs(Vars.TARGET.VALUE)
end
