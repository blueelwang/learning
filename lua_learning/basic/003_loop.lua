
--- while
a=10
while( a < 20 )
do
    print("a 的值为:", a)
    a = a+1
end


--- 数值for循环
for i=10,19,2 do    -- 从10变化到19，每次+2
    print(i)        -- 10 12 14 16 18
end
for i=10,1,-2 do    -- 从10变化到1，每次-2
    print(i)        -- 10 8 6 4 2
end
for i=10,12,1 do    -- 从10变化到12，每次+1
    print(i)        -- 10 11 12
end
--步长可省，默认为1
for i=10,12 do
    print(i)        -- 10 11 12
end
--for的三个表达式在循环开始前一次性求值，以后不再进行求值。
--比如下面的f(x)只会在循环开始前执行一次，其结果用在后面的循环中
function f(x)
    print("function")
    return x*2
end
a = 3
for a=1,f(a) do     -- 先根据全局变量a=3计算f(a), 再执行 for a=1,6 do ...
    print(a)        -- 输出从1到6
end
print(a)            -- 输出3，for中产生的变量a是局部变量，不影响全局变量的值

---泛型for循环
a = {"one", "two", "three"}
for i, v in ipairs(a) do
    print(i, v)
end


--- repeat util
a = 10
repeat
    print("a的值为:", a)
    a = a + 1
until( a > 15 )


--- break 语句
--Lua 编程语言 break 语句插入在循环体中，用于退出当前循环或语句，并开始脚本执行紧接着的语句。
--- lua中没有continue