# -*- coding: utf-8 -*-

scores = [0, 1, 2, 3, 4, 2, 'a', [-1, 9]]   # 可以有不同的类型
scores[0]               # print scores[5] index out of boundary
scores[-1]              # print scores[-6] index out of boundary
scores[0:2]
scores[1:-1]
scores[0:5:2]

1 in scores             # True
'1' not in scores       # True
scores.index(2)         # 2在scores中第一次出现的下标
scores.count(2)         # 2出现的次数
list1 = ['a', 'b', 'c']
list2 = ['A', 'B', 'C']
print(list1 + list2)    # ['a', 'b', 'c', 'A', 'B', 'C']

# 常用函数
len(scores)
# min(scores)           # 元素全是数字
# max(scores)           # 元素全是数字
# sum(scores)           # 元素全是数字

# 改变可变序列
scores[2] = 2
scores[0:2] = ['a', 'b', 'c']   # OK
# scores[0:2] = 100               # error
scores[0:2] = [100]             # OK
scores[0:2] = [0, 1, 2, 3, 4]   # OK 长度可以不一致,但是赋值的需要是可迭代对象
scores[0:3:2] = [0, 2]          # 此时 长度需要一致
# 删除
del scores[0]                   # 根据下标删除
del scores[0:2]
del scores[0:3:2]
scores.remove(2)                # 根据值删除, 删除第一个值为2的元素
scores.clear()                  # python 2.7 没有
# 添加
scores.append(3)
scores.append([1, 2])           # [1, 2]作为一个元素
scores.extend([2, 3])           # 2和3分别做为一个元素加入  不能传入2这样的非数组
scores.insert(3, 'c')           # 在下标3的位置插入'c'
scores[3:3] = ['c']             # 功能同上
scores.pop(3)                   # 3表示下标 默认-1(最后一个元素)
scores.reverse()                # 就地反转
print(scores)

# 引用赋值
l = [1, 2, 3, 4]
s = l
s[0] = 'a'
print(l)                        # ['a', 2, 3, 4], s 和 l 指下同一份变量, 值改变相互影响
s = l[:]                        # s 和 l 指向两个变量
s = l.copy()                    # 同 s = l[:]

# 排序
scores = [1, 3, 5, 2, 4, 6]
scores.sort()                   # 就地排序, 升序
scores = ['a', 'b', 'c', 'A', 'B', 'C']
scores.sort()
scores = ['Abc', 'ABc', 'Ab']
scores.sort()
scores = ['Ca', 'Bb', 'Ac']
scores.sort(key=lambda x: x[-1])    # 根据第个元素的最后一个字母排序
scores = sorted(scores)             # 不是就地
scores = sorted(scores, key=lambda x: x[-1])
print(scores)

