package lang

import "fmt"

func Demo010() {

    value := 3
    var p = &value
    fmt.Println(p, *p)

    // 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
    p = nil
    
    // 指针数组，存储5个指针的数组
    var ptrArray [5]*int
    fmt.Println(ptrArray)


    a := []int{10,100,200}
    var i int
    var ptr [3]*int;
    for  i = 0; i < 3; i++ {
        ptr[i] = &a[i]
    }
    fmt.Println(ptr)


    // 指向函数的指针
    v := func(c int, d int) int {
        return c + d
    }
    var ptr2 *func(int, int) int = &v
    fmt.Println((*ptr2)(1, 3))

}