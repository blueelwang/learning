---
--- @author Sun Jianwei <327021593@qq.com>
--- @date 2019-09-17 11:54
---

local str = 'ULT=EScAAAAAAAAyMDNlZTY3OGYyMTE0NThmODU5ZDk5NzM1YzhjNTI5Yw%3D%3D'

--print(string.gsub(string.pack('>I4', 127), "(.)", function(x) return string.format("%02x", string.byte(x)) end))

local login_token = 'EScAAAAAAAAyMDNlZTY3OGYyMTE0NThmODU5ZDk5NzM1YzhjNTI5Yw=='
login_token = ngx.decode_base64(login_token)
local uid_str = string.sub(login_token, 1, 8)
local uid = 0
for i = 1, 8 do
    uid = uid + string.byte(uid_str, i) * (8 ^ (8 * (i - 1)))
end
print(uid)


