package lang

import (
    "fmt"
)

func Demo011() {

    /*
    Go 语言的数组的定长性和值拷贝限制了其使用场景， Go 提供了另一种数据类型 slice (切片)
    这是一种变长数组，其数据结构中有指向数组的指针，所以是一种引用类型 
    slice 的定义，src/runtime/slice.go (go.19.1)
    type slice struct {
        array   unsafe.Pointer
        len     int
        cap     int
    }
    */

    // 可以声明一个未指定大小的数组来定义切片
    var nums []int
    // 切片初始化
    nums = []int{1, 2, 3}
    // 通过数组创建切片，注意修改切片或原数组，都会相互影响
    array := [4]int{1, 2, 3, 4}
    nums = array[:]
    nums[1] = 11
    array[2] = 22
    fmt.Println(nums, array)    // [1 11 22 4] [1 11 22 4]
    // 也可以使用make()函数来创建切片:
    nums = make([]int, 5)
    // 也可以指定容量
    nums = make([]int, 5, 20)
    //nums[6] = 10 // runtime error: index out of range
    


    //切片的赋值是引用
    nums2 := nums[:]    // 这里nums2是nums的一个引用，修改nums会影响num2的值
    fmt.Println(nums2, len(nums2))  // [0 0 0 0 0] 5
    nums[1] = 111
    fmt.Println(nums, nums2)    // [0 111 0 0 0] [0 111 0 0 0]
    nums2[2] = 222
    fmt.Println(nums, nums2)    // [0 111 222 0 0] [0 111 222 0 0]

    arr1 := [5]int{0, 1, 2, 3, 4}
    arr2 := arr1    // 数组的赋值不同于切片，这里arr2是arr1的拷贝，不会相互影响
    arr1[2] = 22
    fmt.Println(arr1, arr2) // [0 1 22 3 4] [0 1 2 3 4]
    arr2[3] = 333
    fmt.Println(arr1, arr2) // [0 1 22 3 4] [0 1 2 333 4]
    
    nums = []int{0, 1, 2, 3, 4}
    nums2 = nums[:]             // 这里也是引用
    nums2[1] = 11111            // [0 11111 2 3 4] [0 11111 2 3 4]
    fmt.Println(nums, nums2)    // [0 111 2 3 4] [0 111 2 3 4]
    nums2 = nums[1:3]           // 从下标1位置到下标3位置（不包含3）
    fmt.Println(nums, nums2)    // [0 11111 2 3 4] [11111 2]
    nums2[0] = 11
    fmt.Println(nums, nums2)    // [0 11 2 3 4] [11 2]
    nums2 = nums[1:]            // 从第2个元素到最后
    fmt.Println(nums, nums2)    // [0 11 2 3 4] [2 3 4]
    nums2 = nums[:2]            // 从下标0到下标2（不含2）
    fmt.Println(nums, nums2)    // [0 11 2 3 4] [0 11]
    // nums2 = nums[:-1]        // 编译报错，不支持负数语法
    

    //  len() 获取元素数    cap() 获取切片当前容量
    fmt.Println(len(nums2), cap(nums2))     // 2 5
    nums2 = []int{0, 1, 2, 3, 4, 5, 7, 9}
    fmt.Println(len(nums2), cap(nums2))     // 8 8
    nums2 = make([]int, 5, 20)
    fmt.Println(len(nums2), cap(nums2))     // 5 20

    // 切片在未初始化之前默认为 nil，长度为 0
    var nums3 []int
    fmt.Println(nums3, len(nums3), cap(nums3), nums3 == nil)    // [] 0 0 true
    nums4 := []int{}
    fmt.Println(nums4, len(nums4), cap(nums4), nums4 == nil)    // [] 0 0 false


    // append() 函数向切片中添加元素
    nums5 := []int{0, 1, 2}
    nums5 = append(nums5, 3)        // 添加一个元素：3
    fmt.Println(nums5)              // [0 1 2 3]
    nums5 = append(nums5, 4, 5, 6)  // 添加多个元素
    fmt.Println(nums5)              // [0 1 2 3 4 5 6]

    // copy(a, b) 函数把切片b拷贝到a, 但是不会改变a的大小
    // 如果len(a) < len(b)，则b多出的部分丢弃
    // 如果len(a) > len(b)，则只把b中的元素对应覆盖到a，a中多出的部分还保持原值
    nums4 = []int{1, 2, 3}
    nums5 = []int{11, 22}
    copy(nums4, nums5)
    fmt.Println(nums4, nums5)   // [11 22 3] [11 22]

    nums4 = []int{1, 2}
    nums5 = []int{11, 22, 33}
    copy(nums4, nums5)
    fmt.Println(nums4, nums5)   // [11 22] [11 22 33]

    // copy() 不同于切片的直接赋值，两个变量的修改不会相互影响
    nums5[0] = 0
    fmt.Println(nums4, nums5)   // [11 22] [0 22 33]


    
    
}