---
--- @author jianwei18 <jianwei18@staff.weibo.com>
--- @date 2018-11-01 21:28:55
--- @copyright Copyright (c) 2009-2018 Weibo.com All Rights Reserved.
---

t3 = {}
t4 = t3
t4.name = 'www.daemoncoder.com'
print(t3.name)                  -- www.daemoncoder.com t3和t4指向同一块内存，修改t4的值，t3也受影响
t4 = nil                        -- 删除t4不会影响t3，相当于t4不再指向原来的那块内存，但是t3还是指向原来的内存
print(t3.name)                  -- www.daemoncoder.com

print("1" + "2")        -- 3
print("1" + 2)          -- 3
print("1.1" + "1")      -- 2.1
print("2e3" + "1")      -- 2001
print("2e-1" + "1")     -- 1.2
print("1" + 'a')
--[[
... attempt to perform arithmetic on a string value
stack traceback:
    learning/lua_learning/tmp.lua:9: in main chunk
    [C]: at 0x01013ed1f0
]]

print("1" + "2")        -- 3
print("1" + 2)          -- 3
print("1.1" + '1')      -- 2.1
print("2e3" + '1')      -- 2001
print("2e-1" + '1')     -- 1.2


local t = {
    'a',
    name = 'www.daemoncoder.com',
    'b',
    age = 18,
    ['hello world'] = 'cool',
}
print(t[1])                 -- a
print(t['1'])               -- nil
print(t['name'])            -- www.daemoncoder.com
print(t[2])                 -- b
print(t.age)                -- 18
print(t['hello world'])     -- cool

print('-------------------------\n')


t1 = {}
t2 = {}
t2[t1] = 'www.daemoncoder.com'
print(t2[t1])
print(t2['t1'])
print(t2.t1)

print('-------------------------\n')

s1 = 'www.daemoncoder.com'
s2 = 'www.daemoncoder.com'
print(s1 == s2)             -- true
t1 = {}
t2 = {}
print(t1 == t2)             -- false
print({} == {})             -- false

t3 = {}
t4 = t3
t4.name = 'www.daemoncoder.com'
print(t3.name)                  -- www.daemoncoder.com t3和t4指向同一块内存，修改t4的值，t3也受影响
t4 = nil                        -- 删除t4不会影响t3，相当于t4不再指向原来的那块内存，但是t3还是指向原来的内存
print(t3.name)                  -- www.daemoncoder.com


