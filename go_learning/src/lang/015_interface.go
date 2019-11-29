package lang

import (
    "fmt"
)

/*
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，方法只需要声明没有具体实现。
接口定义的内部方法声明没有 func 关键字和方法体。
任何其他类型只要实现了这些方法就是实现了这个接口。

单纯地声明一个接口变量没有任何意义，接口只有被初始化为具体的类型时才有意义，
没有初始化的接口变量，其默认值是 nil

接口的初始化：
1. 如果具体类型实例的方法集是某个接口的方法集的超集，则称该具体类型实现了接口，
可以将该具体类型的实例直接赋值给接口类型的变量
2. 已经初始化的接口类型变量a，直接赋值给另一种接口变量b，要求b的方法集是a的方法集的子集
var b B = a;  // a至少要包含类型B定义的所有方法

接口方法调用的最终地址是在运行期决定的


*/

/* 
最常使用的接口字面量类型就是空接口 interface{}。由于空接口的方法集为空，所以任意类
型都被认为实现了空接口。任意类型的实例都可以赋值或传递给空接口，包括非命名类型的实例。
*/
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}
// 接口中也可以嵌入另一个接口类型匿名字段，编译器会自动展开
// 下面这三种定义是等价的，最终都会转成第三种
type ReadWriter1 interface {
    Reader
    Writer
}
type ReadWriter2 interface {
    Reader
    Write([]byte) (int, error)
}
type ReadWriter3 interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}




/*
Phone 接口定义
*/
type Phone interface {
    call()
}

// Nokia 结构体
type Nokia struct {
}
// 定义 Nokia 结构体实现 call() 方法
func (Nokia Nokia) call() {
    fmt.Println("I am Nokia, I can call you!")
}

// IPhone 结构体
type IPhone struct {
}
// 定义 IPhone 结构体实现 call() 方法
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}
type Samsung struct {
    os string
}
// 定义 IPhone 结构体实现 call() 方法
func (iPhone Samsung) call() {
    fmt.Println("I am Samsung, I can call you!")
}

func DemoInterface() {
    defer exceptionHandler()
    
    var phone Phone

    // IPhone 实现了 Phone 接口，可以把 IPhone 实例赋值给声明为 Phone 类型的变量
    phone = new(IPhone)
    phone.call()

    phone = new(Nokia) 
    phone.call()


    // 有时需要确认已经初始化的接口变量指向实例的具体类型 是什么，也需要检查运行时的接口类型。

    // 类型断言 (Type Assertion)
    // 断言可以把接口A的实例a转成另一种接口B的实例b，语法为  var b B = a.(B)
    // 把接口实例转成具体实现类的实例，如果接口绑定的实例不是要转成的类型的，会抛出panic
    // 也可以把接口A的实例a转成另一种接口B的实例，前提是A的方法集是B的一个子集
    // 注意区分实例是否是指针类型
    
    // 此时phone是 *Nokia 类型的
    // var iphone IPhone = phone.(IPhone)  
    // 抛出panic：interface conversion: lang.Phone is *lang.Nokia, not lang.IPhone

    // var nokia Nokia = phone.(Nokia)  
    // 报错：interface conversion: lang.Phone is *lang.Nokia, not lang.Nokia

    var nokia *Nokia = phone.(*Nokia)
    fmt.Println(nokia)  // &{}

    // 返回的变量是接口绑定的实例值的副本，如果是后指针类型，那就是指针值的副本
    var phone2 Phone = Samsung{"Android 7"}
    var phone3 Samsung = phone2.(Samsung)   // phone3 是 phone2 的副本
    fmt.Println(phone3)     // {Android 7}
    phone3.os = "Android 8" // 修改新变量的值，原来的变量 phone2 不会被影响
    fmt.Println(phone2)     // {Android 7}
    fmt.Println(phone3)     // {Android 8}



    // 接口的类型查询
    // 判断工语法为 a.(type)
    // 注意，只能用在switch中，单独使用会报错：use of .(type) outside type switch
    // .(type)也只能用在接口类型的变量上，具体类型的变量上不能调用.(type)
    switch phone.(type) {
    case Samsung:
        fmt.Println("I am Samsung")
    case Nokia:
        fmt.Println("I am Nokia")
    // 指针类型和指向的类型是不同的类型
    case *Nokia:
        fmt.Println("I am *Nokia")      // 输出
    case IPhone:
        fmt.Println("I am Samsung")
    default:
        fmt.Println("Others")
    }
    // case后面可以跟 类型名，也可以是接口名，如果变量是对应类型，或实现了对应接口，则能匹配上
    // fallthrough 不能在 Type Switch 中使用

    // 也可以赋值给一个新的变量，v是switch内的一个局部变量
    switch v := phone.(type) {
    case Samsung, Nokia, *Nokia, IPhone:
        fmt.Println(v)
    default:
        fmt.Println("Others")
    }



    // 空接口和nil
    // 空接口不是真的为空，接口有类型和值两个概念
    var dog *Dog = nil
    var animal Animal = dog
    fmt.Printf("%p\n", dog)     // 0x0
    fmt.Printf("%p\n", animal)  // 0x0
    fmt.Println(dog == nil)     // true
    fmt.Println(animal == nil)  // false
    // 下面的语句会导致 panic
    // 方法转换为函数调用，第一个参数是Dog类型，由于dog是 nil，无法获取指针所指的对象值
    // animal.Talk()
    animal.Listen()    // 这个可以成功调用
    /*
    注意上面的 animal != nil，空接口有两个字段，一个是实例类型，另一个是指向绑定实例的指针，只有两个都为 nil 时 ，空接口才为 nil
    */
    var animal2 Animal              // 接口没有初始化，此时为nil
    fmt.Println(animal2 == nil)     // true


    // 接口的底层数据结构：
    /*
    src/runtime/runti me2 .go
    type iface struct {
        tab *itab           //itab 存放类型及方法指针信息
        data unsafe.Pointer //数据信息
    }
    */

}

type Animal interface {
    Listen()
    Talk()
}
type Dog struct {}
func (Dog) Talk() {
    fmt.Println("wang")
}
func (*Dog) Listen() {
    fmt.Println("xu...")
}

func exceptionHandler() {
    err := recover()
    if err != nil {
        fmt.Println(err)
    }
}




