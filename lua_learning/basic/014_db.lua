---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-22 16:06:03
---

package.path='E:/code/learning/lua_learning/basic/?.lua;./?.lua;'
require('mysql_conf')
local db_driver = require "luasql.mysql"

--luasql的错误处理，像sql语句错误等错误，会返回 nil,errmsg
--没有错误 errmsg 为 nil

--创建环境对象
env = db_driver.mysql()

--连接数据库
conn, errmsg = env:connect("tmp_test", mysql_conf.user, mysql_conf.passwd, mysql_conf.ip, mysql_conf.port)
print(conn, errmsg)
--成功返回如：MySQL connection (0x2321338)    nil
--失败返回如：nil	LuaSQL: error connecting to database. MySQL: Unknown database 'tmp_testxxx'

--设置数据库的编码格式
result, errmsg = conn:execute("SET NAMES UTF8")
print(result, errmsg)   -- 0    nil

result, errmsg = conn:execute("SETxxx NAMES UTF8")
print(result, errmsg)
--nil	LuaSQL: error executing query. MySQL: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'SETxxx NAMES UTF8' at line 1

str = [[
"'_
%\/*]]
print(str)
print(conn:escape(str)) -- \"\'_\n%\\/*

--设置是否自动提交
--成功返回true，否则false
result, errmsg = conn:setautocommit(false)
print(result, errmsg)   -- true nil

--执行sql语句，查询返回cursor对象，更新返回影响行数
cur, errmsg = conn:execute("select * from t1 for update")
if (errmsg) then
    --回滚，成功返回true，失败返回false
    result, errmsg = conn:rollback()
    print(result, errmsg)
end

--cur:getcolnames()
--Returns: a list (table) of column names.
--cur:getcoltypes()
--Returns: a list (table) of column types.


--获取结果集的下一行
--cursor:fetch(table, modestring)
-- 如果没有参数，直接返回下行数据，如果有table参数，会把下行数据写入table并返回
-- 可选参数table,用来存下一行数据
-- modestring: 'a',关联数组; 'n':数字索引数组
row = cur:fetch({},"a")
while row do
    print(row.a, row.b)
    row = cur:fetch({},"a")
    --最后一行之行再调用fetch()，会关闭对应的cursor
end

--关闭cursor对象，成功返回true，失败或已经被关闭返回false
result = cur:close()
print(result, errmsg)   -- false nil

--事务提交，返回true或false
result, errmsg = conn:commit()
print(result, errmsg)   -- true nil

--关闭链接，关联的cursor对象都已关闭 且 链接没有关闭，才返回true，否则返回false
result, errmsg = conn:close()
print(result, errmsg)   -- true nil

--关闭env对象，与之关联的链接对象都已经关闭 且 env对象没有被关闭，才会返回true，否则返回false
result, errmsg = env:close()
print(result, errmsg)   -- true nil
