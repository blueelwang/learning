package concurrent

import (
	"runtime"
	"fmt"
	"sync/atomic"
	"sync"
	"runtime/debug"
)


func PoolDemo() {

	/*
	sync.Pool 临时对象池类型：
	一个可以存放任务类型值的容器，自动伸缩，线程安全。
	New 成员是一个函数类型，当从临时对象池中取数据，如果池中没有数据可取时自动调用此函数来生成。
	Get() 从池中取一个值，并从池中删除，如果池为空且 New 成员也是空，则返回 nil
	Put() 放入池中一个值。

	临时对象池会为每一个 goroutine 关联的 P 建立一个本地私有池和一个本地共享池。本地共享池里值是在当前临时对象池的范围内共享的。
	当Get()时，先从自身关联 P 的本地私有池和本地共享池中取，没有则从其他 P 的本地共享池中偷一个返回，依然没有则调用 New 成员。

	临时对象池是对垃圾回收友好的，垃圾回收的执行一般会把临时对象池中的对象值全部移除。

	*/

	// 禁用GC，并保证在main函数执行结束前恢复GC
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func () interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}

	// New字段值的作用
	v1 := pool.Get()
	fmt.Printf("Value 1: %v\n", v1)

	// 临时对象池的存取
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	v2 := pool.Get()
	fmt.Printf("Value 2: %v\n", v2)

	// 垃圾回收时对临时对象池的影响
	debug.SetGCPercent(100)
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("Value 3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("value 4: %v\n", v4)


}