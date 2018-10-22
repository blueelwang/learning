---
--- @author jianwei18@staff.weibo.com
--- @version 2018-10-22 10:28:06
---


require('hello_module')

hello_module.func3();
print(hello_module.words)
hello_module.words = 'new value'
print(hello_module.words)



hm = require('hello_module')   --模块重命名
hm.func1()
print(hm.words)

