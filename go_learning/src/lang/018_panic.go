package lang

import "fmt"

// defer 关键字
// Go 函数里提供了 defer关键字，可以注册多个延迟调用，这些调用以 先进后出 的顺序在函数返回前被执行（即最后注册的最先执行）。
// 有点类似于 Java 语言中异常处理中的 finaly 子句。 defer 常用于保证一些资源最终一定能够得到回收和释放。
func deferDemo() {
    i := 1
    defer func(i int) {
        fmt.Printf("我在第%d个 defer 注册的方法中\n", i)
    }(i)    // i 此时为1
    i++
    defer func(i int) {
        fmt.Printf("我在第%d个 defer 注册的方法中\n", i)
    }(i)    // i 此时为2
    // defer 函数的实参在注册时通过值拷贝传递进去。后续再对传入的参数变量做改动，也不再会改变原来注册时传入参数的值
    i++     // defer 方法在实际被调用时，传入的参数是在注册时i的值，这里修改不会产生影响
    fmt.Println("我在所有 defer 注册的方法之后")
    if true {
        return
    }
    
    // defer语句必须先注册后才能执行，如果 defer 位于 return 之后，则等于没有注册，不会执行。
    defer func(i int) {
        fmt.Printf("我在第%d个 defer 注册的方法中\n", i)
    }(i)    // i 此时为1
    /*
    输出：
    我在所有 defer 注册的方法之后
    我在第2个 defer 注册的方法中
    我在第1个 defer 注册的方法中
    */

    // 注意：
    // 主动 调用 as.Exit(int) 退出进程时，defer将不再被执行(即使已经提前注册〉。
}

/*
Go 在运行时发生错误时，也像Java语言一样，会抛出异常，也可以在代码中调用 panic() 手动抛出异常。
抛出异常时，当前方法里未执行到的语句不再执行，而是直接返回，不过返回前也会正常调用 defer 注册的函数。

recover() 函数可以捕获异常。
如果函数内所有defer函数都没有用recover()捕获异常，上层调用此函数的地方，也不会执行函数体后面的语句，会继续抛出异常。
如果上层方法依然没有捕获异常，会继续向上层调用处抛出，直到异常被捕获，或者运行到最外层函数而退出。

注意：recover() 只有在 defer 后面的函数体内被直接调用才能捕获异常，否则返回 nil，继续向外抛异常，详见error2()中错误使用

当在defer函数中继续调用panic()抛出异常时，后面捕获的异常是最后一次抛出的

函数并不能捕获内部新启动的 goroutine 所抛出的异常
*/

func Demo018() {
	deferDemo()
	defer ignore("defer in Main()")
	error1()
	fmt.Println("after call error1()")
}





func error1() {
	defer catch()
	error2()
	fmt.Println("after call error2()")  // error2()里没有recover()调用方式不对，没有成功捕获异常，这里走不到
}

func error2() {
	// 如下面这几种使用捕获异常会失败
	defer recover()
	defer fmt.Println(recover())
	// 这个嵌套两层也会捕获失败
	defer func() {
		func() {
			recover()
		}()
	}()

	error3()
	fmt.Println("after call error3()")    // error3()里没有调用recover()处理异常，这里走不到
}

func error3() {
	defer ignore("defer in error3()")
	panic("custom error occurred")
	// fmt.Println("after panic()")    // panic()之后走不到
}

func ignore(msg string) {
	fmt.Println("ignore:" + msg)
}

func catch() {
	fmt.Print("recovering ...\ngot error: ")
	err := recover()
	fmt.Println(err)
}

/*
ignore:defer in error3()
<nil>
recovering ...
got error: custom error occurred
after call error1()
ignore:defer in Main()
*/



