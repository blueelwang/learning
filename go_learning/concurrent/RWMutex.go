package concurrent

import (
	"sync"
)

func RWMutexDemo() {

	/*
	读写锁
	读写互斥，但是读读不互斥
	同互斥量，对未锁定的对象上执行解锁（无论是 Unlock 还是 RUnlock）都会报错
	*/
	var mutex sync.RWMutex
	mutex.Lock()	// 写锁定
	mutex.Unlock()	// 写解锁

	mutex.RLock()	// 读锁定
	mutex.RLock()
	mutex.RUnlock()	// 读解锁一次，此时锁还处于读锁定的状态，因为有两处加锁
	mutex.RUnlock()	// 再进行一次读解锁，此时锁才被解开
	

	// 如果上面读锁解锁次数不够的话，这里会死锁
	mutex.Lock()	// fatal error: all goroutines are asleep - deadlock!
	mutex.Unlock()

	// RLocker() 方法把调用者实例强转成 rlocker 类型，rlocker 类型是 sync.RWMutex 类
	// 型的一个别名，但是只定义了 Lock() 和 Unlock() 两个方法，这两个方法的内部实现分别是
	// 调用 sync.RWMutex 上的 RLock() 和 RUnlock()，所以 rlocker 类型上的操作都是读锁。
	rlock := mutex.RLocker()
	rlock.Lock()	// 此时等效于 mutext.RLock()
	// mutex.Lock()	// 这里会死锁，因为上面已经通过 rlock.Lock() 加了读锁
	// 如果换成调用 mutex.RLock() 不会死锁，因为两个加的都是读锁
	


}