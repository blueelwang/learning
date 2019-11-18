---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-28 22:07:24
---

local json = require('cjson')


local str  = [[ {"key" :"value"} ]]

local t    = json.decode(str)
ngx.say(" --> ", type(t))



local data = {1, 5}
data[5] = 99
ngx.say(json.encode(data))


ngx.say(json.encode({}))                       -- {}
ngx.say(json.encode({dogs = {}}))              -- {"dogs":{}}

json.encode_empty_table_as_object(false)
print(json.encode({}))                         -- []
print(json.encode({dogs = {}}))                -- {"dogs":[]}
