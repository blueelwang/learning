---
--- @author      jianwei18@staff.weibo.com
--- @version     2018-10-25 18:48:38
---

--[[
math.rad(x)	角度x转换成弧度
math.deg(x)	弧度x转换成角度
math.max(x, ...)	返回参数中值最大的那个数，参数必须是number型
math.min(x, ...)	返回参数中值最小的那个数，参数必须是number型
math.random (m, n)	不传入参数时，返回 一个在区间[0,1)内均匀分布的伪随机实数；只使用一个整数参数m时，返回一个在区间[1, m)内均匀分布的伪随机整数；使用两个整数参数时，返回一个在区间[m, n)内均匀分布的伪随机整数
math.randomseed (x)	为伪随机数生成器设置一个种子x，相同的种子将会生成相同的数字序列
math.abs(x)	返回x的绝对值
math.fmod(x, y)	返回 x对y取余数
math.pow(x, y)	返回x的y次方
math.sqrt(x)	返回x的算术平方根
math.exp(x)	返回自然数e的x次方
math.log(x)	返回x的自然对数
math.log10(x)	返回以10为底，x的对数
math.floor(x)	返回最大且不大于x的整数
math.ceil(x)	返回最小且不小于x的整数
math.pi	圆周率
math.sin(x)	求弧度x的正弦值
math.cos(x)	求弧度x的余弦值
math.tan(x)	求弧度x的正切值
math.asin(x)	求x的反正弦值
math.acos(x)	求x的反余弦值
math.atan(x)	求x的反正切值
--]]


--另外使用 math.random() 函数获得伪随机数时，如果不使用 math.randomseed()
--设置伪随机数生成种子或者设置相同的伪随机数生成种子，那么得得到的伪随机数序列是一样的。
math.randomseed (100) --把种子设置为100
print(math.random())        -->output  0.011139255958739
print(math.random(100))     -->output  4
math.randomseed (100)
print(math.random())        -->output  0.011139255958739
print(math.random(100))     -->output  4
math.randomseed (os.time())
print(math.random())        -->output  0.0069887386700034

