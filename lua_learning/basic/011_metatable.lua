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
setmetatable(mytable,mymetatable)     -- 把 mymetatable 设为 mytable 的元表


---__index 元方法

--当你通过键来访问 table 的时候，如果这个键没有值，那么Lua就会寻找该table的metatable（假定有metatable）中的__index 键。
--如果__index包含一个表格，Lua会在表格中查找相应的键。
other = { foo = 3 }
t = setmetatable({}, { __index = other })
print(t.foo)

print('-----------------------')

mytable = setmetatable({key1 = "value1"}, {
    __index = function(mytable, key)
        if 'key1' == key then       --存在的key不会调用此方法
            print('key exist')
            return 'O.O'
        else
            print('key not exist')
            return '=.='            --不存在的key调用此方法，这里返回的值当作key对应的值
        end
    end
})

print(mytable.key1,mytable.key2)    --value1	=.=

print('-----------------------------')


---__newindex 元方法

--__newindex 元方法用来对表更新，__index则用来对表访问 。
--当你给表的一个缺少的索引赋值，解释器就会查找__newindex 元方法：如果存在则调用这个函数而不进行赋值操作。
mymetatable = {}
mytable = setmetatable({key1 = "value1"}, { __newindex = mymetatable })
print(mytable.key1)                         --value1
mytable.newkey = "新值2"
print(mytable.newkey,mymetatable.newkey)    --nil	新值2

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



