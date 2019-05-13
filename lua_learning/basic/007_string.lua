
--- string

string1 = "www.daemoncoder.com"
string2 = 'www.daemoncoder.com'
string3 = [[www.daemoncoder.com]]
string4 = [[
www.daemoncoder.com]]
string5 = [[
www.daemoncoder.com
]]
print(#string3)                     -- #是取长度符号，19
print(#string4)                     -- 19，www.daemoncoder.com前面的换行不包含在字符串中
print(string3 == string4)           -- true
print(#string5)                     -- 20，多出了最后一个\n，后面的换行是字符串的一部分
print(string3 == string5)           -- false
print(string3 .. "\n" == string5)   -- true

string6 = string.gsub(string1, 'daemoncoder', 'DaemonCoder')
print(string1)                      -- www.daemoncoder.com string1的值没有被改变
print(string6)                      -- www.DaemonCoder.com

string1 = "Lua"
print("\"字符串 1 是\"",string1)    -- "字符串 1 是"	Lua
string2 = 'runoob.com'
print("字符串 2 是",string2)        -- 字符串 2 是	runoob.com
string3 = [[Lua学习]]
print("字符串 3 是",string3)        -- 字符串 3 是	Lua学习

string1 = "Lua\101";
print(string.upper(string1))
print(string.lower(string1))

print(string.gsub('abcdefg abcdefg','bcd','*', 1))
-- 输出 a*efg abcdefg	1
print(string.gsub('abcdefg abcdefg','bcd','*'))
-- 输出 a*efg a*efg	2

print(string.find('abcdefg abcdefg', 'abc', 1))
-- 输出 1	3

print(string.reverse('abcd'))
-- 输出 dcba
print(string.reverse('我是汉字'))
-- 输出 ��剱毘摈�

print(string.format("%05d", 17))
-- 输出 00017
print(string.format("PI = [%-5.2f]",3.14159))
-- 输出 PI = [3.14 ]

print(string.len("abc"))            -- 3
print(string.len("我是汉字"))         -- 12

print(string.rep('Abc ', 3))    -- Abc Abc Abc

print('Daemon' .. 'Coder')          -- DaemonCoder

for word in string.gmatch("Hello World", "%a+") do
    print(word)
end
--Hello
--World

print(string.match("I have 2 questions for you.", "%d+ %a+"))




