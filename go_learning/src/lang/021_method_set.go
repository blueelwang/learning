package lang

type Data struct{}

/*
将接收者为值类型 T 的方法的集合记录为 S，
将接收者为指针类型 *T 的方法的集合统称为 *S
T  类型的方法集为 S
*T 类型的方法集为 S 和 *S

在直接使用类型实例调用类型的方法时，无论值类型变量还是指针类型变量，
都可以调用类型的所有方法，原因是编译器在编译期间做了自动的转换。

但是如果通过方法的 值调用 和 表达式调用 时，有些情况编译器是没有做自动转换的。

通过类型字面量显式地进行值调用和表达式调用，编译器不会做自动转换。

*/
func (Data) TestValue() {}
func (*Data) TestPointer() {}

func Demo021() {

	// 通过类型实例调用方法，无论值类型实例还是指针类型的实例，都能调用全部方法
	value := Data{}
	value.TestPointer()
	value.TestPointer()
	ptr := &value
	ptr.TestPointer()
	ptr.TestValue()
	// 如果类型实例是一个表达式，编译器不会做自动转换。
	(*Data)(&struct{}{}).TestPointer()
	(*Data)(&struct{}{}).TestValue()
	// (Data)(struct{}{}).TestPointer()		//编译不通过
	(Data)(struct{}{}).TestValue()
	getPointer().TestPointer()
	getPointer().TestValue()
	// getValue().TestPointer()				//编译不通过
	getValue().TestValue()




	// 使用表达式调用时编译器不会进行转换
	Data.TestValue(value)
	// Data.TestPointer(ptr)
	// Data.TestValue(ptr)

	(*Data).TestPointer(ptr)
	(*Data).TestValue(ptr)
	// (*Data).TestValue(value)
	// (*Data).TestPointer(value)




	// 使用值调用方式调用时编译器会进行自动转换
	f := value.TestPointer
	f()
	f = value.TestValue
	f()
	f = ptr.TestPointer
	f()
	f = ptr.TestValue
	f()
	
	
	
}

func getValue() Data {
	return Data{}
}

func getPointer() *Data {
	return &Data{}
}