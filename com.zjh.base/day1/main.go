package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	_ = http.ListenAndServe(":9999", nil)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.PATH = %q\n", request.URL.Path)
}
