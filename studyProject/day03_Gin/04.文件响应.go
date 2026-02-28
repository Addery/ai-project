package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 静态文件
	r.Static("file", "static")
	r.StaticFile("test", "static/test.txt")

	// 后端实现的文件下载
	r.GET("/upFile", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")               // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		c.Header("Content-Disposition", "attachment; filename=04.文件响应.go") // 用来指定下载下来的文件名
		c.File("04.文件响应.go")
	})

	fmt.Println("server running in port :8080")
	r.Run(":8080")
}
