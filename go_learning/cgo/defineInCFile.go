package cgo

/*
void SayHello(const char* s);
*/
import "C"

func DemoCFile() {
    C.SayHello(C.CString("defined in c file"))
}
