package main

import (
	"time"
	"fmt"
	"net/http"
    _ "net/http/pprof"
)

func main() {
	// http://localhost:10086/debug/pprof
	go func() {
        http.ListenAndServe(":10086", nil)
	}()

	for i, sum := 0, 0; true; i++ {
		sum += i
		fmt.Println(i, sum)
		time.Sleep(1 * time.Millisecond)
	}
}