package lang

import (
	"reflect"
	"fmt"
)

type Student struct {
	Code string		// 字段名是小写的话，不能通过反射获取value，不然报错：cannot return value obtained from unexported field or method
	Name string
	Age int
}

type INT int

func DemoReflect() {

	var s Student = Student{"001", "张三", 18}
	

	var t reflect.Type = reflect.TypeOf(s)
	fmt.Println(t.PkgPath())		// lang
	fmt.Println(t.Name())			// Student
	fmt.Println(t.NumField())		// 3
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	sf, _ := t.FieldByName("Age")	// StructField
	fmt.Println(sf.Name)			// Age

	v2 :=reflect.New(t) // 根据类型创建一个新的对象
	fmt.Println(v2)
	var s2 *Student = v2.Interface().(*Student)
	s2.Age = 16
	fmt.Println(s2)
	


	// INT 和 int 是两个类型，不相等
	var i1 INT = 13
	var i2 int = 13

	t1 := reflect.TypeOf(i1)
	t2 := reflect.TypeOf(i2)
	if t1 == t2 {
		println("t1 == t2")
	} else {
		println("t1 != t2") // 输出这句
	}

	// 底层类型
	fmt.Println(t1.Kind()) // int
	fmt.Println(t2.Kind()) // int
	if t1.Kind() == t2.Kind() {
		println("t1.Kind() == t2.Kind()") // 输出这句
	} else {
		println("t1.Kind() == t2.Kind()")
	}

	var v reflect.Value = reflect.ValueOf(s)
	fmt.Println(v)
	fmt.Println(v.Type() == t) // true

	for i:= 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Println(field)
		fmt.Println(value)
	}




}