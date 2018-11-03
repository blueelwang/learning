
matrix = {
    {'11', '12', '13'},
    {'21', '22', '23'},
    {'31', '32', '33'},
}
for i = 1, #matrix do
    for j = 1, #matrix[1] do
        print(string.format('%d,%s = %s', i, j, matrix[i][j]))
    end
end
--1,1 = 11
--1,2 = 12
--1,3 = 13
--2,1 = 21
--2,2 = 22
--2,3 = 23
--3,1 = 31
--3,2 = 32
--3,3 = 33
print({'One', 'Two', 'Three'} == {'One', 'Two', 'Three'})   --false table的比较和数值不同，是引用比较

-- #array 稀疏矩阵用#计算出的长度不是全部的键值对数

array = {[1] = 'One', [2] = 'Two', [3] = 'Three' }
print(#array, table.concat(array, ' '))                 -- 3 One Two Three
array = {[1] = 'One', [2] = 'Two', ['3'] = 'Three' }
print(#array, table.concat(array, ' '))                 -- 2 One Two
array = {'One', 'Two', 'Three'}
print(#array, table.concat(array, ' '))                 -- 3 One Two Three

table.insert(array, 'Four')
print(table.concat(array, ' '))                 -- One Two Three Four
table.insert(array, 1, 'Half')      --在索引位置之前插入
print(table.concat(array, ' '))                 -- Half One Two Three Four

--返回table数组部分位于pos位置的元素. 其后的元素会被前移. pos参数可选, 默认为table长度, 即从最后一个元素删起。
table.remove (array, 1)
print(table.concat(array, ' '))                 -- One Two Three Four

table.sort (array)
print(table.concat(array, ' '))                 -- Four One Three Two


