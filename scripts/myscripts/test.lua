METADATA = {
    AUTHOR = {"Author1 <author1@mail.com>",
                "Author2 <author2@mail.com>",
                "Author3 <author3@mail.com>"
            },
    VERSION = "0.1",
    TAGS = {"example","test"},
    INFO = [[Lorem ipsum dolor sit amet, consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris
nisi ut aliquip ex ea commodo consequat.]]
}

VARS = {
    DATA = {VALUE="my data", NEEDED="yes", DESCRIPT="Any"},
}

function Init()
    Meta(METADATA)
    LoadVars(VARS)
end

function Main()
    PrintSuccs(VARS.DATA.VALUE)
end
