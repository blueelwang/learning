
--- while
i = 0
while i < 5 do      -- while 后面的括号可省略
    print("i=" .. i)
    i = i + 1
end

print('---------- 1 ----------')
--- 数值for循环
for i=10,12,1 do    -- 从10变化到12（包含10和12），每次+1
    print(i)        -- 10 11 12
end
print('---------- 2 ----------')
for i=10,9,1 do     -- 不会执行循环体
    print(i)        -- 没有输出
end
print('---------- 2.1 ----------')
for i=10,10,1 do
    print(i)        -- 输出：10
end
print('---------- 3 ----------')
for i=10,15,2 do    -- 从10变化到15，每次+2
    print(i)        -- 10 12 14
end
print('---------- 4 ----------')
for i=10,6,-2 do    -- 从10变化到6，每次-2
    print(i)        -- 10 8 6
end
print('---------- 5 ----------')
for i=10,7,-2 do    -- 从10变化到7，每次-2
    print(i)        -- 10 8
end
print('---------- 6 ----------')
for i=10,10,-2 do
    print(i)        -- 10
end
print('---------- 7 ----------')
for i=10,11,-2 do
    print(i)        -- 没有输出
end
print('---------- 8 ----------')
for i=10,12 do      -- 步长可省，默认为1
    print(i)        -- 10 11 12
end
print('---------- 9 ----------')
--for的三个表达式在循环开始前一次性求值，以后不再进行求值。
--比如下面的f(n)只会在循环开始前执行一次，其结果用在后面的循环中
function f(n)
    print("change number")
    return n + 2
end
n = 3
for n=1,f(n) do     -- 先根据全局变量n=3计算f(n), 再执行 for n=1,5 do ...
    print(n)        -- 输出从1到5
end
print(n)            -- 输出3，for中产生的变量n是局部变量，不影响全局变量的值

print('---------- 10 ----------')
arr = {'a', 'b', 'c'}
for i, v in ipairs(arr) do
   print(i, v)
end
--[[
输出：
1	a
2	b
3	c
 ]]
print('---------- 11 ----------')
---泛型for循环
arr = {a="A", b="B", c="C"}
for i, v in pairs(arr) do
    print(i, v)
end
--[[
输出：
b	B
a	A
c	C
 ]]

print('---------- 12 ----------')
--- repeat util
i = 0
repeat
    print("i=" .. i)
    i = i + 1
until i >= 5      -- 括号可省略


--- break 语句
--Lua 编程语言 break 语句插入在循环体中，用于退出当前循环或语句，并开始脚本执行紧接着的语句。
--- lua中没有continue