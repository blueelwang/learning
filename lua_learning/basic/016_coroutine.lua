---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by 32702.
--- DateTime: 2018-10-23 下午 04:23
---

co = coroutine.create(function(i) print(i) end)
coroutine.resume(co, 1)   -- 1
print(coroutine.status(co))  -- dead

print("-----------------------------")
co = coroutine.wrap(function(i) print(i) end)
co(1)

print("-----------------------------")
co2 = coroutine.create(
        function()
            for i=1,10 do
                print(i)
                if i == 3 then
                    print(coroutine.status(co2))  --running
                    print(coroutine.running()) --thread:XXXXXX
                end
                coroutine.yield()
            end
        end
)

coroutine.resume(co2) --1
coroutine.resume(co2) --2
coroutine.resume(co2) --3
print(coroutine.status(co2))   -- suspended
print(coroutine.running())


print("-----------------------------")
function foo (a)
    print("foo 函数输出", a)
    return coroutine.yield(2 * a) -- 返回  2*a 的值
end
co = coroutine.create(function (a , b)
    print("第一次协同程序执行输出", a, b) -- co-body 1 10
    local r = foo(a + 1)

    print("第二次协同程序执行输出", r, a, b)
    -- a，b的值为第一次调用协同程序时传入，
    --- yield的返回给主程，主程下次的调用参数给r,s，而不是yield的返回直接给了r,s
    local r, s = coroutine.yield(a + b, a - b)

    print("第三次协同程序执行输出", r, s, a, b)
    return b, "结束协同程序"                   -- b的值为第二次调用协同程序时传入
end)
print("main", coroutine.resume(co, 1, 10)) -- true, 4
print("------------------")
print("main", coroutine.resume(co, "r")) -- true 11 -9
print("------------------")
print("main", coroutine.resume(co, "x", "y")) -- true 10 结束协同程序
print("------------------")
print("main", coroutine.resume(co, "x", "y")) -- false  cannot resume dead coroutine
--[[输出：
第一次协同程序执行输出	1	10
foo 函数输出	2
main	true	4
------------------
第二次协同程序执行输出	r	1	10
main	true	11	-9
------------------
第三次协同程序执行输出	x	y	1	10
main	true	10	结束协同程序
------------------
main	false	cannot resume dead coroutine
--]]
---resume和yield的配合强大之处在于，resume处于主程中，它将外部状态（数据）传入到协同程序内部；而yield则将内部的状态（数据）返回到主程中。