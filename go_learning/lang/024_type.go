package lang

import (
	"fmt"
)

type MyString	string
type MyString2	MyString
type MyString3	string

func TypeDemo() {
	var s1 MyString 	= "1111111111"
	fmt.Println(s1)

	// 不同的类型不能赋值，即使是类型的别名。
	// var s2 MyString2 	= s1 	// cannot use s1 (type MyString) as type MyString2 in assignment
	// s1 = s2

	// 底层类型相同的，可以强转
	var s2 MyString2 = (MyString2)(s1)
	var s3 MyString3 = (MyString3) (s1)

	fmt.Println(s1, s2, s3)


}
