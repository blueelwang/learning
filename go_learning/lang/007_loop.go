package main

import "fmt"

func main() {
    
    /*
    循环
    for init; condition; post { }
    for condition { }   同C语言里的 while
    for {}              和 C 的 for(;;) 一样
    */

    for i := 0; i < 10; i++ {   // i只能在循环内部使用，后面再想用的话需要再定义
        fmt.Println(i)
    }

    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }

    i = 0
    for {
        fmt.Println(i)
        i++
        if i >= 10 {
            break
        }
    }

    /*
    For-each range 循环
    这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。
    */
    strings := []string{"google", "runoob"}
    for i, s := range strings {
        fmt.Println(i, s)
    }

    numbers := [6]int{1, 2, 3, 5}
    for i,x:= range numbers {
            fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
    }

}