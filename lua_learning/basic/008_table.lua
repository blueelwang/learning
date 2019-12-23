
-- 定义空table
local empty_table = {}
-- 定义一个字符串数组
local array = {'daemon', 'coder'}
-- 数组元素也可以多种类型混合
local mix_array = {'Nice', 666}
-- 元素是数组的数组，可以用这种方式定义多维组数。
local matrix = {
    {11, 12, 13, 14},
    {21, 22, 23, 24},
    {31, 32, 33, 34, 35, 36},
}
-- 定义一个键值对形式的组数，键和值用等号隔开，等号左侧是键，右侧是值：
local map = {name = 'DaemonCoder', url = 'daemoncoder.com'}
-- 注意字符串形式的key不需要像字符串定义一样用引号，下面这样定义会有语法错误：
-- local map = {'name' = 'DaemonCoder', 'url' = 'daemoncoder.com'}
-- 如果不得不用引号包起来时（比如字符串中包含了等号等特殊符号），把键用方括号括起来，同时键值也要用字符串的定义格式加引号
local map = {["name"] = 'DaemonCoder', ["url"] = 'daemoncoder.com'}
-- 其实[]形式的键，不仅可以用字符串，还可以是任意类型
local table_key = {[{}] = {}} -- 键和值者是table类型
local int_key = {[0] = {}} -- 键是int类型
local function_key = {[function() end] = '0'} -- 键是函数类型
-- []形式的键还可以是一个表达式：
local expression_key = {
    -- 一个空table作为key
    [{}] = {},

    -- 用一个变量作为key
    [empty_table] = {},

    -- 调用一个函数，返回值作为key
    [table.concat(empty_table, '')] = {},
}
-- 注意上面这些写法都是把键用[]包起来的情况，没有[]的键都是字符串，不能用这样写法。


local mix = {'AAA', name='daemoncoder', 'BBB'}
-- 上面 'AAA'、'BBB' 使用默认的数字键，'AAA'的键为1，'BBB'的键为2。
for k, v in pairs(mix) do
    print(k, v)
end
--[[
输出：
1	AAA
2	BBB
name	daemoncoder
]]


local t = {"zhang3", age = 18}
print(t[1], t['age']) -- zhang3 18
t['age'] = 19
print(t.age)   -- 19
-- print(t.1)  -- 语法错误，1 不是合法的标识符，不能用点号访问，只能用 t[1]
print(t[1])
print(t.name)  -- nil  不存在的key为nil
t.name = "zhang3" -- 给table设置新的值

-- 可以用for循环配合 pairs() 或 ipairs() 来遍历table。
-- ipairs() 用于遍历键为连续数字的table，键从1开始，依次递增，直到某个键对应的值为 nil 就停止，即使还有其他键为数字的值没有被遍历到，也会停止。
-- 因此，键如果不是从1递增的连续数字，要谨慎使用ipairs()，看下面使用示例：
local t = {"111", "222", "333"}
for i, v in ipairs(t) do
    print(i, v)
end
--[[
输出：
1	111
2	222
3	333
]]
print('---------------------')
local t = {"111", "222", nil, "333"}
for i, v in ipairs(t) do
    print(i, v)
end
--[[
输出：
1	111
2	222
]]
print('---------------------')
-- 下面这个什么也不会输出，因为 t[1] 为 nil
local t = {a="AAA", b="BBB", c="CCC"}
for i, v in ipairs(t) do
    print(i, v)
end

print('---------------------')
-- pairs() 可以用于遍历任意table，遇到值为 nil 时，会跳过，而不是停止。仔细想想这是合理的，值为 nil 和 table 中不存在这个键是一样的，因为 Lua 未定义的变量的值都是 nil。
-- 需要注意的是，它访问的顺序并不是table中各字段的定义顺序来的，而是根据键的哈希值来的，所以看上去可能是『乱序』，看下面示例：
local t = {a="AAA", b="BBB", c="CCC"}
for k, v in pairs(t) do
    print(k, v)
end
--[[
输出：
b	BBB
a	AAA
c	CCC
]]
print('---------------------')
local t = {"111", "222", nil, "333"}
for i, v in pairs(t) do
    print(i, v)
end
--[[
下面是输出，注意"333"是可以被输出的，值为 nil 的会被跳过。
1	111
2	222
4	333
]]
print('---------------------')


-- table 的长度
print(#{"111", "222", "333"})       -- 3
print(#{"111", "222", nil, "333"})  -- 4
print(#{"111", "222", nil, nil, "333"})  -- 2
print(#{a="AAA", b="BBB", c="CCC"}) -- 0
print(#{"111", [4]="222", [7]="333"}) -- 4
print(#{"111", [5]="222", [7]="333"}) -- 1
print(#{"111", [3]="222", [4]="333"}) -- 4
print(#{[4]="111", [5]="222", [6]="333"}) -- 6

print('---------------------')
local t = {1, 2, nil, 4, 5}
table.insert(t, 4, "new")   -- {1, 2, nil, 4, "new", 5}
table.insert(t, 3, "new2")  -- {1, 2, "new2", nil, 4, "new", 5}
print(table.concat({"aaa", "bbb", "ccc", "ddd"}, '|')) -- aaa|bbb|ccc|ddd
print(table.concat({"aaa", "bbb", "ccc", "ddd"}, '|', 2)) -- aaa|bbb|ccc|ddd
print(table.concat({"aaa", "bbb", "ccc", "ddd"}, '|', 2, 3)) -- bbb|ccc|ddd
-- print(table.concat({"aaa", "bbb", "ccc", "ddd"}, '|', 2, 5))
-- table.concat() 指定的下标不能越界，不然报错 invalid value (nil) at index 5 in table for 'concat'

-- table.insert(list: table, pos: integer, value: any)
-- table.concat(list: table, sep: string, i: integer, j: integer)



local t1 = {}
local t2 = {}
t2[t1] = 'www.daemoncoder.com'
print(t2[t1])           -- www.daemoncoder.com
print(t2['t1'])         -- nil
print(t2.t1)            -- nil

matrix = {
    {'11', '12', '13'},
    {'21', '22', '23'},
    {'31', '32', '33'},
}
for i = 1, #matrix do
    for j = 1, #matrix[1] do
        print(string.format('%d,%s = %s', i, j, matrix[i][j]))
    end
end
--1,1 = 11
--1,2 = 12
--1,3 = 13
--2,1 = 21
--2,2 = 22
--2,3 = 23
--3,1 = 31
--3,2 = 32
--3,3 = 33
print({'One', 'Two', 'Three'} == {'One', 'Two', 'Three'})   --false table的比较和数值不同，是引用比较

-- #array 稀疏矩阵用#计算出的长度不是全部的键值对数

array = {[1] = 'One', [2] = 'Two', [3] = 'Three' }
print(#array, table.concat(array, ' '))                 -- 3 One Two Three
array = {[1] = 'One', [2] = 'Two', ['3'] = 'Three' }
print(#array, table.concat(array, ' '))                 -- 2 One Two
array = {'One', 'Two', 'Three'}
print(#array, table.concat(array, ' '))                 -- 3 One Two Three

table.insert(array, 'Four')
print(table.concat(array, ' '))                 -- One Two Three Four
table.insert(array, 1, 'Half')      --在索引位置之前插入
print(table.concat(array, ' '))                 -- Half One Two Three Four

--返回table数组部分位于pos位置的元素. 其后的元素会被前移. pos参数可选, 默认为table长度, 即从最后一个元素删起。
table.remove (array, 1)
print(table.concat(array, ' '))                 -- One Two Three Four

table.sort (array)
print(table.concat(array, ' '))                 -- Four One Three Two

-- for pairs 会过滤值为nil的key
for k, v in pairs({ 1, 2, nil, 3}) do
    print(k, v)
end
--[[
1	1
2	2
4	3
 ]]

for k, v in pairs({ 1, 2, 0, 3}) do
    print(k, v)
end
--[[
1	1
2	2
3   0
4	3
 ]]

