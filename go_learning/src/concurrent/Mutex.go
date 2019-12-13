package concurrent

import (
	"fmt"
	"sync"
)

func MutextDemo() {

	// 声明一个互斥量，Mutext的零值表示未加锁，只需要声明就可以使用（此时值是类型的默认零值）
	var mutex sync.Mutex
	// var mutex *sync.Mutex 	// 但是不能支声明一个指针，因为指针的零值是 nil，此时没有对应的实例。
	mutex.Lock()
	defer mutex.Unlock() // 未加锁的mutex上调用 Unlock() 会触发一个不可以被恢复的panic： fatal error: sync: unlock of unlocked mutex

	var m sync.Mutex
	var data []int = make([]int, 3)
	go conumer(&m, data) // 注意 mutex 传递时要用指针传递，不然值传递会新建一个 mutex
	producer(&m, data)

}

func producer(mutex *sync.Mutex, data []int) {
	for {
		mutex.Lock()
		if (len(data) < 3) {
			data = append(data, 1)
			mutex.Unlock()
		} else {
			fmt.Println("full")
			mutex.Unlock()
		}
	}
}

func conumer(mutex *sync.Mutex, data []int) {
	for {
		mutex.Lock()
		if (len(data) > 0) {
			data = data[1:]
			mutex.Unlock()
		} else {
			fmt.Println("empty")
			mutex.Unlock()
		}
	}
}