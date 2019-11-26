package lang

import "fmt"

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



