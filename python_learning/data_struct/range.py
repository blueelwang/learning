# -*- coding: utf-8 -*-

# range 类型, 是不可变类型
r = range(5)
print(type(r))       # <class 'range'>
# r[0] = 9           # 会报错 TypeError: 'range' object does not support item assignment

# 访问 同序列的访问
r[0]
3 in r              # True
min(r)
max(r)
len(r)
r.index(3)
r.count(2)

for i in range(5):
    print(i)            # 0 1 2 3 4
for i in range(2, 5):
    print(i)            # 2 3 4
for i in range(2, 5, 2):
    print(i)            # 2 4
