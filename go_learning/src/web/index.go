package web

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request.Method)
	fmt.Fprintln(writer, request.URL.Path)
	fmt.Fprintln(writer, request.Proto)
	fmt.Fprintln(writer, request.Header)
	fmt.Fprintln(writer, request.Body)
}

func WebDemo() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8900", nil)

}