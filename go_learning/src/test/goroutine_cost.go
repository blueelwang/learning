package test

import (
	"time"
	"fmt"
	"sync"
)


func GoroutineCostTest() {
	t1 := time.Now().UnixNano()
	fmt.Println()
	var wg sync.WaitGroup
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
		}()
	}
	wg.Wait()
	t2 := time.Now().UnixNano()
	fmt.Println(t1, t2, t2 - t1, (t2 - t1) / 1000000)
	// 1576224839690213000 1576224841977187000 2286974000 2286
	// 1ms 大概可以创建3000个线程
	
}