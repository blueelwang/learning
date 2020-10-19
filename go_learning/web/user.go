package web

import (
	"net/http"
	"strconv"
	"fmt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 读Cookie
	ult, _ 	:= r.Cookie("ULT")


	uid 	:= r.PostFormValue("uid")
	passwd 	:= r.PostFormValue("passwd")
	cookie := http.Cookie{
		Name: "ULT",
		Value: uid + "." + strconv.Itoa(len(passwd)),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "ULT[" + ult.Name + "_" + ult.Value + "][" + uid + "_" + passwd + "]")
}