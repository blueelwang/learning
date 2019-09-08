
--- Lua操作符

--- 算术运算

-- 除了加、减、乘、除等常见的运算之外，Lua还支持取整除法、取模和指数运算。
print(-1 + 3)   -- 3
print(0.1 + 0.2 == 0.3) -- false 和C语言类似浮点运算不精确，相等判断会有问题
print(1 - 2)    -- -1
print(2 * 3)    -- 6
print(5 / 2)    -- 2.5
--print(5 // 2) -- 2    Lua5.3引入
print(5 % 3)    -- 2
print(-5 % 3)   -- 1
print(2 ^ 3)    -- 8

-- Lua中的除法的结果都是浮点数（即使两个操作数都是整数，且能整除），如果想要得到向下取整的结果，可以用math.floor()函数。
-- Lua5.3为这种向下取整的除法引入了一个新的运算符//，对除法得到的结果向负无穷取整。

--print(math.type(4 / 2))     -- float
print(math.floor(5 / 2))    -- 2
--print(math.type(math.floor(5 / 2))) -- integer 注意，Lua5.3之前所有的数值类型都浮点数，5.3才引入整数
--print(-9 // 2)              -- -5

-- 取模运算需要注意有负数的情况，和其他语言有可能得出的结果不同，Lua的计算方法如下：
-- a % b == a - ((a // b) * b)
-- 结果的符号和第二个操作数一致
print(5 % 3)    -- 2
print(-5 % 3)   -- 1
print(5 % -3)   -- -1

--[[
C语言中：
5 % 3   == 2
-5 % 3  == -2
5 % -3  == 2
]]

--- 关系运算

--[[
  Lua提供了以下的关系运算：
  小于 <
  大于 >
  等于 ==
  不行于 ~=
  小于或等于 <=
  大于或等于 >=
  这些关系运算的结果都是布尔型。
  ]]

-- 不同类型的变量不相等，但是整数、浮点数这种细分的子类型不区分，都视为数字类型。
-- Lua中的等于判断对变量所对应内存地址的判断，两个表即使内容相同也是不相等的。
print(2 > 3)        -- false
print(2 < 3)        -- true
print(2 == 3)       -- false
print(2 ~= 3)       -- true
print(2 == 2.0)     -- true
print(nil == false) -- false
print(nil == 0)     -- false
print(1 == true)    -- false
print(0 == false)   -- false 没有类似C语言的类型转换
print(1 == "1")     -- false 不同的类型不相等
print('1' == "1")   -- true
print('a' == 'a')   -- true
print({} == {})     -- false

--- 逻辑运算

-- and 逻辑与操作符。 若 A 为 false，则返回 A，否则返回 B。               (A and B) 为 false。
-- or  逻辑或操作符。 若 A 为 true，则返回 A，否则返回 B。                (A or B) 为 true。
-- not 逻辑非操作符。与逻辑运算结果相反，如果条件为 true，逻辑非为 false。    not(A and B) 为 true。
-- 和其他语言类似，判断有短路效果，即：
-- and操作第一个为false，就不再执行后面的判断
-- or操作第一个为true，也不再执行后台的判断
local a, b = 1, 2;
function change_value()
    b = 3
    return true
end
if (a == 2 and change_value()) then end
print(b)    -- 输出2  a==2为false,不再执行change_value()
if (a == 1 or change_value()) then end
print(b)    -- 输出2  a==1为true,不再执行change_value()
if (a == 2 or change_value()) then end
print(b)    -- 输出3  a==2为false,需要继续执行change_value()判断
if (not(a == 2)) then
    print('a ~= 2') -- 会执行到这里
else
    print('a == 2')
end
-- 复杂的判断中使用not时最好不要省略括号，否则运算符的优先级会导致和想象中不一致的结果（后面会有各运算符的优先级顺序）
if (not a == 2) then    -- not的优先级比==高，会先计算 not a 得出false，再计算 false == 2，最终结果是false
    print('true')
else
    print('false')  -- 会执行到这里
end

---5.3提供了位运算，5.2及以下会报错

-- &    按位与
-- |    按位或
-- ~    按异或，当用作一无运算符时是按位取反
-- >>   逻辑右移，用0填充，Lua没有提供算术右移（即用符号位填充）
-- <<   逻辑左移
--print(0xff & 0xabcd) --205

---其他运算符

-- ..  连接两个字符串
-- #   一元运算符，返回字符串或表的长度。 #"Hello" 返回 5
print('www.' .. 'daemoncoder' .. '.com')    -- www.daemoncoder.com
print('wow ' .. 23333)                      -- wow 23333
print('wow ' .. -666)                       -- wow -666
print('wow ' .. 2.333)                      -- wow 2.333
--print('wow ' .. nil)                      -- 报错
--print('wow ' .. false)                    -- 报错
--print('wow ' .. true)                     -- 报错

print(#'www.daemoncoder.com')               -- 19
print(#'汉字')                              -- 6  多字节字符使用时需要注意，计算出的长度是字节长度。
print(#{'www.', 'daemoncoder', '.com'})     -- 3


---运算符优先级

--[[
    Lua中的运算符优先级从高到低排列如下：
    ^
    -(这里是一元运算符取反，不是减号)  #  not
    *  /  //  %
    +  -(这里是二元运算符减号，不是取反)
    ..
    <<  >>
    &
    ~
    |
    <  >  <=  >=  ~=  ==
    and
    or
 ]]

