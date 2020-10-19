package web

import(
	"fmt"
	"io"
	"net/http"
)



func Request(w http.ResponseWriter, r *http.Request) {
	// Url(w, r)
	// RequestHeader(w, r)
	// Body(w, r)
	// Form(w, r)
	// FormValue(w, r)
}

func Url(w http.ResponseWriter, r *http.Request) {
	// type URL struct {
	// 	Scheme     string
	// 	Opaque     string    // encoded opaque data
	// 	User       *Userinfo // username and password information
	// 	Host       string    // host or host:port
	// 	Path       string    // path (relative paths may omit leading slash)
	// 	RawPath    string    // encoded path hint (see EscapedPath method)
	// 	ForceQuery bool      // append a query ('?') even if RawQuery is empty
	// 	RawQuery   string    // encoded query values, without '?'
	// 	Fragment   string    // fragment for references, without '#'
	// }

	fmt.Fprintf(w, "r.URL=%+v\n", r.URL)				
	fmt.Fprintf(w, "r.URL.Scheme=%s\n", r.URL.Scheme)	
	fmt.Fprintf(w, "r.URL.Host=%s\n", r.URL.Host)		
	fmt.Fprintf(w, "r.URL.Path=%s\n", r.URL.Path)		
	fmt.Fprintf(w, "r.URL.RawPath=%s\n", r.URL.RawPath)
	fmt.Fprintf(w, "r.URL.ForceQuery=%v\n", r.URL.ForceQuery)
	fmt.Fprintf(w, "r.URL.RawQuery=%s\n", r.URL.RawQuery)
	fmt.Fprintf(w, "r.URL.Fragment=%s\n", r.URL.Fragment)

	/*
	请求url:
		http://127.0.0.1:9000/request/url?key=value#anchor
	输出：
		r.URL=/request/url?key=value
		r.URL.Scheme=
		r.URL.Host=
		r.URL.Path=/request/url
		r.URL.RawPath=
		r.URL.ForceQuery=false
		r.URL.RawQuery=key=value
		r.URL.Fragment=					// 这里为空是因为浏览器的原因，请求时会把这部分去掉
	*/
}

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	// 	map[
	// 		Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9] 
	// 		Accept-Encoding:[gzip, deflate, br] 
	// 		Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7] 
	// 		Cache-Control:[max-age=0] 
	// 		Connection:[keep-alive] 
	// 		Cookie:[ULT=.0] 
	// 		Sec-Fetch-Dest:[document] 
	// 		Sec-Fetch-Mode:[navigate] 
	// 		Sec-Fetch-Site:[none] 
	// 		Sec-Fetch-User:[?1] 
	// 		Upgrade-Insecure-Requests:[1] 
	// 		User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36]
	// 	]

	fmt.Fprintln(w, h["Accept-Encoding"])	// [gzip, deflate, br]	取到的值是一个字符串切片
	fmt.Fprintln(w, h.Get("Accept-Encoding")) // gzip, deflate, br	取到一个字段串，如果有多个值，用逗号隔开
}

func Body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	rlen, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Fprintf(w, "read body err: %s, read len:%d", err, rlen)
	}
	fmt.Fprintln(w, string(body))
	// r.Body.Close()
}

func Form(w http.ResponseWriter, r *http.Request) {
	
	// curl -X POST "http://127.0.0.1:9000/request/url?key=value#anchor" -d "key2=value2&key=newvalue"

	r.ParseForm()				// 不调用，r.Form为空map，
	// 如果之前通过 r.Body.Read() 读过body数据了，这里就不会再解析body中的数据了，下面的 r.Form 和 r.PostForm 也不会包含body中的数据。

	// r.Form 值是url解码后的值，包含url中的参数和body中的参数，body中的数据在前
	fmt.Fprintln(w, r.Form)		// map[key:[newvalue value] key2:[value2]]， 注意:newvalue在value之前
	
	// r.PostForm 只包含body中解析出的数据
	fmt.Fprintln(w, r.PostForm)	// map[key:[newvalue] key2:[value2]]

}

func FormValue(w http.ResponseWriter, r *http.Request) {

	// curl -X POST "http://127.0.0.1:9000/request/url?key=value&key2=value2#anchor" -d "key=newvalue&key3=value3"

	// r.FormValue() 会在需要的时候自动调用 r.ParseForm() 和 r.ParseMultipartForm()，无需再手动调用。
	// 当给定键有多个值时，只会从 Form 中取第一个值
	// r.FormValue() 从 r.Form 中返回数据
	// r.PostFormValue() 从 r.PostForm 中返回数据，只包含body中的数据
	fmt.Fprintln(w, r.FormValue("key"))			// newvalue
	fmt.Fprintln(w, r.PostFormValue("key2"))	// 空
	fmt.Fprintln(w, r.Form["key"])				// [newvalue value]
	fmt.Fprintln(w, r.FormValue("key2"))		// value2
	fmt.Fprintln(w, r.Form["key2"])				// [value2]
}

func MultipartForm(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		buf := make([]byte, 1024)
		len, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Fprintln(w, "read file failed, err:%s", err)
			return
		}
		fmt.Fprintln(w, string(buf[0:len]))
	}
}



