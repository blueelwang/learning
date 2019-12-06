
--optional_function_scope function function_name( argument1, argument2, argument3..., argumentn)
--    function_body
--    return result_params_comma_separated
--end
--optional_function_scope: 该参数是可选的制定函数是全局函数还是局部函数，默认全局，加local为局部函数。
--result_params_comma_separated: 函数返回值，Lua语言函数可以返回多个值，每个值以逗号隔开


-- 函数的定义
local function func_name(arg1, arg2, ...)
    -- statements
    return arg1, arg2, ...
end

local function sum(a, b)
    return a + b
end
print(sum(2, 3))    -- 5

local function get_msg() 
    return "Hello", true -- 返回多个值
end
local msg, err = get_msg() -- 接收返回的多个值
print(msg, err) -- Hello	true

local sub = function(a, b)
    return a - b
end
print(sub(8, 5)) -- 3

local math = {
    add = sum,
    sub = sub,
}
print(math.add(5, 3))   -- 8
print(math.sub(5, 3))   -- 2

-- Lua的函数，不要求传入的实参个数和形参个数一致。实参个数大于形参个数时，多传入的实参被忽略；实参个数小于形参个数时，不足的参数在函数体中的值为 nil 。
function say(a, b)
    print(a, b == nil)
end
say('www.daemoncoder.com')      -- www.daemoncoder.com	true
say('daemon', 'coder', '.com')  -- daemon	false


-- 默认参数，不支持！！！

-- Lua不支持函数重载，这里的写法是函数的重定义！
function overload(a, b)
    print("In first function", a, b)
end
function overload(a, b, c)
    print("In second function", a, b, c == nil)
end
overload(1, 2) -- In second function	1	2	true


--print(add(3,4,5,6,7))           -- add在后面，此时调用会出错，add还没有定义，是在nil上调用
function add(...)
    local s = 0
    local args = {...}
    for _, v in ipairs(args) do      -- {...} 表示一个由所有变长参数构成的数组
        s = s + v
    end
    return s
end
print(add(1, 2, 3, 4))           -- 10



function average(...)
    local result = 0
    local arg={...}    --> arg 为一个表，局部变量
    for i,v in ipairs(arg) do
        result = result + v
    end
    return #arg, result/#arg
end
print(average(10,5,3,4,5,6))    -- 6	5.5


--通常在遍历变长参数的时候只需要使用 {…}，然而变长参数可能会包含一些 nil，那么就可以用 select 函数来访问变长参数了：select('#', …) 或者 select(n, …)
--select('#', …) 返回可变参数的长度
--select(n, …) 用于访问 n 到 select('#',…) 的参数
--调用select时，必须传入一个固定实参selector(选择开关)和一系列变长参数。如果selector为数字n,那么select返回它的第n个可变实参，
--否则只能为字符串"#",这样select会返回变长参数的总数
function foo(...)
    for i = 1, select('#', {...}) do  -->获取参数总数
        local arg = select(i, {...}); -->读取参数
        print("arg", arg);
    end
end
foo(1, 2, 3, 4);


function jump()
    run()
    print("jump")
end
-- jump(); -- 会报错：attempt to call a nil value (global 'run')
function run()
    print("run")
end
jump(); -- OK

