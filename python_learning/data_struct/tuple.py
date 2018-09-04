# -*- coding: utf-8 -*-

# tuple是不可变的
t = ()
t = (1, 'Tom', 3.14, ('a', 'b'), [3, 5])
print(t)

t = ('Tom')         # t 是一个字符串,而不是一个元组
print(type(t))      # <class 'str'>
t = ('Tom',)
print(type(t))      # <class 'tuple'>
t = 1, 2, 3,
print(type(t))      # <class 'tuple'>
print(t)            # (1, 2, 3)
t = tuple(range(10))

# 元组访问
t = (0, 1, 2, 3, 4, 5, 6)
print(t[-3])        # 同list,不能越界
print(t[-2:])
print(t[0:5:2])
t1 = ('a', 'b', 'c')
t2 = ('A', 'B', 'C')
print(t1 + t2)      # ('a', 'b', 'c', 'A', 'B', 'C')

print(2 in t)       # True
print(len(t))
print(min(t))
print(t.count(2))
print(t.index(2))
a, b = 5, 10        # 个数需要一致  5, 10, 15 和 5, 都会报错
print(a)
print(b)

t = (3, 5, 2, 4)
print(sorted(t))    # [2, 3, 4, 5] 转成了数组
# t.sorted()        元组不可变，所以没有 .sort()方法

# 元组的应用：
# 1.在映射类型中当做键来使用
# 2.函数的特殊类型参数
# 3.函数的特殊返回类型


def foo(args1, *argst):     # 可变长的函数参数（元组）
    print(args1)            # Hello
    print(argst)            # ('Zhang3', 'Li4', 'Wang5', 'Zhao6')

foo('Hello', 'Zhang3', 'Li4', 'Wang5', 'Zhao6')


def foo():
    return 1, 2, 3
print(foo())                # (1, 2, 3)





