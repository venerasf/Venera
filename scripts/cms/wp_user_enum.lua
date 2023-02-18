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
    TAGS = {"enum","http","wordpress","scanner","cms"},
    INFO = [[Wordpress user enumeration with /?author= api.]]
}

-- Arguments/Variables needed to execute script
VARS = {
    URL = {VALUE="http://example.com", NEEDED="yes", DESCRIPT="URL"},
    VERBOSE = {VALUE="false", NEEDED="no", DESCRIPT="Verbose output."},
    USERAGENT = {VALUE="Venera-Framework", NEEDED="no", DESCRIPT="User-Agent"}
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    local payload = "/?author=1"

    local request = http.request("GET", VARS.URL.VALUE)
    PrintInfoln("Requesting "..VARS.URL.VALUE..payload)
    request:header_set("User-Agent", VARS.USERAGENT.VALUE)
    local result, err = client:do_request(request)
    if err then error(err) end

    PrintInfoln("Code "..result.code)
    print(dump(result.headers))

end

function dump(o)
    if type(o) == 'table' then
       local s = '{ '
       for k,v in pairs(o) do
          if type(k) ~= 'number' then k = '"'..k..'"' end
          s = s .. '['..k..'] = ' .. dump(v) .. ','
       end
       return s .. '} '
    else
       return tostring(o)
    end
 end
 