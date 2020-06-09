package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 解析模板
	router.LoadHTMLGlob("templates/**/*")
	// 设置静态文件的请求路径和实际位置
	router.Static("/static", "./static")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "baidu.com",
		})
	})

	router.GET("/users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "/users/index",
		})
	})

	router.GET("/posts/index", func(context *gin.Context) {
		// 指定 templates中的文件名字
		context.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "/posts/index",
		})
	})
	router.Run(":9099")
}
