---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-25 17:46:42
---

-- 在 OpenResty 的世界里，不推荐使用这里的标准时间函数，因为这些函数通常会引发不止一个昂贵的系统调用，
-- 同时无法为 LuaJIT JIT 编译，对性能造成较大影响。推荐使用 ngx_lua 模块提供的带缓存的时间接口，如
-- ngx.today, ngx.time, ngx.utctime, ngx.localtime, ngx.now, ngx.http_time，以及 ngx.cookie_time 等。

print(os.time())    --unix_timestamp

local t1 = os.time({
    year    = 2018,
    month   = 10,
    day     = 25,
})
print(t1)   -- 1540440000


--指定时间
--必须含有 year、month、day 字段，其他字缺省时段默认为中午（12:00:00）
local t2 = os.time({
    year    = 2018,
    month   = 10,
    day     = 25,
    hour    = 17,
    min     = 46,
    sec     = 42,
    isdst   = false, -- 是否夏令时
})
print(t2)   -- 1540460802

print(os.difftime (t2, t1)) -- 20802


print(os.date('*t', t2))                -- 返回一个table
print(os.date('%Y-%m-%d %H:%M:%S'))     --当前时间
print(os.date('%Y-%m-%d %H:%M:%S', t2)) --2018-10-25 17:46:42
--[[
%a	一星期中天数的简写（例如：Wed）
%A	一星期中天数的全称（例如：Wednesday）
%b	月份的简写（例如：Sep）
%B	月份的全称（例如：September）
%c	日期和时间（例如：07/30/15 16:57:24）
%d	一个月中的第几天[01 ~ 31]
%H	24小时制中的小时数[00 ~ 23]
%I	12小时制中的小时数[01 ~ 12]
%j	一年中的第几天[001 ~ 366]
%M	分钟数[00 ~ 59]
%m	月份数[01 ~ 12]
%p	“上午（am）”或“下午（pm）”
%S	秒数[00 ~ 59]
%w	一星期中的第几天[1 ~ 7 = 星期天 ~ 星期六]
%x	日期（例如：07/30/15）
%X	时间（例如：16:57:24）
%y	两位数的年份[00 ~ 99]
%Y	完整的年份（例如：2015）
%%	字符'%'
 ]]


