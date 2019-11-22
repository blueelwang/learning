package main

import "fmt"

func main() {

    /*
    map 变量声明格式：
    var variable_name map[key_type]value_type

    使用 make 函数 
    variable_name := make(map[key_type]value_type)
    */
    var m map[string]int        // 不初始化默认是 nil 
    m = make(map[string]int)    // 使用 make 函数定义
    m = map[string]int{"Go": 1, "PHP": 2}   // 也可以用数据直接初始化map
    fmt.Println(m)                          // map[Go:1 PHP:2]

    m["Java"] = 3
    fmt.Println(m)                          // map[Go:1 PHP:2 Java:3]
    m["C++"]++                              // 可以对一个原本不存在的key执行++
    fmt.Println(m)                          // map[Go:1 PHP:2 Java:3 C++:1]
    fmt.Println(m["C"])                     // 输出一个不存在的值，这里输出int的默认值：0

    // delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key
    delete(m, "Java")
    fmt.Println(m)                          // map[Go:1 PHP:2 C++:1]

    // key 不能是 切片、函数、map等复杂类型
    fmt.Println(make(map[bool]int))
    fmt.Println(make(map[float32]int))
    fmt.Println(make(map[int]int))
    fmt.Println(make(map[[3]int]int))
    // fmt.Println(make(map[[]int]int))             // invalid map key type []int
    // fmt.Println(make(map[func(int, int)]int))    // invalid map key type func(int, int)
    // fmt.Println(make(map[map[string]int]int))

    // 值可以是 切片、函数、map等复杂类型
    fmt.Println(make(map [string] []int))
    fmt.Println(make(map [string] [][]int))
    fmt.Println(make(map [string] func(int, int) int))
    fmt.Println(make(map [string] map[string]int))

}