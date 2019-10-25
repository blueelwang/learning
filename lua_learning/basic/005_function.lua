
--optional_function_scope function function_name( argument1, argument2, argument3..., argumentn)
--    function_body
--    return result_params_comma_separated
--end

--optional_function_scope: 该参数是可选的制定函数是全局函数还是局部函数，默认全局，加local为局部函数。
--result_params_comma_separated: 函数返回值，Lua语言函数可以返回多个值，每个值以逗号隔开

local function sum(a, b)
    return a + b
end
print(sum(2, 3))    -- 5

--print(add(3,4,5,6,7))           -- add在后面，此时调用会出错，add还没有定义，是在nil上调用
function add(...)
    local s = 0
    for i, v in ipairs{...} do      -- {...} 表示一个由所有变长参数构成的数组
        s = s + v
    end
    return s
end
print(add(3,4,5,6,7))           -- 25



function average(...)
    result = 0
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
    for i = 1, select('#', ...) do  -->获取参数总数
        local arg = select(i, ...); -->读取参数
        print("arg", arg);
    end
end
foo(1, 2, 3, 4);


