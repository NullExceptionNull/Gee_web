package main

import (
	"GeeWeb/com.zjh.base/gee02"
	"net/http"
)

func main() {
	engine := gee02.New()
	engine.GET("/", func(ctx *gee02.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	engine.GET("/hello", func(ctx *gee02.Context) {
		ctx.String(http.StatusOK, "hello %s ,you are %s\n", ctx.Query("name"), ctx.Path)
	})

	engine.POST("/login", func(ctx *gee02.Context) {
		ctx.JSON(http.StatusOK, gee02.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})
	engine.Run(":9999")
}
