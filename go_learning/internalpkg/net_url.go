package internalpkg

import (
	"net/url"
	"fmt"
)

func NetUrlDemo() {
	str := "aaa bbb http://www.baidu.com/s?wd=b bb"
	fmt.Println(url.QueryEscape(str))

	fmt.Println(url.QueryUnescape("|%20|+|%|"))
	fmt.Println("----------------------------------")
}

