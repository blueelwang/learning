# -*- coding: utf-8 -*-


# 不可变序列

# s = 'c:\abc\zxy\tag'  会报错 \需要转义
s = r'c:\abc\zxy\tag'
print(s)

s = '我是汉字'
print(len(s))       # python3:4     2.7:12
print(s[0])         # python3 '我'
print(s * 5)        # 5个s
'是' in s            # True


# 字符串没有隐式转换
# '33' + 5  会报错, 5是int型的,不能和字符串相加
# int 转为 字符串
print('33' + str(5))    # 335
# 字符串 转为 int类型
print(int('33') + 5)    # 38


# 常用操作
s = '01234567'
s = s.replace('123', '')    # 字符串是不可变的, replace不是就地替换
print(s)
s = 'abcabcabc'
s = s.replace('ab', 'AB', 2)    # 替换前两个
print(s)

s = 'abcdabab'
s = s.capitalize()      # 首字母转为大写
s = s.upper()           # 全部转为大写
s = s.lower()
s.find('ab', 1, 6)
s.isalpha()             # 是否是 字母
s.isnumeric()           # 是否是 数字
s.isalnum()             # 是否是 数字或字母
print(s.split('ab'))    # ['', 'cd', '', '']
s.split()               # 只包含s数组
l = ['a', 'b', 'c', 'd']
print(','.join(l))      # a,b,c,d

# 格式化
s = '姓名:{0}, 姓名:{0}, 姓名:{0}, 年龄:{1}, 工作:{2}'.format('张三', 24, 'RD')
s = '姓名:{}, 年龄:{}, 工作:{}'.format('张三', 24, 'RD')
s = '姓名:{}, 年龄:{}, 工作:{}, 部门:{department}'.format('张三', 24, 'RD', department='KS')
print(s)
s = '{0:5}={1:10}'.format('不够5个', 3.14)             # 0点5个字符宽度, 1点12个
s = '{0:>5}={1:<10}'.format('不够5个', 3.14)           # 默认字符串左对齐, 数字右对齐,  可以>指定右对齐, <指定左对齐
s = '{0:3}={1:10}'.format('超过3个', 3.14)             # 实际超过不会被截断
s = '{0:f},{0:.3f},{0:4.2f},{0:25.20f},{0}'.format(3.14159265)
# 3.141593,3.142,3.14,   3.14159265000000020862,3.14159265
# 四舍五入,默认保留6位,位数不足补齐,小数点前用空格,后用0,浮点有误差
print(s)
s = '{0:X},{0:o},{0:b}'.format(127)     # 7F,177,1111111    16 | 8 | 2 进制
s = 'aaa{{aaa}}aaa'.format()            # aaa{aaa}aaa   去掉{}的特殊含义
print(s)

world = 'Python'
s = 'Hello %s'%(world)
print(s)

s = '汉字'
print(len(s))               # 2
print(list(s))              # ['汉', '字']
u = s.encode('utf-8')
print(u)                    # b'\xe6\xb1\x89\xe5\xad\x97'
print(len(u))               # 6
print(list(u))              # [230, 177, 137, 229, 173, 151]
u.decode()

# \OOO 八进制OOO代表的字符
# \xXX 十六进制XX代表的字符
print('\101\t\x41\n')       # A	A
print('\141\t\x61\n')       # a	a



