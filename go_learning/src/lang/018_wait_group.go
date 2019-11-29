package lang

import (
	"fmt"
	"net/http"
	"sync"
)

func Demo018() {
	
	/*
	没有缓冲的通道具有 goroutine 同步的功能，除此之外，sync包也提供了多个同步的机制，主要是通过 WaitGroup 实现的。
	WaitGroup 用来等待多个 goroutine 完成， main goroutine 调用 Add() 设置需要等待 goroutine 的数目，然后调
	用 Wait() 等待所有的 goroutine 完成。每一个 goroutine 结束时调用 Done()，当等待的任务都调用了 Done() 之后，
	Wait() 函数就会返回。
	*/

	var wg sync.WaitGroup
	var urls = []string{
		"https://www.daemoncoder.com/",
		"https://daemoncoder.com/",
		"https://api.daemoncoder.com/",
	}

	for _, url := range urls {
		// 每一个任务启动一个 goroutine，同时给wg加1
		wg.Add(1)
		
		go func(url string) {
			// 当前 goroutine 结束后给wg计数减1, wg.Done()等价于wg.Add(-1)
			// defer wg.Add(-1)
			defer wg.Done()

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
			}
			
		}(url)
	}
	// 等待所有请求结束
	wg.Wait()

}