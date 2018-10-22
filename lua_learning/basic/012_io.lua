---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-22 14:30:34
---

file_name = 'data.txt'

---简单模式

-- 以只读方式打开文件
file = io.open(file_name, 'r')
-- 设置默认输入文件为 test.lua
io.input(file)
-- 输出文件第一行
print(io.read())
-- 关闭打开的文件
io.close(file)

-- 以附加的方式打开只写文件
file = io.open(file_name, 'a')
-- 设置默认输出文件为 test.lua
io.output(file)
-- 在文件最后一行添加 Lua 注释
io.write(".idea\n")
-- 关闭打开的文件
io.close(file)
--io.read() 方法还可以带参数，参数可以是下表中的一个：
--"*n"	        读取一个数字并返回它。                             例：file.read("*n")
--"*a"	        从当前位置读取整个文件。                            例：file.read("*a")
--"*l"（默认）	读取下一行，在文件尾 (EOF) 处返回 nil。              例：file.read("*l")
--number        返回一个指定字符个数的字符串，或在 EOF 时返回 nil。    例：file.read(5)


---完全模式:使用外部的文件句柄来实现。它以一种面对对象的形式，将所有的文件操作定义为文件句柄的方法

-- 通常我们需要在同一时间处理多个文件。我们需要使用 file:function_name 来代替 io.function_name 方法
-- 以只读方式打开文件
fileOjb = io.open(file_name, 'r')
-- 输出文件第一行
print(fileOjb:read())
-- 关闭打开的文件
fileOjb:close()
-- 以附加的方式打开只写文件
fileOjb = io.open(file_name, 'a')
-- 在文件最后一行添加 Lua 注释
fileOjb:write("--test\n")
-- 关闭打开的文件
fileOjb:close()

---file:seek(optional whence, optional offset):

-- 设置和获取当前文件位置,成功则返回最终的文件位置(按字节),失败则返回nil加错误信息。
-- 参数 whence 值可以是:
-- "set": 从文件头开始
-- "cur": 从当前位置开始[默认]
-- "end": 从文件尾开始
-- offset:默认为0

---file:flush():

--向文件写入缓冲中的所有数据


---io.lines(optional file name):

--打开指定的文件filename为读模式并返回一个迭代函数,每次调用将获得文件中的一行内容,
--当到文件尾时，将返回nil,并自动关闭文件。
for line in io.lines(file_name) do
    print(line)
end

-- 以只读方式打开文件
file = io.open(file_name, "r")
file:seek("end",-25)
print(file:read("*a"))
-- 关闭打开的文件
file:close()



