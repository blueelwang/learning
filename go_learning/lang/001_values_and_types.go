package main

import "fmt"

var a = 3   // 全局变量

func demo001() { // { 不能在像C++那样单独的一行
    // 我是注释
    /*
    我也是注释
    */

    // 分号可以省略，一行如果有多条语句可以用分号隔开
    var b bool = false      // 定义变量  var name type
    if (&b == nil) {
        fmt.Println(b)      
    }

    // 注意：声明的变量未使用，会直接编译不通过！！
    var hello = "Hello"
    var world = "World"
    fmt.Println(hello + " " + world)
    /*
    当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
    那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），
    这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见
    的，但是他们在整个包的内部是可见并且可用的
     */



    
    // 数据类型
    /* 布尔型 */
    var truevalue bool = true
    var falsevalue bool = false
    fmt.Println(truevalue, falsevalue)

    /* 数字类型 */
    
    // 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
    // Go 也有基于架构的类型: int、uint 和 uintpt，
    // uint8、uint16、uint32、uint64
    // int8、 int16、 int32、 int64
    
    // 浮点数：float32、float64
    // complex64    32位实数和虚数
    // complex128   64位实数和虚数
    
    // 其他:
    // byte 类似unit8
    // rune 类似int32
    // uint 32或64位
    // int  与uint一样大小
    // uintptr 无符号整型，用于存放一个指针
    var intvalue int = 30
    var uintvalue uint = 20
    var int8value int8 = 127    // 数值大小超出类型所表示的范围，编译不通过，此处不能是128
    int8value *= 2              // 运行时超出不会报错，但是数据被截断处理，此处计算出的值为 -2
    var fvalue float32 = 3.14
    var fvalue2 float64 = 3.14
    var c128 complex128 = 180 + 23i
    fmt.Println(intvalue, uintvalue, int8value, fvalue, fvalue2, c128)

    // 支持八进制、十六进制表示
    fmt.Println(0x0a)   // 10
    fmt.Println(0101)   // 65

    /* 字符串 */
    // 用双引号定义，单引号只用于表示字符，同Java
    // Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
    var stringvalue string = "I am string"
    var charvalue = 'a'         // char不是类型名，不能写成: var charvalue char = 'a'
    fmt.Println(stringvalue)
    fmt.Println(charvalue)      // 输出： 97
    fmt.Println("a\tbc\x64\x65\0777")    // a    bcde
    // 跨行字符串，第一个换行不包含在字符串，最后一个包含
    longstring := `
    123480
    `
    fmt.Println(longstring)

    /* 
        派生类型:
        包括：
        (a) 指针类型（Pointer）
        (b) 数组类型
        (c) 结构化类型(struct)
        (d) Channel 类型
        (e) 函数类型
        (f) 切片类型
        (g) 接口类型（interface）
        (h) Map 类型
    */
    strings := []string{"google", "runoob"}
    numbers := [6]int{1, 2, 3, 5}
    fmt.Println(strings, numbers)


    // 变量声明
    // 用户var来声明一个变量，格式  var variable_name type
    var v1 int
    // 也可以不显示声明类型，根据初始化的值，自动推断
    var v2 = 3
    // 也可以省略 var 关键字，用户 := 声明并初始化一个变量，但是已经声明过的变量不能再用 := ，否则会报错
    // 但是这只能用于函数中，不能用于全局变量的的声明，如最上面的a
    v3 := "v3"
    // 只声明被初始化，则变量默认为零值，也就是系统默认设置的值，各类型的默认值为：
    // 数值类型（包括complex64/128）为 0
    // bool 为 false
    // string 为 ""
    // 以下几种类型为 nil：
    //  var a *int
    //  var a []int
    //  var a map[string] int
    //  var a chan int
    //  var a func(string) int
    //  var a error // error 是接口

    // 多变量声明: var vname1, vname2, vname3 type
    var v4, v5 int = 1, 2
    var v6, v7 = 6, "v7"    // 类型不同
    v8, v9 := 8, "v9"


    // 空白标识符 _ 也被用于抛弃值，_ 实际上是一个只写变量，你不能得到它的值。
    // 这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
    _, v10 := 3, 4
    // 交换两个变量的值
    v1, v2 = v2, v1     // 两个变量的类型必须是相同


    // GO是强类型语言
    // 声明的一个变量之后，就不能再把其类型的值赋值给它，这一点和PHP等弱类型语言不同
    // v2 = "abc"


    // 变量赋值
    // 所有像 int、float、bool 和 string 这些基本类型都属于值类型，赋值时把内存中的值进行了拷贝
    // 可能通过 &i 来获取变量 i 的内存地址

    // 更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
    // 一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置
    // 当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。
    // 如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。

    fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)





    



}
