package main

import "fmt"
import "unsafe"

func demo002() {

    /* 常量 */
    const V1 = "2"	// 毕竟立即初始化，const V1 string 这样会编译报错
    // V1 = "3"	修改常量编译不通过
    // 不同于普通变量，常量未使用，不会报错

    // 用作枚举
    const (
        Unknown = 0
        Female = 1
        Male = 2
    )
    
    // 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。
    // 常量表达式中，函数必须是内置函数，否则编译不过：
    const (
        a = "abc"
        b = len(V1)
        c = unsafe.Sizeof(a)
    )



    
    // iota
    // iota，特殊常量，可以认为是一个可以被编译器修改的常量。
    // iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)
    // const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
    // iota 可以被用作枚举值：
    const (
        v1 = iota   // 0
        v2 = iota   // 1
    )
    const v3 = iota // 0
    const v4 = iota // 0

    const (
        v5 = iota    // 0
        v6           // 1
        v7           // 2
        v8 = "ha"    // 独立值，iota += 1
        v9           // "ha"   iota += 1
        v10 = 100    // iota +=1
        v11          // 100  iota +=1
        v12 = iota   // 7,恢复计数
        v13          // 8
    )
    const (
        v14 = 1 << iota // 1
        v15 = 3 << iota // 6
        v16             // 12
        v17             // 24
    )
    // var v18 = iota   // 会报错 undefined: iota   不能用在变量的定义，只能用在被编译器处理的常量
    // 同样，变量也不能用下面的方式定义，编译不通过
    // var (
    //     v18 = 3
    //     v19
    // )

    fmt.Println(v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17)
}