package lang

import (
    "time"
    "fmt"
)

func DemoChannel() {

    /*
    通道 channel
    通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
    操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
    */


    // ch := make(chan string)  // 声明一个通道
    // 注意：默认情况下，通道是不带缓冲区的。发送端发送数据时会阻塞，直到接收方从通道中接收了值。
    
    // 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
    // ch := make(chan string, 100)
    // 如果通道带缓冲，发送方则会阻塞，直到发送的值被拷贝到缓冲区内；
    // 如果缓冲区已满，会阻塞，直到某个接收方获取到一个值

    // ch <- v                  // 把 v 发送到通道 ch
    // v2 := <-ch               // 从 ch 接收数据，会阻塞直到收到数据，如果通道已经关闭，没有读到数据也会立即返回


    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go arraySum(s[:len(s)/2], c)
    go arraySum(s[len(s)/2:], c)

    time.Sleep(2000 * time.Millisecond)
    fmt.Println("begin recv")
    x, y := <-c, <-c // 从通道 c 中接收
    fmt.Println("finish recv")

    time.Sleep(2000 * time.Millisecond)
    fmt.Println(x, y, x + y)

    /*
    示例输出顺序:
    before send     // 第一个 goroutine 往通道里发数据，然后阻塞，等待接收
    before send     // 第二个 goroutine 往通道里发数据，然后阻塞，等待接收
    begin recv      // 开始接收
    finish recv     // 接收完成
    after send      // 注意这里，往通道里发数据的 goroutine 在通道被接收之后，发操作才返回
    after send
    17 -5 12
    */


    c = make(chan int, 2)
    go arraySum(s[:len(s)/2], c)
    go arraySum(s[len(s)/2:], c)

    time.Sleep(2000 * time.Millisecond)
    fmt.Println("begin recv")
    x, y = <-c, <-c // 从通道 c 中接收
    fmt.Println("finish recv")

    time.Sleep(2000 * time.Millisecond)
    fmt.Println(x, y, x + y)

    /*
    示例输出：
    before send     // 第一个 goroutine 往通道里发数据，然后阻塞，等待接收
    after send      // 第一个 goroutine 发数据操作已经返回
    before send     // 第二个 goroutine 往通道里发数据，然后阻塞，等待接收
    after send      // 第二个 goroutine 发数据操作已经返回
    begin recv      // 开始接收数据
    finish recv     // 接收完成
    17 -5 12
    */



    /*
    只读通道 与 只写通道

    定义通道时可以指定通道是只读还是可写的，写法上的区别在于 <- 和 chan 的相对顺序。
    */
    readOnly  := make(<-chan int, 10)   // 只读通道
    writeOnly := make(chan<- int, 10)   // 只写通道
    v := <-readOnly
    // readOnly <- v    // invalid operation: readOnly <- v (send to receive-only type <-chan int)
    writeOnly <- v
    // v := <-writeOnly // invalid operation: <-writeOnly (receive from send-only type chan<- int)

    // 只读、只写通道不能相互赋值，但是可以把一个普通的通道赋值给一个只读或只写通道，反过来不行。
    // 这种通常用于方法的赋值，一个形参为只读通道的函数，即使传一个可读可写的通道，在函数内也只能对通道进行读操作。
    readOnly = make(chan int, 10)
    // var ch chan int = readOnly  // cannot use readOnly (type <-chan int) as type chan int in assignment
    writeOnly = make(chan int, 10)
    // var ch chan int = writeOnly // cannot use writeOnly (type chan<- int) as type chan int in assignment





    /*
    遍历通道、关闭通道
    */

    c = make(chan int, 10)
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    go filterLargeValue(nums, 5, c);
    // range 函数遍历从通道接收到的数据，如果不关闭通道，在接收第11个数据的时候会一直阻塞下去，然后报错：
    // all goroutines are asleep - deadlock!
    for i := range c {  // 注意这里不需要 <- 操作符
        fmt.Println(i)
    }


    /*
    通道的 len()、cap()

    len() 返回通道没有被读取的元素数量
    cap() 返回通道的容量
    对于无缓冲的通道，len() 和 cap() 都返回0
    */


}

func arraySum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
            sum += v
    }
    fmt.Println("before send")
    c <- sum // 把 sum 发送到通道 c
    fmt.Println("after send")
}

func filterLargeValue(s []int, threshold int, c chan int) {
    for _, v := range s {
        if (v > threshold) {
            continue
        }
        c <- v
    }

    // 关闭通道，关闭之后，可以继续从通道中读取，但是不能继续写入，也不能重复关闭通道，否则导致panic
    close(c) 
}