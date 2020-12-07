package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(writer, "URL.PATH = %q\n", request.URL.Path)
	case "/hello":
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(writer, "404 NOT FOUND %s\n", request.URL)
	}
}

func main() {
	http.ListenAndServe(":9999", new(Engine))
}
