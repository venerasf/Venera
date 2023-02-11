# This script will list all todo within the project. 
# To create a todo(uppercase) the correct pattern is:
# /*
#   <Any description here, it wont be listed>
#   TOD O: <Description here. accepts line breaks>
# */

#╰─$ python3 todo.py
#./src/wlua/LuaStruct.go
#TOD O: use "REQUIRED" instead of "NEEDED".
#        This change will affect all scripts done.

#./src/wlua/LuaVars.go
#TOD O: escape string terminator that could allow lua injection
#        in SetVarValue func
#        the `VARS.%s.VALUE="%s"` is exploitable.
#*/
#func SetVarValue(L *lua.LState, key string, value string)


import re
import os

regex = r"(?:\/\*)(?:\n|.)*(TODO:[\w\s\n()!@#$%\"'&*.,;:´`=\[\]\?\/\\]+)(?:\n|.)*(?:\*\/)"

plist = []
dirs = ["./src","./scripts"]

for i in dirs:
    for currentpath, folders, files in os.walk(i):
        for file in files:
            plist.append(os.path.join(currentpath, file))

for i in plist:
    matches = re.finditer(regex, open(i,'r').read(), re.MULTILINE)
    for matchNum, match in enumerate(matches, start=1):
        for groupNum in range(0, len(match.groups())):
            groupNum = groupNum + 1
            print(i)
            print ("{group}".format(group = match.group(groupNum)))