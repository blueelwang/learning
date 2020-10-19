package web

import (
	"net/http"
	"math/rand"
	"strconv"
	"fmt"
)

func Cookie(w http.ResponseWriter, r *http.Request) {
	GetCookie(w, r)

	c1 := http.Cookie{
		Name: "randValue",
		Value: strconv.Itoa(rand.Intn(100000)),
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "randValue2",
		Value: strconv.Itoa(rand.Intn(100000)),
		MaxAge: 86400,
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())	// String 方法返回一个经过序列化处理的字符串
	
	// 也可以通过 http.SetCookie() 来设置 cookie，注意第二个参数是指针
	http.SetCookie(w, &c2)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]	// 返回一个字符串切片
	fmt.Fprintln(w, h)

	// 也可以r.Cookie()获取解析好的，返回一个 http.Cookie 结构，不存在返回err，如果多个只返回一个
	randValue, err := r.Cookie("randValue")
	if err == http.ErrNoCookie {
		fmt.Println(w, "cookie[randValue] not found")
	}

	fmt.Println(w, randValue)

}