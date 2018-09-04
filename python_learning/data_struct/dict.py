# -*- coding: utf-8 -*-

# 声明dict
d = {'name': '张三', "age": 3}
d = {'name': '张三', 'address': {'city': 'Beijing', 'country': 'China'}, 'age': 33}   # 可嵌套
d = dict(name='李四', age=4)
d = dict([('name', '张三'), ('age', 3)])
d = dict([['name', '张三'], ['age', 3]])
d = dict((('name', '张三'), ('age', 3)))
d = dict((['name', '张三'], ['age', 3]))

# 访问
print(d['name'])        # d['key'] 访问不存在的key会报错
print(d.get('name'))
print(d.get('key'))     # d.get('name') 访问不存在的key返回 None
print(d.get('key', 'key not exist'))    # 不存在时自定义返回值
print(d.keys())         # 所有的键       dict_keys(['name', 'age'])
print(d.values())       # 所有的值       dict_values([3, '张三'])
print(d.items())        # 所有的键值对    dict_items([('name', '张三'), ('age', '3')])
print('name' in d)
for key in d:
    print(key + ": " + str(d[key]))
for (key, value) in d.items():
    print(key + ": " + str(value))

# 修改
d2 = d                  # 引用赋值, 修改d2, d也会跟随改变
d3 = d.copy()
print('age' in d3)      # 成员判断
d3['name'] = 'd3 Name'  # name不存在也可以
d4 = {'price': 3.14, 'name': 'd4 Name'}
d.update(d4)            # 将d4的items merge到d, d4如果包含d中已有的key, 会覆盖
del d['price']          # 删除 price, 如果 price 不存在, 报错
d.pop('name', None)     # 删除并返回 name, 如果不存在,返回None,而不是抛异常
d.popitem()             # 删除并返回（key, value）

len(d)                  # 成员数量

d = {}.fromkeys(("Zhang3", 'Li4', 'Wang5'), 300)    # {'Wang5': 300, 'Li4': 300, 'Zhang3': 300}
d = {}.fromkeys(["Zhang3", 'Li4', 'Wang5'], 300)

names = ['Zhang3', 'Li4', 'Wang5']
scores = ['99', '98', '0']
print(dict(zip(names, scores)))         # {'Zhang3': '99', 'Wang5': '0', 'Li4': '98'}
scores = ['99', '98']
print(dict(zip(names, scores)))         # score不够时   {'Li4': '98', 'Zhang3': '99'}
names = ['Zhang3', 'Li4']
scores = ['99', '98', '0']
print(dict(zip(names, scores)))         # names不够时 {'Zhang3': '99', 'Li4': '98'}

a = {'name': "zhang3", 'age': 3}
b = a
b['age'] = 4        # a和b指向同一个对象，修改b，a也被修改
print(a['age'])     # 4
b = {'name': "wang5", 'age': 5}     # b指向一个新的对象，而不是修改原有的对象，所以此时a的值不变
print(a)            # {'age': 4, 'name': 'zhang3'}
print(b)            # {'name': 'wang5', 'age': 5}

a = {'name': "zhang3", 'age': 3}
b = a
b.clear()           # 想同时清空a和b，用.clear()方法
print(a)            # {}
print(b)            # {}


def foo(args1, *argst, **argsd):
    print(args1)                    # hello
    print(argst)                    # ('zhang3', 'Li4')
    print(argsd)                    # {'type': 4, 'age': 3}

foo('hello', 'zhang3', 'Li4', age=3, type=4)







