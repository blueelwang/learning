package internalpkg

import (
	"regexp"
	"fmt"
)

func RegExpDemo() {
	reg, err := regexp.Compile("<div>(.+)</div>")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := "<div>content</div><div>content2</div><div>content3</div>"

	res := reg.FindAllStringSubmatch(str, -1)
	fmt.Printf("%#v, len:%d\n", res, len(res))

	// [][]string{
	// 	[]string{
	// 		"<div>content</div><div>content2</div><div>content3</div>", 
	// 		"content</div><div>content2</div><div>content3"
	// 	}
	// }




	// reg, err = regexp.CompilePOSIX("<div>(.+)</div>")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// res = reg.FindAllStringSubmatch(str, 1)
	// fmt.Printf("%#v\n", res)
}
