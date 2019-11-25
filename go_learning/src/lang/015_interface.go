package lang

import (
    "fmt"
)

/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
*/

/*
Phone 接口定义
*/
type Phone interface {
    call()
}

// NokiaPhone 结构体
type NokiaPhone struct {
}
// 定义 NokiaPhone 结构体实现 call() 方法
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

// IPhone 结构体
type IPhone struct {
}
// 定义 IPhone 结构体实现 call() 方法
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

func Demo015() {
    var phone Phone

    // NokiaPhone 实现了 Phone 接口，可以把 NokiaPhone 实例赋值给声明为 Phone 类型的变量
    phone = new(NokiaPhone) 
    phone.call()

    phone = new(IPhone)
    phone.call()

}