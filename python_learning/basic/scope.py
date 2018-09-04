
global_str = 'I am global string'
str = 'test string'


def test_func():
    # global global_str
    local_str = 'I am local string'
    str = "TEST STRING"
    print(global_str + ' ' + local_str + ' ' + str)
    # UnboundLocalError: local variable 'global_str' referenced before assignment
    # 加上下面这名会报错，因为不能 前面已经引用，再赋值，函数开始初用global声明可解决
    # global_str = 'GLOBAL STRING'  # 修改全局生效
test_func()
print(global_str)

