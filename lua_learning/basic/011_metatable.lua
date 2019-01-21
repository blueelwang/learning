---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-22 11:24:15
---



--在 Lua table 中我们可以访问对应的key来得到value值，但是却无法对两个 table 进行操作。
--因此 Lua 提供了元表(Metatable)，允许我们改变table的行为，每个行为关联了对应的元方法。
--例如，使用元表我们可以定义Lua如何计算两个table的相加操作a+b。
--当Lua试图对两个表进行相加时，先检查两者之一是否有元表，之后检查是否有一个叫"__add"的字段，若找到，则调用对应的值。"__add"等即时字段，其对应的值（往往是一个函数或是table）就是"元方法"。
--有两个很重要的函数来处理元表：
--setmetatable(table,metatable): 对指定table设置元表(metatable)，如果元表(metatable)中存在__metatable键值，setmetatable会失败 。
--getmetatable(table): 返回对象的元表(metatable)。

mytable = {}                          -- 普通表
mymetatable = {}                      -- 元表
mt = setmetatable(mytable,mymetatable)     -- 把 mymetatable 设为 mytable 的元表



---__index 元方法

--当你通过键来访问 table 的时候，如果这个键没有值，那么Lua就会寻找该table的metatable（假定有metatable）中的__index 键。
--如果__index包含一个表格，Lua会在表格中查找相应的键，返回对应的值，如果对应的是一个函数，则调用函数，table和键会作为参数传递给函数
other = { foo = 3 }
t = setmetatable({}, { __index = other })
print(t.foo)
print('-----------------------')
mytable = setmetatable({key1 = "value1"}, {
    __index = function(mytable, key)
        if 'key1' == key then       --存在的key不会调用此方法
            print(key .. ' exist')
            return 'O.O'
        else
            print(key .. ' not exist')
            return '=.='            --不存在的key调用此方法，这里返回的值当作key对应的值
        end
    end
})

print(mytable.key1, mytable.key2)    --value1	=.=

print('-----------------------------')


---__newindex 元方法

--__newindex 元方法用来对表更新，__index则用来对表访问 。
--当你给表的一个缺少的索引赋值，解释器就会查找__newindex 元方法。
--如果设置了__newindex元方法，则不在原来的表上进行赋值操作，而是在元表上进行，如果元表设置对应字段的值是函数，则调用函数。
mymetatable = {}
mytable = setmetatable({key1 = "value1"}, { __newindex = mymetatable })
print(mytable.key1)                         --value1
mytable.newkey = "新值2"
print(mytable.newkey,mymetatable.newkey)    --nil	新值2
mytable.newkey = "新值3"
print(mytable.newkey,mymetatable.newkey)    --nil	新值3

mytable.key1 = "新值1"
print(mytable.key1,mymetatable.key1)        --新值1	nil
--以上实例中表设置了元方法 __newindex，在对新索引键（newkey）赋值时（mytable.newkey = "新值2"），会调用元方法，而不进行赋值。
--而如果对已存在的索引键（key1），则会进行赋值，而不调用元方法 __newindex。
print('---------------------------------')
mytable = setmetatable({key1 = "value1"}, {
    __newindex = function(mytable, key, value)
        rawset(mytable, key, "\""..value.."\"")

    end
})
mytable.key1 = "new value"
mytable.key2 = 4
print(mytable.key1,mytable.key2)    --new value	"4"


print('---------------- 为表添加操作符 --------------------')
-- 计算表中最大值，table.maxn在Lua5.2以上版本中已无法使用
-- 自定义计算表中最大键值函数 table_maxn，即计算表的元素个数
local function table_maxn(t)
    local mn = 0
    for k, v in pairs(t) do
        if mn < k then
            mn = k
        end
    end
    return mn
end
mytable = setmetatable({ 1, 2, 3 }, {
    __add = function(mytable, newtable)
        for i = 1, table_maxn(newtable) do
            table.insert(mytable, table_maxn(mytable)+1, newtable[i])
        end
        return mytable
    end
})

secondtable = {4,5,6}
mytable = mytable + secondtable
for k,v in ipairs(mytable) do
    print(k,v)
end
--[[
1	1
2	2
3	3
4	4
5	5
6	6
--]]

--- 操作符：
--- __add	对应的运算符 '+'.
--- __sub	对应的运算符 '-'.
--- __mul	对应的运算符 '*'.
--- __div	对应的运算符 '/'.
--- __mod	对应的运算符 '%'.
--- __unm	对应的运算符 '-'.
--- __concat	对应的运算符 '..'.
--- __eq	对应的运算符 '=='.
--- __lt	对应的运算符 '<'.
--- __le	对应的运算符 '<='.


print('------------------ __call 元方法 -------------------')
-- 定义元方法__call
mytable = setmetatable({10}, {
    __call = function(mytable, newtable)
        local sum = 0
        for i = 1, table_maxn(mytable) do
            sum = sum + mytable[i]
        end
        for i = 1, table_maxn(newtable) do
            sum = sum + newtable[i]
        end
        return sum
    end
})
newtable = {10,20,30}
print(mytable(newtable))    -- 输出：70


print('------------------ __tostring 元方法 -------------------')
mytable = setmetatable({ 10, 20, 30 }, {
    __tostring = function(mytable)
        local sum = 0
        for k, v in pairs(mytable) do
            sum = sum + v
        end
        return "表所有元素的和为 " .. sum
    end
})
print(mytable)      -- 输出：表所有元素的和为 60
