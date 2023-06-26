package controllers

import (
	"github.com/gin-gonic/gin"
)

/*
* 这个层的文件主要是用来处理逻辑 然后返回html
 */
func Good(c *gin.Context) {
	// 可以编写控制层
	c.HTML(200, "index.html", gin.H{
		"title": "name",
	})
}

func Good1(c *gin.Context) {
	c.String(200, "成功")
}
