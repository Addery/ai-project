package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	fmt.Println("Home")
	c.String(200, "Home")
}

func GM1(c *gin.Context) {
	// 值传递
	c.Set("Value", "GM1")
	fmt.Println("GM1 请求部分")
	c.Next()
	fmt.Println("GM1 响应部分")
}

func GM2(c *gin.Context) {
	fmt.Println("GM2 请求部分")
	c.Next()
	fmt.Println("GM2 响应部分")
	fmt.Println(c.Get("Value"))
}

func main() {
	r := gin.Default()
	g := r.Group("api")
	g.Use(GM1, GM2)
	g.GET("users", Home)
	r.Run(":8080")
}
