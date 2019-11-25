package lang

import "fmt"

func Demo009() {

    // 数组定义格式 var variable_name [SIZE] variable_type
    var nums [8] int
    var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    // 初始化数组中 {} 中的元素个数可以小于 [] 中指定的数字，不足的用默认值补充，但是不能多于[]指定的数字
    var balance2 = [8]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    // 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
    var balance3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    var balance4 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    fmt.Println(nums, balance1, balance2, balance3, balance4)
    // var data1 []int = [3]int{1, 2, 3}   // cannot use [3]int literal (type [3]int) as type []int in assignment

    // 组数读写
    nums[0] = 4
    // nums[10] = 30  // invalid array index 10 (out of bounds for 8-element array)
    fmt.Println(nums)
    fmt.Println(len(nums))

    // 多维数组
    var matrix [2][3] int
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            matrix[i][j] = 100 + i * 10 + j
        }
    }
    fmt.Println(matrix)     // [[100 101 102] [110 111 112]]
    matrix = [2][3]int {
        {1000, 1001, 1002},
        {1010, 1011, 1012},
    }
    fmt.Println(matrix)     // [[1000 1001 1002] [1010 1011 1012]]


    // 数组做为函数参数传递，形参和参数定义语句必需严格一致
    // []int{1, 2, 4}  不能 传递给 [3]int{1, 2, 3}，虽然他们实际的大小相同
    func1([3]int{1, 2, 3})
    // func1([]int{1, 2, 3})        // cannot use []int literal (type []int) as type [3]int in argument to func1
    // func1([4]int{1, 2, 3, 4})    // cannot use [4]int literal (type [4]int) as type [3]int in argument to func1
    func2([]int{1, 2, 3})
    //func2([3]int{1, 2, 3})        // cannot use [3]int literal (type [3]int) as type []int in argument to func2

}

func func1(nums [3]int) {
    fmt.Println(len(nums))
    fmt.Println(nums)
}

func func2(nums []int) {
    fmt.Println(len(nums))
    fmt.Println(nums)
}
func func3(nums [4]int) {
    fmt.Println(len(nums))
    fmt.Println(nums)
}
func func4(nums []int, size int) {
    fmt.Println(len(nums))
    fmt.Println(nums)
}