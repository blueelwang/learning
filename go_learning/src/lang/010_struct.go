package lang

import "fmt"

/*
Go语言提供了和C语言类似的结构体，来自定义一种包含多个成员的类型
struct 的存储空间是连续的，字段按照声明时的顺序存放
*/

// Books 结构体定义
type Books struct {
    title string
    author string
    subject string
    bookID int
}
// BookChartItem 结构体定义
type BookChartItem struct {
    book Books
    num int
}
var chart [10]BookChartItem

// 也可以定义类型别名，格式为 type newtype oldtype
type Map map[string]string

// 为定义的类型添加方法，方法的命名空间的可见性和变量一样，大写开头的方法可以在包外被访问，否则只能 在包内可见
func (book Books) buy(num int) {
}
func (book *Books) totalPrice(num int) {
}

// 类型方法有一个限制，就是方法的定义必须和类型的定义在同一个包中，
// 不能再给intbool等预声明类型增加方法，它们是内置的预声明类型，作用域是全局的
// 只能给自定义的命名类型添加方法

// 使用 type 定义的自定义类型是一个新类型，新类型不能调用原有类型的方法，但是底层类型支持的运算可以被新类型继承
func (m Map) Print() {
    for _, key := range m {
        fmt.Println(key)
    }
}

func Demo010() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookID: 6495407})
    // 忽略的字段为 0 或 空
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

    // 使用 new() 内置函数创建，宇段默认初始化为其类型的零值， 返回值是指向结构的指针
    var bookptr *Books = new(Books)
    bookptr.title = "new出来的title"



    var book1 Books = Books{title: "Go 语言", author: "www.runoob.com"}
    book2 := book1              // 结构体的赋值是拷贝
    fmt.Println(book2)
    fmt.Println(book2.title)
    book1.bookID = 3
    fmt.Println(book1.bookID)  // 3
    fmt.Println(book2.bookID)  // 0


    change1(book1)
    fmt.Println(book1)
    change2(&book1)
    fmt.Println(book1)

}

// 值传递，这里的book是一个拷贝，对其修改不会改变原来的变量
func change1(book Books) {
    book.title = "CHANGED"
}

// 传入指针，会改变原来的变量
func change2(book *Books) {
    // 注意这里可以可以用点
    book.title = "CHANGED"
    // 等价于 (*book).title = "CHANGED"  但是Go没有 -> 操作符
}

