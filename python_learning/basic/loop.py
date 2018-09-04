
i = 1
while i < 10:
    i += 1  # python 不能用 i++
    print(i, end=' ')

a = []
for i in a:     # a需要是一个可迭代的对象  String  List  Tuple  Dictionary  File
    print(i, end=' ')
for i in range(3, 11, 2):
    print(i, end=' ')

# for可用于列表解析
a = [i for i in range(0, 5)]
print(a)
a = [i+1 for i in range(0, 10) if i % 2 == 0]
print(a)

# for可用于生成器解析
a = (i+1 for i in range(10) if i % 2 == 0)  # 返回一个生成器，用于数据量较大的场景
print(a)    # <generator object <genexpr> at 0x109a78a40>


# break continue
for i in range(10):
    if i < 5:
        continue
    elif i > 8:
        break
    else:
        print(i)
        i += 1

# while else      循环由break跳出，不执行else，不是于break跳出则执行else
i = 0
while i < 10:
    if i == 3:
        break
    print(i, end=' ')
    i += 1
else:
    print('I am else')







