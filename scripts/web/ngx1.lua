--[[ 
  Script for Venera Framework
  https://github.com/farinap5/Venera
]]--

local http = require("http")
local client = http.client()

-- Metadata
METADATA = {
    AUTHOR = {"farinap5 <null>"},
    VERSION = "1.0",
    TAGS = {"nginx","http","proxy","ssrf"},
    INFO = [[Nginx SSRF via Host header. Try to connect to internal.]]
}

-- Arguments/Variables needed to execute script
VARS = {
    URL = {VALUE="http://example.com", NEEDED="yes", DESCRIPT="URL"}
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    local request = http.request("GET", VARS.URL.VALUE)
    local r, err = client:do_request(request)
    if r.code ~= 403 then
      PrintSuccsln("May be vulnerable to SSRF in Host header.")
    else
      PrintErrln("Not vulnerable to SSRF in Host header.")
    end
end
