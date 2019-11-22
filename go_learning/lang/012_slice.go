package main

import "fmt"

func main() {

    // Go 语言切片是对数组的抽象。
    // 数组的长度不可改变，切片的长度是可变的，可以追加元素，在追加时可能使切片的容量增大。

    // 可以声明一个未指定大小的数组来定义切片
    var nums []int
    // 切片初始化
    nums = []int{1, 2, 3}
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
    // 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
    // nums2 := num[startIndex:endIndex] 
    nums2 = nums[1:3]
    nums2[0] = 11
    fmt.Println(nums, nums2)    // [0 11 2 3 4] [11 2]
    

}