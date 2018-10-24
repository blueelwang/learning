---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-22 16:06:03
---

local lm = require "luasql.mysql"

--创建环境对象
env = lm.mysql()

--连接数据库
conn = env:connect("tmp_test","root","...","23.236.77.86", 3306)

--设置数据库的编码格式
conn:execute("SET NAMES UTF8")

--执行数据库操作
cur = conn:execute("select * from t1")

row = cur:fetch({},"a")

while row do
    print(string.format("%s", row.a))
    row = cur:fetch(row,"a")
end

conn:close()  --关闭数据库连接
env:close()   --关闭数据库环境
