package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "25")
		values := c.QueryArray("key")
		fmt.Println(name, age, values)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
