package concurrent

import (
	"fmt"
	"sync"
)

type Product struct {
	mutex *sync.Mutex
	producerCond *sync.Cond
	consumerCond *sync.Cond
	Num int
	MaxNum int
}

func CondDemo() {

	/*
	条件变量需要有一个绑定的锁，可以是互斥量，也可以是读写锁
	所以不能只声明，还需要定义，并传入绑定的锁实例。
	条件变量支持的操作：
	cond.Wait()
	cond.Signal()
	cond.Broadcast()


	
	Wait() 执行时有三步：
	1. 首先会对绑定的锁进行解锁操作
	2. 然后让当前线程阻塞。
	3. 阻塞被唤醒后，再次对绑定的锁进行加锁
	
	从第1步可以看出，Wait()执行时，绑定的锁必需是处于锁定状态，否则报错。
	如果没有第1步，其他线程可能会因为不能加锁而阻塞，有可能因此造成死锁。

	第1、2步（即解锁和阻塞）两步是原子性的。如果不是原子性，解锁之后，阻塞之前，
	有可能其他线程进入临界区，然后执行了唤醒动作，因为此时当前线程还没有执行阻塞，
	所以会错过这次唤醒，也就是当前线程在真正阻塞时，已经不满足需要阻塞的条件了。
	前2步的原子性保证了阻塞时条件变量的状态没有改变，也不会有唤醒被错过。
	
	当阻塞被唤醒后，还会对绑定的锁再次执行加锁操作（因为阻塞前执行了解锁），加
	锁成功之后 Wait() 才返回。所以 Wait() 被调用的前后都是处于锁定状态的。



	cond.Signal()
	唤醒一个在条件变量上阻塞的线程

	cond.Broadcast()
	唤醒全部在条件变量上阻塞的线程

	*/
	var mutex sync.Mutex
	var producerCond *sync.Cond = sync.NewCond(&mutex)
	var consumerCond *sync.Cond = sync.NewCond(&mutex)

	product := Product{&mutex, producerCond, consumerCond, 0, 10}

	go func() {
		for {
			produce(&product)
		}
	}()
	
	for {
		consume(&product)
	}
	

}

func produce(product *Product) {
	product.mutex.Lock()
	defer product.mutex.Unlock()

	if product.Num < product.MaxNum {
		product.Num++
		product.consumerCond.Signal()
	} else {
		fmt.Println("producer.Wait()")
		product.producerCond.Wait()
		fmt.Printf("producer is signaled! cur[%d]\n", product.Num)
	}
}

func consume(product *Product) {
	product.mutex.Lock()
	defer product.mutex.Unlock()

	if product.Num > 0 {
		product.Num--
		product.producerCond.Signal()
	} else {
		fmt.Println("consumer.Wait()")
		product.consumerCond.Wait()
		fmt.Printf("consumer is signaled! cur[%d]\n", product.Num)
	}
}