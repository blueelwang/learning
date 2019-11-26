package lang

import "fmt"

func Demo003() {

    /*
    算术运算符
    +   相加    A + B 输出结果 30
    -   相减    A - B 输出结果 -10
    *   相乘    A * B 输出结果 200
    /   相除    B / A 输出结果 2
    %   求余    B % A 输出结果 0
    ++  自增    A++ 输出结果 11
    --  自减    A-- 输出结果 9

    整型和浮点型变量不能直接运算，除非是两个字面常量
    */

    A, B, C, D := 7, 5, 3.14, 3.0
    fmt.Println(A + B) // 12
    fmt.Println(A - B) // 2
    fmt.Println(A * B) // 35
    fmt.Println(A / B) // 1
    fmt.Println(A % B) // 2
    fmt.Println(C + D) // 6.140000000000001
    // fmt.Println(A + C)   // invalid operation: A + C (mismatched types int and float64)
    // fmt.Println(A + 1.1) // 编译报错：constant 1.1 truncated to integer
    fmt.Println(7 + 1.1)    // 8.1
    // 浮点运算不精确
    fmt.Println(1.5 - 1.3 == 0.2)   // false
    // 负数除法，同 C、Java、PHP 等语言
    fmt.Println(5 / 3)      // 1
    fmt.Println(-5 / 3)     // -1
    fmt.Println(5 / -3)     // -1
    // 负数取模，同 C、Java、PHP 等语言，但是和 Lua 不同
    fmt.Println(5 % 3)      // 2
    fmt.Println(-5 % 3)     // -2
    fmt.Println(5 % -3)     // 2



    /*
    关系运算符
    ==  检查两个值是否相等，如果相等返回 true 否则返回 false。
    !=  检查两个值是否不相等，如果不相等返回 true 否则返回 false。
    >   检查左边值是否大于右边值，如果是返回 true 否则返回 false。
    <   检查左边值是否小于右边值，如果是返回 true 否则返回 false。
    >=  检查左边值是否大于等于右边值，如果是返回 true 否则返回 false。
    <=  检查左边值是否小于等于右边值，如果是返回 true 否则返回 false。
    */
    fmt.Println(A > B)   // true
    fmt.Println(3 <= 2)  // false
    fmt.Println(A == 7)  // true
    fmt.Println(A != 7)  // false
    // fmt.Println(A > 7.1)  // 编译报错
    // fmt.Println(A > nil)  // 编译报错



    /*
    逻辑运算符
    下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。
    &&	逻辑 AND 运算符
    ||	逻辑 OR 运算符
    !	逻辑 NOT 运算符
    */
    truevalue, falsevalue := true, false
    fmt.Println(truevalue && falsevalue)    // false
    fmt.Println(truevalue || falsevalue)    // true

    longstring := `
    123480
    `
    fmt.Println(longstring)

    
    /*
    位运算符
    位运算符对整数在内存中的二进制位进行操作。下表列出了位运算符 
    &   按位与，这里是双目运算符，注意和取址运算符区分
    |   按位或
    ^   按们异或，相同为1，不同为0
    <<  左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
    >>  右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。低位丢弃，高位补0（正数）或 补1（负数）
    */
    fmt.Println(1 << 3)     // 8
    fmt.Println(16 >> 3)    // 2
    fmt.Println(2 >> 3)     // 0
    fmt.Println(-16 >> 3)   // -2
    fmt.Println(-2 >> 3)    // -1


    /*
    赋值运算
    =	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
    +=	相加后再赋值	C += A 等于 C = C + A
    -=	相减后再赋值	C -= A 等于 C = C - A
    *=	相乘后再赋值	C *= A 等于 C = C * A
    /=	相除后再赋值	C /= A 等于 C = C / A
    %=	求余后再赋值	C %= A 等于 C = C % A
    <<=	左移后赋值	C <<= 2 等于 C = C << 2
    >>=	右移后赋值	C >>= 2 等于 C = C >> 2
    &=	按位与后赋值	C &= 2 等于 C = C & 2
    ^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
    |=	按位或后赋值	C |= 2 等于 C = C | 2
    */
    A++
    A--
    // Go 的自增，自减只能单独语句使用，而不能用作表达式嵌入其他语句中。
    // B = A++  // 编译不通过，syntax error: unexpected ++ at end of statement



    /*
    其他运算符
    &   返回变量存储地址	&a; 将给出变量的实际地址。
    *   指针变量。	*a; 是一个指针变量
    */


    // Go 不支持三目表示式 ?:


}