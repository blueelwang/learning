# -*- coding: utf-8 -*-

# 可变集合set，不可变集合frozenset

names = ['zhang3', 'li4', 'wang5', 'li4', 'Zhang3']
nameSet = set(names)
nameFrozenSet = frozenset(names)
print(nameSet)          # {'zhang3', 'Zhang3', 'li4', 'wang5'}
print(nameFrozenSet)    # {'zhang3', 'Zhang3', 'li4', 'wang5'}

# 集合比较
print('zhang3' in names)            # True      属于
print('zhao4' not in names)         # True      不属于
print(nameSet == nameFrozenSet)     # True      值比较
print(nameSet != nameFrozenSet)     # False     值比较

# 集合关系 属于 包含关系
aeiou = set('aeiou')
print(aeiou)                # {'i', 'e', 'a', 'o', 'u'}
print('u' in aeiou)         # True
ao = set('ao')
print(ao < aeiou)               # True              ao是aeiou的真子集
print(aeiou > ao)               # True
print(set('aeiou') < aeiou)     # False
print(set('aeiou') <= aeiou)    # True

# 集合关系运算
aSet = set('sunrise')
bSet = set('sunset')
print(aSet & bSet)              # 交集 aSet ∩ bSet    {'u', 's', 'e', 'n'}
print(aSet | bSet)              # 并集 aSet ∪ bSet    {'s', 'i', 'r', 'u', 'e', 't', 'n'}
print(aSet - bSet)              # 差集                {'i', 'r'}
print(bSet - aSet)              # 差集                {'t'}
print(aSet ^ bSet)              # {'r', 'i', 't'}  并集 - 差集    对称差集：集合A与集合B的对称差集定义为集合A与集合B中所有不属于A∩B的元素的集合，记为A△B


aSet.remove('s')    # 's'不存在抛 KeyError
aSet.discard('s')   # 's'不存在，do nothing
print(set('asdf').issubset(set('asdf')))    # True
aSet.difference(bSet)
cSet = aSet.copy()

aSet.add('x')
aSet.clear()
aSet.update('qwer')     # 把 'q' 'w' 'e' 'r'四个成员分别加入，而不是一个'qwer'







