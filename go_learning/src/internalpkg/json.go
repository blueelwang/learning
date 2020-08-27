package internalpkg

import(
	"fmt"
	"encoding/json"
)

func JsonDemo() {

	i := int(1)
	j, _ := json.Marshal(i)
	fmt.Printf("json=%s\n", j)	// json=1

	l := int64(2)
	j, _ = json.Marshal(l)
	fmt.Printf("json=%s\n", j)	// json=2

	f := float32(3.1)
	j, _ = json.Marshal(f)
	fmt.Printf("json=%s\n", j)	// json=3.1

	d := float64(3.14)
	j, _ = json.Marshal(d)
	fmt.Printf("json=%s\n", j)	// json=3.14

	c := 'A'
	j, _ = json.Marshal(c)
	fmt.Printf("json=%s\n", j)	// json=65

	s := "string"
	j, _ = json.Marshal(s)
	fmt.Printf("json=%s\n", j)	// json="string"


	names := []string{"zhang3", "li4"}
	j, _ = json.Marshal(names)
	fmt.Printf("json=%s\n", j)	// json=["zhang3","li4"]

	dict := map[string]interface{}{
		"i": i,
		"l": l,
		"f": f,
		"d": d,
		"string": s,
		"slice": names,
	}
	j, _ = json.Marshal(dict)
	fmt.Printf("json=%s\n", j)	// json={"d":3.14,"f":3.1,"i":1,"l":2,"slice":["zhang3","li4"],"string":"string"}

	type Student struct{
		Name 	string	`json:"name"`
		Age		uint8	`json:"age"`
	}
	zhang3 := Student{
		Name: "zhang3",
		Age: 18,
	}
	j, _ = json.Marshal(zhang3)
	fmt.Printf("json=%s\n", j)	// json={"name":"zhang3","age":18}
	
	var stu3 *Student
	err := json.Unmarshal(j, stu3)
	fmt.Printf("json=%+v err=%s\n", stu3, err)	// json=<nil> err=json: Unmarshal(nil *internalpkg.Student)

	stu3 = &Student{}
	err = json.Unmarshal(j, stu3)
	fmt.Printf("json=%+v err=%s\n", stu3, err)	// json=&{Name:zhang3 Age:18} err=%!s(<nil>)

	// 接收实例有多余字段，不报错，字段使用为默认值
	type MoreField struct {
		Name	string	`json:"name"`
		Age		uint8	`json:"age"`
		Gender	string	`json:"gender"`
	}
	moreField := MoreField{}
	err = json.Unmarshal(j, &moreField)
	fmt.Printf("json=%+v err=%s\n", moreField, err)	// json={Name:zhang3 Age:18 Gender:} err=%!s(<nil>)

	// 接收实例比json少字段，也不报错
	type LessField struct {
		Name	string	`json:"name"`
	}
	lessField := LessField{}
	err = json.Unmarshal(j, &lessField)
	fmt.Printf("json=%+v err=%s\n", lessField, err)	// json={Name:zhang3} err=%!s(<nil>)


	// 指针
	type Score struct {
		Score	float32
		Student	*Student
	}
	score := Score{
		Score:		80,
		Student:	stu3,
	}
	j, _ = json.Marshal(score)
	fmt.Printf("json=%s\n", j)	// json={"Score":80,"Student":{"name":"zhang3","age":18}}
	err = json.Unmarshal(j, &score)
	fmt.Printf("json=%+v err=%s\n", score, err)	// json={Score:80 Student:0xc000090240} err=%!s(<nil>)





	
}