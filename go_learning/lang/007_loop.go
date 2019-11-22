package main

import "fmt"

func demo007() {
    
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


    /*
    Go 语言中 break 语句用于以下两方面：
    用于循环语句中跳出循环，并开始执行循环之后的语句。
    break 在 switch 中在执行一条 case 后跳出语句的作用。
    在多重循环中，可以用标号 label 标出想 break 的循环。

    continue 类似，后面跟一个 label，表示要跳过哪层循环一步
    */

    for1:
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
        for2:
        for j := 11;; j++ {
            if j % i == 0 {
                continue for1   // 跳过最外层循环一步
            } else if j == 12 {
                continue        // 跳过最内层循环一步
            } else if j == 14 {
                break for2      // 结束2层循环
            } else if j == 20 {
                goto end
            }
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
    end:

    /*
    从上面的代码可以看出，Go语言支持 goto 语句，后跟一个标签表示要跳转到个里
    同其他语言，goto 也不能从一个块的外部跳转到块内，否则编译不通过
    goto for2    // goto for2 jumps into block starting at ./007_loop.go:58:29
    */
}