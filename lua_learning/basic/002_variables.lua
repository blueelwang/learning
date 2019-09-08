--变量在使用前，必须在代码中进行声明，即创建该变量。
--编译程序执行代码之前编译器需要知道如何给语句变量开辟存储区，用于存储变量的值。
--Lua 变量有三种类型：全局变量、局部变量、表中的域。
--Lua 中的变量默认是全局变量，那怕是语句块或是函数里，除非用 local 显式声明为局部变量。
--全局变量不仅可以跨语句块，还可以跨文件访问
--局部变量的作用域为从声明位置开始到所在语句块结束。
--变量的默认值均为 nil。


a = 'www.daemoncoder.com'
local b = 1

print(a)


--- 赋值语句
-- Lua可以对多个变量同时赋值，变量列表和值列表的各个元素用逗号分开，赋值语句右边的值会依次赋给左边的变量。
a, b = 10, 2 * 6
--遇到赋值语句Lua会先计算右边所有的值然后再执行赋值操作，所以我们可以这样进行交换变量的值
a, b = b, a
--当变量个数和值的个数不一致时，Lua会一直以变量个数为基础采取以下策略：
--1.变量个数 > 值的个数             按变量个数补足nil
--2.变量个数 < 值的个数             多余的值会被忽略
a, b, c, d = 1, 2, 3
print(a, b, c, d) --1	2	3	nil
a, b, c = 4, 3, 2, 1
print(a, b, c, d) --4	3	2	nil
--多值赋值经常用来交换变量，或将函数调用返回给变量：
function f()
    return 'A', 'B'
end

a, b = f()
print(a, b) -- A	B


--应该尽可能的使用局部变量，有两个好处：
--1. 避免命名冲突。
--2. 访问局部变量的速度比全局变量更快。

--- 示例1
-- lua变量的赋值类似C语言的指针赋值
a = 3       -- a变量类似C语言的一个指针，指向一块内存区域，这块内存的内容是3
b = a       -- 把a指针的指向的地址，赋值给b，此时a,b指向同一块内存，都是3
print(a, b) -- 3   3
a = 4       -- 修改a指针的地址，指向一个新的内存，存储的值是4，此时b指针还是指向原来的内存，没有改变
print(a, b) -- 4   3
b = a
a = nil     -- 把a指针清空，只是清除了a指针保存的地址，没有清除指向地址所存储的内容，
print(a, b) -- nil 3

--- 示例2
-- 字符串变量的赋值，和数值原理相同
a = "A"
b = a
a = "a"
print(b) --A
a = nil
print(b) --A

--- 示例3
-- table赋值
a = { key = 'value1' }
b = a  -- 把a赋值给b,此时a,b指向同一块内存地址
for key, value in pairs(b) do
    print(key, value) --key value1
end
a['key'] = 'value2'  -- 修改a指向内存存储的值，因为此时a,b指向同一块内存，b也会被影响
for key, value in pairs(b) do
    print(key, value) --key value2
end
a = nil  -- 把a指针设置为nil，原来指向的内存没被改变，b指向还是指向原来的内存，不受影响
for key, value in pairs(b) do
    print(key, value) --key value2
end

--- 示例4
-- table赋值
a = { key = 'value1' }
b = a
a = { key = 'value2' }  -- 把a指针指向一块新的内存地睛，此时a,b指针指向不再相同，b的指向和对应的内存没变
for key, value in pairs(b) do
    print(key, value) --key value1
end

a = { key = 'value1' }
b = a
a = nil  -- 把a指针设置为nil，原来指向的内存没被改变，b指向还是指向原来的内存，不受影响
for key, value in pairs(b) do
    print(key, value) --key value1
end

--[[
上面的几种示例，都是进行了这几步操作
    1. 对a赋一个初始值
    2. 把a赋值给b
    3. 修改a
    4. 看b是否也被修改
然后步骤4的结果表现出不一致的情况，数字类型（示例1）和字符串类型（示例2）的b不会被修改，
而表类型却有的被修改（示例3），有的没被修改（示例4）。最根本的原因在于第3步修改a时，是否在修改了a的指向。
如果修改a是修改了a指向的地址，则b不被影响，如：
a = 4
a = 'a'
a = { key = 'value3' }
如果没有修改a的指向，而是修改a指向地址所存储的值，则b也会被影响，如：
a['key'] = 'value2'
]]

-- 函数传参（表）
a = { key = 'value1' }
local function change_table_value(b)
    -- 此时a,b指向同一块内存，下面的改动没有改变a,b指针的指向，但是改变了他们指向地址所存储的内容
    b.key = 'value2'
end

change_table_value(a)
for key, value in pairs(a) do
    print(key, value) --key value2
end

a = { key = 'value1' }
local function change_to_new_table(b)
    -- 此时a,b指向同一块内存，下面的改动把b指向一块新的内存，没有改变a指针的指向，更没有改变a指向内存存储的内容
    b = { key = 'value2' }
end
change_to_new_table(a)
for key, value in pairs(a) do
    print(key, value) --key value1
end


