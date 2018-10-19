
--控制结构的条件表达式结果可以是任何值，Lua认为false和nil为假，true和非nil为真。
--[ 0 为 true ]
if (false) then
    print("false")
elseif (0) then
    print("0 为 true")
else
    print("0 为 false")
end
