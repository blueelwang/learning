package web

import (
	"net/http"
)

func ResponseHeader(w http.ResponseWriter, r *http.Request) {

	// // curl -vL "http://127.0.0.1:9000/response/redirect"

	// 写响应头
	w.Header().Set("Location", "/index")
	w.Header().Set("Custom-Key1", "value1")

	// 写响应码，省略此调用，默认为200。
	// 此调用之后可以继续写响应体，但是不能再写响应头
	w.WriteHeader(302)
	w.Header().Set("Custom-Key2", "value2")	// 实际未写入
}

func ResponseContent(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:9000/response/welcome
	str := `<html>
<head><title>Welcome</title></head>
<body>Hello</body>
</html>`
	w.Write([]byte(str))
}

