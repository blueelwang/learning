package lang

import "fmt"


// Go 语言最少有个 main() 函数。
func Demo008() {
    sayHello()
    a, b := 1, 2
    swap(a, b)
    fmt.Println(a, b)
    swap2(&a, &b)
    fmt.Println(a, b)
    swap3(&a, &b)
    fmt.Println(a, b)
    dynamicAdd()
    sum(1, 2, 3, 4, 5, 6)

    var c1 Circle
    c1.radius = 10
    fmt.Println("圆的面积 = ", c1.getArea())
    
    deferDemo()
}

/*
Go 函数的格式：
func func_name(param1 param1_type, ...) return_types {
    ...
    return ...
}
*/
func hi(name string) string {
    return "Hi " + name
}
// 参数、返回值没有的话，可以省略
func hello() {
    fmt.Println("Hello World!")
}

// 参数类型省略：几个连续的的参数如果类型相同，可以只对最后一个参数指定类型
func add(a, b int) int {
    return a + b
}
func addAll(a float32, b, c, d int, e, f float64) float64 {
    // b, c, d 都是 int 类型
    // e, f 都是 float64 类型
    return float64(a) + float64(b) + float64(c) + float64(d) + e + f
}

// 函数返回多个值
func say(msg string) (bool, string) {
    fmt.Println(msg)
    return true, msg
}

// 函数调用方法，参数个数必需和定义严格相等
func sayHello() {
    res, msg := say("Hello World")
    fmt.Println(res, msg)
}

// 函数参数的值传递和引用传递
// 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
// 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数
func swap(x, y int) {
    x, y = y, x
}
// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
// 引用传递指针参数传递到函数内，以下使用了引用传递：
func swap2(x, y *int) {
    // 这里只把x,y存储的地址做了交换，但是没有交换地址指向内存的值，调用处还是指向原来的地址，不影响原来的值
    x, y = y, x
}
func swap3(x, y *int) {
    // 影响方法调用处原来变量的值
    *x, *y = *y, *x
}

// 函数做为参数传递
func call(action func(int, int) int, a int, b int) int {
    return action(a, b)
}
func dynamicAdd() {
    sum := call(add, 1, 2)
    fmt.Printf("Calling dynamic function, function[%s] param1[%d] param2[%d] return[%d]\n", add, 1, 2, sum)
} 

// 可变参数列表，必需放在函数参数的最后一部分，在函数体里可以当作切片使用。
func sum(vals ... int) {
    total := 0
    for _, val := range vals {
        total += val
    }
    fmt.Println(total)
}

// 支持有名的返回值，参数名就相当于函数体内最外层的局部变量，命名返回值变量会被初始化为类型零值，最后的 return 可以不带参数名直接返回。
func sumWithReturnName(a, b int) (sum int) {
    sum = a + b
    return // 因为返回参数定义处指定了变量，这里的 return 相当于 return sum，当然也可以显示地写 return sum
}

/*
Go 语言函数没有类似Java的重载，会报错重复声明
也不支持类似PHP的参数默认值
*/

// 闭包
// Go语言支持匿名函数，可作为闭包。
func getSequence() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// defer 关键字
// Go 函数里提供了 defer关键字，可以注册多个延迟调用，这些调用以 先进后出 的顺序在函数返回前被执行（即最后注册的最先执行）。
// 有点类似于 Java 语言中异常处理中的 finaly 子句。 defer 常用于保证一些资源最终一定能够得到回收和释放。
func deferDemo() {
    i := 1
    defer func(i int) {
        fmt.Printf("我在第%d个 defer 注册的方法中\n", i)
    }(i)    // i 此时为1
    i++
    defer func(i int) {
        fmt.Printf("我在第%d个 defer 注册的方法中\n", i)
    }(i)    // i 此时为2
    // defer 函数的实参在注册时通过值拷贝传递进去。后续再对传入的参数变量做改动，也不再会改变原来注册时传入参数的值
    i++     // defer 方法在实际被调用时，传入的参数是在注册时i的值，这里修改不会产生影响
    fmt.Println("我在所有 defer 注册的方法之后")
    if true {
        return
    }
    
    // defer语句必须先注册后才能执行，如果 defer 位于 return 之后，则等于没有注册，不会执行。
    defer func(i int) {
        fmt.Printf("我在第%d个 defer 注册的方法中\n", i)
    }(i)    // i 此时为1
    /*
    输出：
    我在所有 defer 注册的方法之后
    我在第2个 defer 注册的方法中
    我在第1个 defer 注册的方法中
    */

    
    // 注意：
    // 主动 调用 as.Exit(int) 退出进程时，defer将不再被执行(即使已经提前注册〉。
}




/*
方法
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
所有给定类型的方法属于该类型的方法集。语法格式如下：
func (variable_name variable_type) function_name() return_type{
    ...
}
*/
/* Circle 结构体定义 */
type Circle struct {
    radius float64
}
//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
    //c.radius 即为 Circle 类型对象中的属性
    return 3.14 * c.radius * c.radius
}
func (c *Circle) setArea(radius float64) bool {
    c.radius = radius
    return true
}


