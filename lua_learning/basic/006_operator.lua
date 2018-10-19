

a = 21
b = 10
print("Line 1 - c 的值为 ", a + b)
print("Line 2 - c 的值为 ", a - b)
print("Line 3 - c 的值为 ", a * b)
print("Line 4 - c 的值为 ", a / b)
print("Line 5 - c 的值为 ", a % b)
print("Line 6 - c 的值为 ", a ^ 2)
print("Line 7 - c 的值为 ", -a)

if (0.1 + 0.2 == 0.3) then
    print('=====', 0.1 + 0.2)
else
    print('!!!!!', 0.1 + 0.2)   -- !!!!!	0.3
end


--      ==  等于，检测两个值是否相等，相等返回 true，否则返回 false           (A == B) 为 false。
--      ~=  不等于，检测两个值是否相等，相等返回 false，否则返回 true          (A ~= B) 为 true。
--      >   大于，如果左边的值大于右边的值，返回 true，否则返回 false          (A > B) 为 false。
--      <   小于，如果左边的值大于右边的值，返回 false，否则返回 true          (A < B) 为 true。
--      >=  大于等于，如果左边的值大于等于右边的值，返回 true，否则返回 false    (A >= B) 返回 false。
--      <=  小于等于， 如果左边的值小于等于右边的值，返回 true，否则返回 false   (A <= B) 返回 true。


--      and 逻辑与操作符。 若 A 为 false，则返回 A，否则返回 B。               (A and B) 为 false。
--      or  逻辑或操作符。 若 A 为 true，则返回 A，否则返回 B。                (A or B) 为 true。
--      not 逻辑非操作符。与逻辑运算结果相反，如果条件为 true，逻辑非为 false。    not(A and B) 为 true。


--其他运算符
--      ..  连接两个字符串	a..b ，其中 a 为 "Hello " ， b 为 "World", 输出结果为 "Hello World"。
--      #   一元运算符，返回字符串或表的长度。 #"Hello" 返回 5