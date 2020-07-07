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
	ServeWithDefaultServeMux()
	// ServeWithCustomServeMux()
	// HttpsDemo()
}

func ServeWithDefaultServeMux() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":9000", nil)
}

func ServeWithCustomServeMux() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr: ":9001",
		Handler: mux,
	}
	server.ListenAndServe()
}

func HttpsDemo() {
	http.ListenAndServeTLS(":9002", "cert.pem", "key.pem", nil)
}