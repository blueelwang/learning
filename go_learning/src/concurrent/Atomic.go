package concurrent

import (
	"fmt"
	"sync/atomic"
)

func AtomicDemo() {

	// Add
	var i32 int32 = 32
	newvalue := atomic.AddInt32(&i32, -10)	// 把i32减10，并返回减之后的值
	fmt.Println(i32)		// 22
	fmt.Println(newvalue)	// 22

	var ui32 uint32 = 32
	var step int = -10
	atomic.AddUint32(&ui32, ^uint32(-step-1))	// AddUint32() 第二个参数是 uint32，不能是负数，通过这个方法来实例减法
	fmt.Println(ui32)



	// CompareAndSwap
	var changed bool = atomic.CompareAndSwapInt32(&i32, 22, 222)
	fmt.Println(changed, i32) // true 222
	changed = atomic.CompareAndSwapInt32(&i32, 1234, 1235)
	fmt.Println(changed, i32) // false 222



	// Load
	i32 = atomic.LoadInt32(&i32)



	// Store
	atomic.StoreInt32(&i32, 128)
	fmt.Println(i32)


	// Swap 设置新值，返回原有值
	newvalue = atomic.SwapInt32(&i32, 256)
	fmt.Println(newvalue, i32) // 128 256



	var atomicVal atomic.Value
	fmt.Println(atomicVal.Load())	// <nil>
	atomicVal.Store(student{"张三"})
	fmt.Println(atomicVal.Load())	// {张三}




}

type student struct{
	name string
}