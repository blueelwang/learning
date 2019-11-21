package main

import "fmt"

func demo005() {

    /*
    switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
    switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
    switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
    switch val1 {
        case val1:
            ...
        case val2:
            ...
        default:
            ...
    }
    */

    level := "青铜"
    switch level {
        case "青铜":
            fmt.Println("倔强" + level)
        case "白银":
            fmt.Println("轶序" + level)
        case "黄金":
            fmt.Println("荣耀" + level)
        case "铂金":
            fmt.Println("尊贵" + level)
        case "钻石":
            fmt.Println("永恒" + level)
        case "星耀":
            fmt.Println("至尊" + level)
        case "王者":
            fmt.Println("最强" + level)
        default:
            fmt.Println("荣耀王者")
    }

    // 变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
    // 可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
    var grade string = ""
    var marks int = 90
    switch marks {
        case 90: grade = "A"
        case 80: grade = "B"
        case 50,60,70 : grade = "C"
        default: grade = "D"
    }
    fmt.Println(grade)

    warnned := true
    switch {
        case marks >= 90: 
            grade = "A"
        case marks >= 80 && marks < 90: 
            grade = "B"
        case marks >= 60 && marks < 80: 
            grade = "C"
        case marks < 60 && warnned == false: 
            warnned = true
            grade = "D"
        default: 
            grade = "退学吧"
    }
    fmt.Println(grade)

    /*
    switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
    Type Switch 语法格式如下：
    switch x.(type){
        case type:
            statement(s);      
        case type:
            statement(s); 
        default: // 可选
            statement(s);
    }
    */
    var x interface{}
    switch i := x.(type) {
       case nil:  
          fmt.Printf(" x 的类型 :%T",i)
       case int:  
          fmt.Printf("x 是 int 型")
       case float64:
          fmt.Printf("x 是 float64 型")
       case func(int) float64:
          fmt.Printf("x 是 func(int) 型")
       case bool, string:
          fmt.Printf("x 是 bool 或 string 型" )
       default:
          fmt.Printf("未知型")
    }
    
}