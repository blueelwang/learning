package cgo

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"	// import C 上面的注释会被解析，所以上面不能随便写注释


func Demo() {
	// 没有释放使用C.CString创建的C语言字符串会导致内存泄漏
	C.puts(C.CString("Hello, World\n"))
}