package concurrent

import (
	"fmt"
	"sync"
)

func OnceDemo() {

	var once sync.Once
	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Printf("run only 1 time, i=%d\n", i)
		})
	}

}