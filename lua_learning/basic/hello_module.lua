---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-22 10:31:55
---

hello_module = {}

-- 定义一个变量
hello_module.words = 'Hello World!'

-- 定义一个函数
function hello_module.func1()
    io.write("这是一个公有函数！\n")
end

local function func2()
    print("这是一个私有函数！")
end

function hello_module.func3()
    func2()
    hello_module.func1()
end

--local function hello_module.func3() end   --编译就不通过



return hello_module



