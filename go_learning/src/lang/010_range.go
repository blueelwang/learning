package lang

import (
    "fmt"
)


func DemoRange() {

    // Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
    // 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
    
    nums := []int{0, 1, 2}
    for index, num := range nums {
        fmt.Printf("index[%d] num[%d]\n", index, num)
    }
    /*
    输出：
    index[0] num[0]
    index[1] num[1]
    index[2] num[2]
    */
    

    // range也可以遍历 map，但是不保证每次选代元素的顺序。
    data := map[string]string{"a": "AAA", "b": "BBB"}
    for k, v := range data {
        fmt.Printf("%s -> %s\n", k, v)
    }
    /*
    输出：
    b -> BBB
    a -> AAA
    */
    
    // range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    // 注意这里下标的不是按字符个数算的下标，而按字节算的下标，包含多字节字符时有可能不连续
    str := "Hello 中文"
    for i, c := range str {
        fmt.Printf("index[%d] char[%c]\n", i, c)
    }
    /*
    输出：
    index[0] char[H]
    index[1] char[e]
    index[2] char[l]
    index[3] char[l]
    index[4] char[o]
    index[5] char[ ]
    index[6] char[中]
    index[9] char[文]
    */

    fmt.Println(len(str))   // 12 字节数
    fmt.Println(str[6], str[7], str[8], str[9])


}