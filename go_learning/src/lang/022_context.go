package lang

import (
	"context"
	"time"
	"sync"
	"fmt"
)

func DemoContext() {

	/*
	goroutine 里并没有父子的概念，怎么通知所有 goroutine 退出？
	Go 1.7 提供了一 个标准库 context 来解决这个问题。它提供两种功能: 退出通知和元数据传递。 
	context 库的设计目的就是跟踪 goroutine 调用 ， 在其内部维护一个调用树，并在这些调用树中传递通知和元数据。
	*/
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(ctx, &wg)
    }

    time.Sleep(time.Second)
    cancel()

    wg.Wait()

}

func worker(ctx context.Context, wg *sync.WaitGroup) error {
    defer wg.Done()

    for {
        select {
        default:
            fmt.Println("hello")
        case <-ctx.Done():
            return ctx.Err()
        }
    }
}
