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


a = 3
b = a
a = 4
print(b) --3
a = nil
print(b) --3
-- 当我们为 table a 并设置元素，然后将 a 赋值给 b，则 a 与 b 都指向同一个内存。如果 a 设置为 nil ，则 b 同样能访问 table 的元素。
-- 如果没有指定的变量指向a，Lua的垃圾回收机制会清理相对应的内存。
--
a = { key = 'value' }
b = a
for key, value in pairs(b) do
    print(key, value) --key value
end

a['key'] = 'value2'
for key, value in pairs(b) do
    print(key, value) --key value2
end

a = { key = 'value3' }
for key, value in pairs(b) do
    print(key, value) --key value2
end

local function change_value(data)
    data.key = 'value4'
end

change_value(a)
for key, value in pairs(b) do
    print(key, value) --key value2
end

a = nil
for key, value in pairs(b) do
    print(key, value) --key value2
end


