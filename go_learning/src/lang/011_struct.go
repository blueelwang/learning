package lang

import "fmt"

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


func Demo011() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookID: 6495407})

    // 忽略的字段为 0 或 空
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

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

