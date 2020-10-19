package lang

import "fmt"

func DemoIf() {

    value := 1
    if value > 0 {  // 和方法的定义相同，这里的括号也不能单起一行，只能放在行尾
        fmt.Printf("%d > 0\n", value)
    } else if value < 0 {
        fmt.Printf("%d < 0\n", value)
    } else {
        fmt.Printf("%d == 0\n", value)
    }

    /*
    if 表示式中只能传入bool类型的值，不允许传入整型值，否编译不通过：non-bool 1 (type int) used as if condition
    if 1 {   
        value = 2
    }

    if 表示式中不能定义变量，syntax error: value2 := true used as value
    if value2 := true {
        value = 1;
    }
    
    也不能有赋值操作，syntax error: assignment value2 = false used as value
    value2 := true
    if value2 = false {
        value = 1;
    }
    */

    

}