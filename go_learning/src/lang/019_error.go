package lang

import (
    "errors"
    "fmt"
)


/*
Go内置了一个接口 error，定义如下：
type error interface {
	Error() string
}

errors 包下面还定义了一个 errorString 类型并实现了 error 接口：
type errorString struct {
	s string
}
func (e *errorString) Error() string {
	return e.s
}

errors 包中还定义了一个 New() 函数，用于创建一个 error 接口的实例：
func New(text string) error {
	return &errorString{text}
}

可以给自己写的函数，添加一个 error 类型的返回值，出错时调用 errors.New("...") 生成一个变量返回，没有出错时可以直接返回 nil
调用此函数时，根据返回的 error 是不是 nil 来判断调用是否出错。


*/

func Demo019() {
	res, err := divide(1, 0)
	fmt.Println(res, err.Error())	// 0 can not divide by 0

    res, err = divide(1, 3.14)		// 0.3184713375796178 <nil>
    fmt.Println(res, err)
}

func divide(a , b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("can not divide by 0")
    }
	return a / b, nil
}