
--- Lua是动态类型语言，变量不要类型定义,只需要为变量赋值。
--- Lua中有8个基本类型分别为：
--- nil、boolean、number、string、userdata、function、thread和table。

--- nil
--  这个最简单，只有值nil属于该类，表示一个无效值（在条件表达式中相当于false）。

--- boolean
--  包含两个值：false和true。Lua 把 false 和 nil 看作是"假"，其他的都为"真"

--- number
--  表示双精度类型的实浮点数
print(type(2))
print(type(2.2))
print(type(0.2))
print(type(2e+1))                   --  2e+1  ==  20
print(type(0.2e-1))
print(type(7.8263692594256e-06))
print(0.1 + 0.2 == 0.3)             --  false

--- string
--  字符串由一对双引号或单引号来表示，也可以用2个方括号 "[[]]" 来表示多行的字符串
text=[[
aa
bb
]]
print(text)
-- 使用 # 来计算字符串的长度，如：
len = "www.runoob.com"
print(#len)
print(#"www.runoob.com")    -- 都输出14

--- function
--  由 C 或 Lua 编写的函数

--- userdata
--  表示任意存储在变量中的C数据结构

--- thread
--  表示执行的独立线路，用于执行协同程序

--- table
--  Lua 中的表（table）其实是一个"关联数组"（associative arrays），数组的索引可以是数字或者是字符串。
--  在 Lua 里，table 的创建是通过"构造表达式"来完成，最简单构造表达式是{}，用来创建一个空表。
a = {}
--  Lua 中的表（table）其实是一个关联数组，数组的索引可以是数字或者是字符串。但是字符串和数字不会自动转换，可以同时存在
a["10"] = "十"
a[10] = 22
a[10] = a[10] + 11          -- a["10"] 和 a[10]同时存在
for k, v in pairs(a) do
    print(k .. " : " .. v)
end
--  不同于其他语言的数组把 0 作为数组的初始索引，在 Lua 里表的默认初始索引一般以 1 开始。
local tbl = {"apple", "pear", "orange", "grape"}
for key, val in pairs(tbl) do
    print(type(key), key, val)  -- number	1	apple ...
end
local tbl2 = {a="apple", p="pear", o="orange", g="grape"}




print(type(true))                   --> boolean
print(type(10.4*3))                 --> number
print(type("Hello world"))          --> string
print(type(print))                      --> function
print(type(type))                       --> function
print(type(nil))                    --> nil
print(type(type(X)))                    --> string
print(type(2))                      --> number
print(type(2.2))                    --> number
print(type(0.2))                    --> number
print(type(2e+1))                   --> number
print(type(0.2e-1))                 --> number
print(type(7.8263692594256e-06))    --> number

--- nil（空）
--- 对于全局变量和 table，nil 还有一个"删除"作用，给全局变量或者 table 表里的变量赋一个 nil 值，等同于把它们删掉
