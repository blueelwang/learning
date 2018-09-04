
def get_sum(a, b):
    """I am docstring"""
    sum = a + b
    print(sum)
    return sum
get_sum(2, -6)    # 少传参数报错


def incr(num, step=1):    # 默认参数，step=1行号两边没有空格，同其他语言，默认参数在右，后面不能再有非默认参数
    print(num + step)
incr(3, -1)
incr(3)
incr(step=2, num=-5)        # 关键字参数 方式，传入参数有顺序和定义顺序可不同，但是参数名不能写错
# incr(step=6, 3)           # wrong 关键字参数在最右，后面不能再有非关键字参数
incr(3, step=6)             # ok


# 传递函数
def call(func, args):       # func不能传字符串
    func(args)
call(incr, -2)

# lambda 匿名函数
r = lambda num, step=-1: incr(num, step)    # 函数传入x，返回 incr(num, step)
r(8, 3)
r(8)



