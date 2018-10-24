---
--- @author jianwei18@staff.weibo.com
--- @version 2018-10-22 10:28:06
---

package.path='E:/code/learning/lua_learning/basic/?.lua;'
require('hello_module')


hello_module.func3();
print(hello_module.words)
hello_module.words = 'new value'
print(hello_module.words)



hm = require('hello_module')   --模块重命名
hm.func1()
print(hm.words)


--[[
require 用于搜索 Lua 文件的路径是存放在全局变量 package.path 中，当 Lua 启动后，会以环境变量 LUA_PATH 的值来初始这个环境变量。
如果没有找到该环境变量，则使用一个编译时定义的默认路径来初始化。
文件路径以 ";" 号分隔，最后的 2 个 ";;" 表示新加的路径后面加上原来的默认路径，如
export LUA_PATH="~/lua/?.lua;;"
如果找过目标文件，则会调用 package.loadfile 来加载模块。否则，就会去找 C 程序库。
搜索的文件路径是从全局变量 package.cpath 获取，而这个变量则是通过环境变量 LUA_CPATH 来初始。
搜索的策略跟上面的一样，只不过现在换成搜索的是 so 或 dll 类型的文件。如果找得到，那么 require 就会通过 package.loadlib 来加载它。
--]]

