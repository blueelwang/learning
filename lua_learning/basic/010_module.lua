---
--- @author jianwei18@staff.weibo.com
--- @version 2018-10-22 10:28:06
---

package.path='/Users/jianwei18/code/vg/learning/lua_learning/basic/?.lua;'

-- hello_module定义处有local，这里require的返回要赋给一个变量，不然取不到hello_module变量
local hello_module = require('hello_module')

hello_module.func3();
print(hello_module.words)
hello_module.words = 'new value'
print(hello_module.words)           -- new value

print('----------------------')

hm = require('hello_module')        -- 模块重命名
hm.func1()
print(hm.words)
--这里也输出修改之后的值 new value
--对于不同的请求实例（require 方法得到的对象）是相同的，因为 module 会被缓存到全局环境中。

print(g_data)       -- global value  全局变量


--[[
require 用于搜索 Lua 文件的路径是存放在全局变量 package.path 中，当 Lua 启动后，会以环境变量 LUA_PATH 的值来初始这个环境变量。
如果没有找到该环境变量，则使用一个编译时定义的默认路径来初始化。
文件路径以 ";" 号分隔，最后的 2 个 ";;" 表示新加的路径后面加上原来的默认路径，如
export LUA_PATH="~/lua/?.lua;;"
如果找过目标文件，则会调用 package.loadfile 来加载模块。否则，就会去找 C 程序库。
搜索的文件路径是从全局变量 package.cpath 获取，而这个变量则是通过环境变量 LUA_CPATH 来初始。
搜索的策略跟上面的一样，只不过现在换成搜索的是 so 或 dll 类型的文件。如果找得到，那么 require 就会通过 package.loadlib 来加载它。
--]]
