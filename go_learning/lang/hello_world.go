package main

import "fmt"

func main() {
    // 我是注释
    /*
    我也是注释
    */

    var b bool = false
    if (&b == nil) {
        fmt.Println(b)
    }

    
    fmt.Println("Hello, World!")
    /*
    当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，
    那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），
    这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见
    的，但是他们在整个包的内部是可见并且可用的
     */
}
