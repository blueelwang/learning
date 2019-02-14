---
--- @author jianwei18@staff.weibo.com
--- @version 2018-10-22 10:31:55
---

g_data = 'global value'

-- 对于不同的请求实例（require 方法得到的对象）是相同的，因为 module 会被缓存到全局环境中。
-- 所以在这个位置千万不要放单请求内个性信息，例如 ngx.ctx 等变量
local hello_module = {}

-- 定义一个变量
hello_module.words = 'Hello World!'

-- 定义一个函数
function hello_module.func1()
    io.write("这是一个公有函数！\n")
end

local function func2()
    print("这是一个私有函数！")
end

-- func2想成为私有函数，没门
hello_module.func2 = func2;

function hello_module.func3()
    func2()
    hello_module.func1()
end

--local hello_module.func4() end   --编译就不通过



return hello_module



