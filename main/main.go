package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func main() {
	r := gin.Default()
	// 添加 get 路由和回调

	r.LoadHTMLFiles("./views/index.html")

	//r.GET("/get", controllers.Good1)
	r.GET("/get", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "学习",
		})
	})

	// 添加 post 路由和回调
	r.POST("/post", func(c *gin.Context) {
		// 返回的 code 和 字符串返回
		c.String(200, "这是一个 post 方法\n")
	})

	// 添加 delete 路由和回调
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		// 返回的 code 和 字符串返回
		c.String(200, "这是一个 delete 方法\n")
	})

	// 添加 any 路由和回调
	r.Any("/any", func(c *gin.Context) {
		// 返回的 code 和 字符串返回
		c.String(200, "这是一个 any 方法\n")
	})

	// 使用包函数打印
	fmt.Printf("版本: %s \n", runtime.Version())

	// 启动框架程序, 默认 8080 端口
	r.Run()
}
