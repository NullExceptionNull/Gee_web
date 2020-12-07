package main

import (
	"GeeWeb/com.zjh.base/gee01"
	"fmt"
	"net/http"
)

func main() {
	engine := gee01.New()

	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	engine.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	engine.Run(":9999")
}
