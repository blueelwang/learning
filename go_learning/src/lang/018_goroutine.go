package lang

import (
    "time"
    "fmt"
    "runtime"
)




func DemoGoroutine() {

    fmt.Println("strat")


    /*
    操作系统可以进行线程和进程的调度，本身具备并发处理能力，但进程切换代价还是过于 高昂，进程切换
    需要保存现场，耗费较多的时间。如果应用程序能在用户层再构筑一级调度，将并发的粒度进一步降低。

    Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine ，go 关键宇后面必须跟一个函
    数，不能是语句或其他东西，函数的返回值被忽略。go 关键字指定以一个不同的、新创建的 goroutine 
    来执行一个函数。同一个程序中的所有 goroutine 共享同一个地址空间。
    
    goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。 
    
    goroutine 语法格式：    go 函数名( 参数列表 )

    注意:
    go 的执行是非阻塞的，不会等待。
    go 后面的函数的返回值会被忽略。
    调度器不能保证多个 goroutine 的执行次序。
    Go 程序在执行时会单独为 main 函数创建一个 goroutine，遇到其他 go 关键字时再去创 建其他
    的 goroutine。main goroutine 执行结束，其他的 goroutine 即使没有执行完，被也会终止。
    Go 没有暴露 goroutine id 给用户，所以不能在一个 goroutine 里面显式地操作另一个 goroutine，
    不过 runtime 包提供了 一些函数访 问和设置 goroutine 的相关信息
    */
    go task("111111111")
    go task("222222222")
    go task("3333333333")
    go task("4444444444")
    go task("55555555555")
    go task("666666666")
    go task("77777777777")
    // task("main")


    // runtime 包的 func NumGoroutine() int 可以返回当前 goroutine 的数量
    fmt.Println("NumGoroutine=", runtime.NumGoroutine())

    // fun GOMAXPROCS(n int) int
    // 查询或设置可以井发执行的 goroutine 最大数量
    fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))       // 传入参数0表示查询当前值
    fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(100))     // 参数大于0表示设置指定值，并返回原来的值
    fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))

    // func Goexit() 用来结束当前 goroutine 的运行， Goexit在结束当前 goroutine 运行之前会
    // 调用当前 goroutine 已经注册的 defer。Goexit 并不会产生 panic，所以该 recover 调用都
    // 返回 nil
    // 不能在 main goroutine 里调用，否则会报错：
    // fatal error: no goroutines (main called runtime.Goexit) - deadlock!
    // runtime.Goexit()

    // func Gosched() 表示放弃当前调度执行机会，将当前 goroutine 放到队列中等待下次被调度。
    runtime.Gosched()



}

func task(str string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Print(str + "\t\t" + string(i) + "\n")
        runtime.Gosched()
    }
    if str != "main" {
        runtime.Goexit()
    }
    
}