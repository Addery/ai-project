package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 加载模板 只有这里加载了模板，下面才能用
	r.LoadHTMLGlob("templates/*")
	//r.LoadHTMLFiles("templates/index.html")
	r.GET("", func(c *gin.Context) {
		c.HTML(200, "index.html", map[string]any{
			"title": "title内容",
		})
	})
	r.Run(":8080")
}
