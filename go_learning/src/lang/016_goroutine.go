package lang

import (
    "time"
    "fmt"
)




func Demo016() {

    fmt.Println("strat")


    /*
    Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
    Go 允许使用 go 语句开启一个新的运行期线程，即 goroutine，goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
    go 关键字指定以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
    goroutine 语法格式：
    go 函数名( 参数列表 )
    */
    go task("Hello")
    task("World")
}

func task(str string) {
    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(string(i) + ":\t" + str)
    }
}