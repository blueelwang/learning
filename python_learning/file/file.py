# -*- coding: utf-8 -*-

f = open('data.txt', 'r', encoding='UTF-8')      # .不是此文件的所处的路径,而是python执行时的当前路径
s = f.read(1)       # 设置了utf-8,read(1)会读出一个汉字, rb方式会读出一个字节
print(s)
l = f.readline()    # \n没有被去除
print([l])
l = f.readlines()   # 从流的当前位置开始
print(l)

print('\n--------------')
f.seek(0)
for line in f.readlines():
    print(line, end='')

print('\n--------------')
f.seek(0)
for line in f:
    print(line, end='')
f.close()

print('\n--------------')
f = open('./data.txt', 'w', encoding='UTF-8')
f.write('我是文件\n第二行')
f.flush()
f.close()

print('\n----------------------------------')
# with 语句会自动关闭文件
with open('data.txt', 'r') as fd:
    lines = fd.readlines()
    print(lines)
